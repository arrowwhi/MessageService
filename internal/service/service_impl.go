package service_impl

import "MessageService/internal/interfaces/service"

var _ service.Service = (*Service)(nil)

type Service struct {
}

func (s *Service) GetMessages(request *service.GetMessagesRequest) (*service.GetMessagesResponse, error) {
	//GetMessages
	panic("implement me")
}

func (s *Service) SendMessage(request *service.SendMessageRequest) error {
	//AddMessage
	//Notify
	panic("implement me")
}

func (s *Service) UpdateMessageStatus(request *service.UpdateMessageStatusRequest) error {
	//UpdateMessageStatus
	//Notify
	panic("implement me")
}

func (s *Service) AddChat(request *service.AddChatRequest) error {
	//AddChat
	//Notify
	panic("implement me")
}

func (s *Service) GetChats(request *service.GetChatsRequest) (*service.GetChatsResponse, error) {
	//GetChats
	panic("implement me")
}

func (s *Service) GetStatusInfo(request *service.GetStatusInfoRequest) (*service.GetStatusInfoResponse, error) {
	//GetStatusInfo
	panic("implement me")
}

func (s *Service) UpdateStatus(request *service.UpdateStatusRequest) error {
	// UpdateStatus
	// Notify
	panic("implement me")
}
