package service_impl

import (
	"MessageService/internal/converter"
	"MessageService/internal/interfaces/infra/postgres"
	intf "MessageService/internal/interfaces/service"
	"context"
	"go.uber.org/zap"
)

var _ intf.Service = (*service)(nil)

type service struct {
	logger      *zap.Logger
	userRepo    postgres.Users
	messageRepo postgres.Message
	chatRepo    postgres.Chats

	cvt converter.RepoConverter
}

func New(
	logger *zap.Logger,
	userRepo postgres.Users,
	messageRepo postgres.Message,
	chatRepo postgres.Chats,
	cvt converter.RepoConverter,
) intf.Service {
	return &service{
		logger:      logger,
		userRepo:    userRepo,
		messageRepo: messageRepo,
		chatRepo:    chatRepo,
		cvt:         cvt,
	}
}

func (s *service) GetMessages(ctx context.Context, request *intf.GetMessagesRequest) (*intf.GetMessagesResponse, error) {
	ans, err := s.messageRepo.GetMessages(ctx, request.ChatId,
		request.Limit,
		request.Offset,
	)
	if err != nil {
		return nil, err
	}

	return &intf.GetMessagesResponse{Messages: ans}, nil
}

func (s *service) SendMessage(ctx context.Context, request *intf.SendMessageRequest) error {
	err := s.messageRepo.AddMessage(ctx, request.Message)
	if err != nil {
		return err
	}
	//todo Notify
	return nil
}

func (s *service) UpdateMessageStatus(ctx context.Context, request *intf.UpdateMessageStatusRequest) error {
	err := s.messageRepo.UpdateMessageStatus(ctx, request.MessageId, true)
	if err != nil {
		return err
	}
	//todo Notify
	return nil
}

func (s *service) AddChat(ctx context.Context, request *intf.AddChatRequest) error {
	err := s.chatRepo.AddChat(ctx, request.Chat)
	if err != nil {
		return err
	}
	//todo Notify
	return nil
}

func (s *service) GetChats(ctx context.Context, request *intf.GetChatsRequest) (*intf.GetChatsResponse, error) {
	chats, err := s.chatRepo.GetChats(ctx, request.UserId, request.Limit, request.Offset)
	if err != nil {
		return nil, err
	}
	return &intf.GetChatsResponse{Chats: chats}, nil
}

func (s *service) GetStatusInfo(ctx context.Context, request *intf.GetStatusInfoRequest) (*intf.GetStatusInfoResponse, error) {
	s.logger.Debug("ok")
	users, err := s.userRepo.GetStatusInfo(ctx, request.Ids)
	if err != nil {
		return nil, err
	}
	return &intf.GetStatusInfoResponse{Users: users}, nil
}

func (s *service) UpdateStatus(ctx context.Context, request *intf.UpdateStatusRequest) error {
	err := s.userRepo.UpdateStatus(ctx, request.UserId, request.Online)
	if err != nil {
		return err
	}
	// todo Notify
	return nil
}
