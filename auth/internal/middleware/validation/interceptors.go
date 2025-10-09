package middleware_validation

import (
	"context"

	"buf.build/go/protovalidate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func ValidationInterceptor(validator protovalidate.Validator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		if protoMsg, ok := req.(proto.Message); ok {
			if err := validator.Validate(protoMsg); err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "Validation error: %v", err)
			}
		}

		return handler(ctx, req)
	}
}
