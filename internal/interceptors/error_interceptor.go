package interceptors

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func APIErrorInterceptor(zapLogger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		resp, err := handler(ctx, req)
		if err == nil {
			return resp, nil
		}
		zapLogger.Error(fmt.Sprintf("error in method: %s", info.FullMethod), zap.Error(err))

		if errGrpc, ok := status.FromError(err); ok {
			return nil, errGrpc.Err()
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}
}
