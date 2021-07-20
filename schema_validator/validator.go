// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package grpc_schema_validator

import (
	"bytes"
	"context"

	"github.com/xeipuuv/gojsonschema"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type validator interface {
	Validate() (*gojsonschema.Result, error)
}

// UnaryServerInterceptor returns a new unary server interceptors that validates incoming messages.
//
// Invalid messages will be rejected with `InvalidArgument` before reaching any userspace handlers.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if v, ok := req.(validator); ok {
			result, err := v.Validate()
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
			if result != nil && !result.Valid() {
				return nil, status.Error(codes.InvalidArgument, flatten(result))
			}
		}
		return handler(ctx, req)
	}
}

// StreamServerInterceptor returns a new streaming server interceptors that validates incoming messages.
//
// The stage at which invalid messages will be rejected with `InvalidArgument` varies based on the
// type of the RPC. For `ServerStream` (1:m) requests, it will happen before reaching any userspace
// handlers. For `ClientStream` (n:1) or `BidiStream` (n:m) RPCs, the messages will be rejected on
// calls to `stream.Recv()`.
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapper := &recvWrapper{stream}
		return handler(srv, wrapper)
	}
}

type recvWrapper struct {
	grpc.ServerStream
}

func (s *recvWrapper) RecvMsg(m interface{}) error {
	if err := s.ServerStream.RecvMsg(m); err != nil {
		return err
	}
	if v, ok := m.(validator); ok {
		result, err := v.Validate()
		if err != nil {
			return status.Error(codes.InvalidArgument, err.Error())
		}
		if result != nil && !result.Valid() {
			return status.Error(codes.InvalidArgument, flatten(result))
		}
	}
	return nil
}

func flatten(result *gojsonschema.Result) string {
	var buf bytes.Buffer
	for _, r := range result.Errors() {
		buf.WriteString(r.Field() + ": " + r.Description() + "\n")
	}
	return buf.String()
}
