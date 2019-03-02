package helpers

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var _ grpc.UnaryServerInterceptor = UnaryServerInterceptor
var _ grpc.StreamServerInterceptor = StreamServerInterceptor

func UnaryServerInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer handleCrash(ctx, func(ctx context.Context, r interface{}) {
		err = toError(r)
	})

	return handler(ctx, req)
}

func StreamServerInterceptor(srv interface{}, stream grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer handleCrash(stream.Context(), func(ctx context.Context, r interface{}) {
		err = toError(r)
	})

	return handler(srv, stream)
}

func handleCrash(ctx context.Context, handler func(context.Context, interface{})) {
	if r := recover(); r != nil {
		handler(ctx, r)
	}
}

func toError(r interface{}) error {
	return grpc.Errorf(codes.Internal, "%v", r)
}
