/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
