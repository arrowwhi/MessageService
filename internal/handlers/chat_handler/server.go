package chat_handler

import (
	"MessageService/internal/interfaces/service"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	pb "MessageService/proto/message_service/chats/v1"
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
	pb.RegisterChatServiceServer(server, s)
}

func (s *Service) RegisterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return pb.RegisterChatServiceHandler(ctx, mux, conn)
}
