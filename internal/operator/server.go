package operator

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/PicoTools/pico/internal/constants"
	"github.com/PicoTools/pico/internal/ent"
	p "github.com/PicoTools/pico/internal/ent/pki"
	"github.com/PicoTools/pico/internal/middleware/grpcauth"
	"github.com/PicoTools/pico/internal/middleware/grpclog"
	"github.com/PicoTools/pico/internal/middleware/grpcrecover"
	"github.com/PicoTools/pico/internal/tls"
	operatorv1 "github.com/PicoTools/pico/pkg/proto/operator/v1"
	"github.com/PicoTools/pico/pkg/shared"

	"github.com/go-faster/sdk/zctx"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

// Serve starts serving of listener's GRPC server
func Serve(ctx context.Context, cfg Config, db *ent.Client) error {
	lg := zctx.From(ctx).Named("operator")

	tlsOpts, f, err := tls.NewTlsTransport(ctx, db, p.TypeOperator)
	if err != nil {
		return err
	}
	// create server
	srv := grpc.NewServer(
		grpc.Creds(credentials.NewTLS(tlsOpts)),
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    constants.GrpcKeepalivePeriod,
				Timeout: constants.GrpcKeepalivePeriod * time.Duration(constants.GrpcKeepaliveCount),
			},
		),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             constants.GrpcKeepalivePeriod,
				PermitWithoutStream: true,
			},
		),
		grpc.ChainUnaryInterceptor(
			grpcrecover.UnaryServerInterceptor(),
			grpclog.UnaryServerInterceptor(lg),
			grpcauth.UnaryServerInterceptorOperator(db.Operator),
		),
		grpc.ChainStreamInterceptor(
			grpcrecover.StreamServerInterceptor(),
			grpclog.StreamServerInterceptor(lg),
			grpcauth.StreamServerInterceptorOperator(db.Operator),
		),
		grpc.MaxRecvMsgSize(shared.MaxProtobufMessageSize),
		grpc.MaxSendMsgSize(shared.MaxProtobufMessageSize),
		grpc.MaxConcurrentStreams(constants.GrpcMaxConcurrentStreams),
	)
	// add service
	operatorv1.RegisterOperatorServiceServer(srv, &server{
		ctx: ctx,
		db:  db,
		lg:  lg,
	})
	// create net listener
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.IP, cfg.Port))
	if err != nil {
		return err
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		lg.Info("start serving",
			zap.String("ip", cfg.IP.String()),
			zap.Int("port", cfg.Port),
			zap.String("fingerprint", f),
		)
		return srv.Serve(l)
	})

	g.Go(func() error {
		<-ctx.Done()
		srv.Stop()
		lg.Info("stop serving")
		return nil
	})

	return g.Wait()
}
