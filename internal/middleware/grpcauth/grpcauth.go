package grpcauth

import (
	"context"
	"strings"
	"time"

	"github.com/PicoTools/pico/internal/ent"
	"github.com/PicoTools/pico/internal/ent/listener"
	"github.com/PicoTools/pico/internal/ent/operator"
	picoErrors "github.com/PicoTools/pico/internal/errors"
	"github.com/PicoTools/pico/internal/middleware"
	"github.com/PicoTools/pico/internal/pools"
	operatorv1 "github.com/PicoTools/pico/pkg/proto/operator/v1"
	"github.com/PicoTools/pico/pkg/shared"
	"github.com/go-faster/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type listenerCtxId struct{}

func ListenerToCtx(ctx context.Context, id int64) context.Context {
	return context.WithValue(ctx, listenerCtxId{}, id)
}

func ListenerFromCtx(ctx context.Context) int64 {
	return ctx.Value(listenerCtxId{}).(int64)
}

type operatorCtxId struct{}

func OperatorToCtx(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, operatorCtxId{}, username)
}

func OperatorFromCtx(ctx context.Context) string {
	return ctx.Value(operatorCtxId{}).(string)
}

// UnaryServerInterceptorListener returns unary interceptor to handle listener's GRPC authentication
func UnaryServerInterceptorListener(db *ent.ListenerClient) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Internal, picoErrors.MissingRequestMetadata)
		}
		tokens := meta.Get(shared.GrpcAuthListenerHeader)
		if len(tokens) != 1 {
			return nil, status.Error(codes.Unauthenticated, picoErrors.UnauthenticatedRequest)
		}
		v, err := db.Query().
			Where(listener.Token(tokens[0])).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, status.Error(codes.Unauthenticated, picoErrors.UnauthenticatedRequest)
			} else {
				return nil, status.Error(codes.Internal, err.Error())
			}
		}
		v, err = v.Update().
			SetLast(time.Now()).
			Save(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		// notify operators with listener last checkout update
		go pools.Pool.Listeners.Send(&operatorv1.ListenerLastResponse{
			Id:   v.ID,
			Last: timestamppb.New(v.Last),
		})
		return handler(ListenerToCtx(ctx, v.ID), req)
	}
}

// StreamServerInterceptorListener returns stream interceptor to handle listener's GRPC authentication
func StreamServerInterceptorListener(db *ent.ListenerClient) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := ss.Context()
		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return status.Error(codes.Internal, picoErrors.MissingRequestMetadata)
		}
		tokens := meta.Get(shared.GrpcAuthListenerHeader)
		if len(tokens) != 1 {
			return status.Error(codes.Unauthenticated, picoErrors.UnauthenticatedRequest)
		}
		v, err := db.Query().
			Where(listener.Token(tokens[0])).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return status.Error(codes.Unauthenticated, picoErrors.UnauthenticatedRequest)
			} else {
				return status.Error(codes.Internal, err.Error())
			}
		}
		v, err = v.Update().
			SetLast(time.Now()).
			Save(ctx)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		// notify operators with listener last checkout update
		go pools.Pool.Listeners.Send(&operatorv1.ListenerLastResponse{
			Id:   v.ID,
			Last: timestamppb.New(v.Last),
		})
		return handler(srv, &middleware.SrvStream{ServerStream: ss, Ctx: ListenerToCtx(ctx, v.ID)})
	}
}

// UnaryServerInterceptorOperator returns unary interceptor to handle operator's GRPC authentication
func UnaryServerInterceptorOperator(db *ent.OperatorClient) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Internal, picoErrors.MissingRequestMetadata)
		}
		tokens := meta.Get(shared.GrpcAuthOperatorHeader)
		if len(tokens) != 1 {
			return nil, status.Error(codes.Unauthenticated, picoErrors.UnauthenticatedRequest)
		}
		v, err := db.Query().
			Where(operator.Token(tokens[0])).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, status.Error(codes.Unauthenticated, picoErrors.UnauthenticatedRequest)
			} else {
				return nil, status.Error(codes.Internal, err.Error())
			}
		}
		v, err = v.Update().
			SetLast(time.Now()).
			Save(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		// notify operators with operator last checkout update
		go pools.Pool.Operators.Send(&operatorv1.OperatorLastResponse{
			Username: v.Username,
			Last:     timestamppb.New(v.Last),
		})
		return handler(OperatorToCtx(ctx, v.Username), req)
	}
}

// StreamServerInterceptorOperator returns stream interceptor to handle operator's GRPC authentication
func StreamServerInterceptorOperator(db *ent.OperatorClient) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := ss.Context()
		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return status.Error(codes.Internal, picoErrors.MissingRequestMetadata)
		}
		tokens := meta.Get(shared.GrpcAuthOperatorHeader)
		if len(tokens) != 1 {
			return status.Error(codes.Unauthenticated, picoErrors.UnauthenticatedRequest)
		}
		v, err := db.Query().
			Where(operator.Token(tokens[0])).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return status.Error(codes.Unauthenticated, picoErrors.UnauthenticatedRequest)
			} else {
				return status.Error(codes.Internal, errors.Wrap(err, "query operator token").Error())
			}
		}
		v, err = v.Update().
			SetLast(time.Now()).
			Save(ctx)
		if err != nil {
			return status.Error(codes.Internal, errors.Wrap(err, "update operator last").Error())
		}
		// notify operators with operator last checkout update
		go pools.Pool.Operators.Send(&operatorv1.OperatorLastResponse{
			Username: v.Username,
			Last:     timestamppb.New(v.Last),
		})
		return handler(srv, &middleware.SrvStream{ServerStream: ss, Ctx: OperatorToCtx(ctx, v.Username)})
	}
}

// UnaryServerInterceptorManagement returns stream interceptor to handle management's GRPC authentication
func UnaryServerInterceptorManagement(token string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Internal, picoErrors.MissingRequestMetadata)
		}
		tokens := meta.Get(shared.GrpcAuthManagementHeader)
		if len(tokens) != 1 {
			return nil, status.Error(codes.Unauthenticated, picoErrors.UnauthenticatedRequest)
		}
		if strings.Compare(tokens[0], token) != 0 {
			return nil, status.Error(codes.Unauthenticated, picoErrors.UnauthenticatedRequest)
		}
		return handler(ctx, req)
	}
}

// StreamServerInterceptorManagement returns stream interceptor to handle management's GRPC authentication
func StreamServerInterceptorManagement(token string) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := ss.Context()
		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return status.Error(codes.Internal, picoErrors.MissingRequestMetadata)
		}
		tokens := meta.Get(shared.GrpcAuthManagementHeader)
		if len(tokens) != 1 {
			return status.Error(codes.Unauthenticated, picoErrors.UnauthenticatedRequest)
		}
		if strings.Compare(tokens[0], token) != 0 {
			return status.Error(codes.Unauthenticated, picoErrors.UnauthenticatedRequest)
		}
		return handler(srv, ss)
	}
}
