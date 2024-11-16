package chat_handler

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "MessageService/proto/message_service/chats/v1"
)

func (s *Service) AddChat(ctx context.Context, req *v1.AddChatRequest) (*emptypb.Empty, error) {
	fmt.Println("implement me")
	err := s.service.AddChat(s.cvt.AddChatToService(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Service) GetChats(ctx context.Context, req *v1.GetChatsRequest) (*v1.GetChatsResponse, error) {
	chats, err := s.service.GetChats(s.cvt.GetChatsToService(req))
	if err != nil {
		return nil, err
	}
	return s.cvt.GetChatsToHandler(chats), nil
}
