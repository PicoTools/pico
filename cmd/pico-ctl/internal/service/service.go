package service

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/PicoTools/pico/cmd/pico-ctl/internal/middleware"
	managementv1 "github.com/PicoTools/pico/pkg/proto/management/v1"
	"github.com/PicoTools/pico/pkg/shared"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var mgmtConn = &grpcConn{}

type grpcConn struct {
	ctx  context.Context
	conn *grpc.ClientConn
	svc  managementv1.ManagementServiceClient
}

// Init initializes GRPC client connection
func Init(ctx context.Context, host string, token string) error {
	var err error
	mgmtConn.ctx = ctx

	if mgmtConn.conn, err = grpc.NewClient(
		host,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})),
		grpc.WithUnaryInterceptor(middleware.UnaryClientInterceptor(token)),
		grpc.WithStreamInterceptor(middleware.StreamClientInterceptor(token)),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(shared.MaxProtobufMessageSize),
			grpc.MaxCallSendMsgSize(shared.MaxProtobufMessageSize),
		),
	); err != nil {
		return err
	}

	mgmtConn.svc = managementv1.NewManagementServiceClient(mgmtConn.conn)
	return nil
}

// Close closes GRPC connection
func Close() error {
	if mgmtConn.conn != nil {
		return mgmtConn.conn.Close()
	}
	return nil
}

func getSvc() managementv1.ManagementServiceClient {
	return mgmtConn.svc
}

// AddOperator creates operator specified by username
func AddOperator(username string) (*managementv1.Operator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rep, err := getSvc().NewOperator(ctx, &managementv1.NewOperatorRequest{Username: username})
	if err != nil {
		return nil, err
	}
	return rep.GetOperator(), nil
}

// ListOperators returns list of registered operators
func ListOperators() ([]*managementv1.Operator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rep, err := getSvc().GetOperators(ctx, &managementv1.GetOperatorsRequest{})
	if err != nil {
		return nil, err
	}
	return rep.GetOperators(), nil
}

// RegenerateTokenOperator regenerates access token for operator specified by username
func RegenerateTokenOperator(username string) (*managementv1.Operator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rep, err := getSvc().RegenerateOperator(ctx, &managementv1.RegenerateOperatorRequest{Username: username})
	if err != nil {
		return nil, err
	}
	return rep.GetOperator(), nil
}

// RevokeTokenOperator revokes access token for operator specified by username
func RevokeTokenOperator(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := getSvc().RevokeOperator(ctx, &managementv1.RevokeOperatorRequest{Username: username})
	return err
}

// AddListener creates new listener
func AddListener() (*managementv1.Listener, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rep, err := getSvc().NewListener(ctx, &managementv1.NewListenerRequest{})
	if err != nil {
		return nil, err
	}
	return rep.GetListener(), nil
}

// ListListeners returns list of registered listeners
func ListListeners() ([]*managementv1.Listener, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rep, err := getSvc().GetListeners(ctx, &managementv1.GetListenersRequest{})
	if err != nil {
		return nil, err
	}
	return rep.GetListeners(), nil
}

// RegenerateTokenListener regenerates access token for listener specified by ID
func RegenerateTokenListener(id int64) (*managementv1.Listener, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rep, err := getSvc().RegenerateListener(ctx, &managementv1.RegenerateListenerRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return rep.GetListener(), nil
}

// RevokeTokenListener revokes access token for listener specified by ID
func RevokeTokenListener(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := getSvc().RevokeListener(ctx, &managementv1.RevokeListenerRequest{Id: id})
	return err
}

// GetCertCA returns CA certificate from PKI
func GetCertCA() (*managementv1.GetCertCAResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return getSvc().GetCertCA(ctx, &managementv1.GetCertCARequest{})
}

// GetCertOperator returns operator GRPC certificate from PKI
func GetCertOperator() (*managementv1.GetCertOperatorResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return getSvc().GetCertOperator(ctx, &managementv1.GetCertOperatorRequest{})
}

// GetCertOperator returns listener's GRPC certificate from PKI
func GetCertListener() (*managementv1.GetCertListenerResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return getSvc().GetCertListener(ctx, &managementv1.GetCertListenerRequest{})
}
