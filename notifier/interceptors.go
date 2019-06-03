package notifier

import (
	"golang.org/x/net/context"
	utilerrors "gomodules.xyz/notify/errors"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor wraps a unary server with error notifier
//
// Invalid messages will be rejected with `Internal` before reaching any userspace handlers.
func UnaryServerInterceptor(fn grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		v, err := fn(ctx, req, info, handler)
		if err != nil {
			utilerrors.HandleError(err)
		}
		return v, err
	}
}

// StreamServerInterceptor wraps a unary server with error notifier
func StreamServerInterceptor(fn grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		err := fn(srv, stream, info, handler)
		if err != nil {
			utilerrors.HandleError(err)
		}
		return err
	}
}
