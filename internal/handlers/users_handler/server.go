package users_handler

import (
	"MessageService/internal/interfaces/service"
	pb "MessageService/proto/message_service/users/v1"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Service struct {
	service service.Service
	logger  *zap.Logger
}

func New(logger *zap.Logger, service service.Service) *Service {
	return &Service{
		logger:  logger,
		service: service,
	}
}

func (s *Service) RegisterServer(server *grpc.Server) {
	pb.RegisterUsersServiceServer(server, s)
}

func (s *Service) RegisterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return pb.RegisterUsersServiceHandler(ctx, mux, conn)
}
