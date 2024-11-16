package service_impl

import (
	"MessageService/internal/converter"
	"MessageService/internal/interfaces/infra/postgres"
	"MessageService/internal/interfaces/service"
	"go.uber.org/zap"
)

var _ service.Service = (*Service)(nil)

type Service struct {
	logger      *zap.Logger
	userRepo    postgres.Users
	messageRepo postgres.Message
	chatRepo    postgres.Chats

	cvt converter.RepoConverter
}

func (s *Service) GetMessages(request *service.GetMessagesRequest) (*service.GetMessagesResponse, error) {
	ans, err := s.messageRepo.GetMessages(request.ChatId,
		request.Limit,
		request.Offset,
	)
	if err != nil {
		return nil, err
	}

	return &service.GetMessagesResponse{Messages: ans}, nil
}

func (s *Service) SendMessage(request *service.SendMessageRequest) error {
	err := s.messageRepo.AddMessage(request.Message)
	if err != nil {
		return err
	}
	//todo Notify
	return nil
}

func (s *Service) UpdateMessageStatus(request *service.UpdateMessageStatusRequest) error {
	err := s.messageRepo.UpdateMessageStatus(request.MessageId, true)
	if err != nil {
		return err
	}
	//todo Notify
	return nil
}

func (s *Service) AddChat(request *service.AddChatRequest) error {
	err := s.chatRepo.AddChat(request.Chat)
	if err != nil {
		return err
	}
	//todo Notify
	return nil
}

func (s *Service) GetChats(request *service.GetChatsRequest) (*service.GetChatsResponse, error) {
	chats, err := s.chatRepo.GetChats(request.UserId, request.Limit, request.Offset)
	if err != nil {
		return nil, err
	}
	return &service.GetChatsResponse{Chats: chats}, nil
}

func (s *Service) GetStatusInfo(request *service.GetStatusInfoRequest) (*service.GetStatusInfoResponse, error) {
	users, err := s.userRepo.GetStatusInfo(request.Ids)
	if err != nil {
		return nil, err
	}
	return &service.GetStatusInfoResponse{Users: users}, nil
}

func (s *Service) UpdateStatus(request *service.UpdateStatusRequest) error {
	err := s.userRepo.UpdateStatus(request.UserId, request.Online)
	if err != nil {
		return err
	}
	// todo Notify
	return nil
}
