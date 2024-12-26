package grpcrecover

import (
	"context"
	"fmt"
	"runtime/debug"

	picoErrors "github.com/PicoTools/pico/internal/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// recoverFrom returns internal error in case of panic
func recoverFrom(_ context.Context, p interface{}) error {
	fmt.Printf("\n%s\n%s\n", p, debug.Stack())
	return status.Errorf(codes.Internal, picoErrors.Internal)
}

// UnaryServerInterceptor returns unary interceptor for panic recovering
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				err = recoverFrom(ctx, r)
			}
		}()

		resp, err := handler(ctx, req)
		panicked = false
		return resp, err
	}
}

// StreamServerInterceptor returns stream interceptor for panic recovering
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				err = recoverFrom(stream.Context(), r)
			}
		}()

		err = handler(srv, stream)
		panicked = false
		return err
	}
}
