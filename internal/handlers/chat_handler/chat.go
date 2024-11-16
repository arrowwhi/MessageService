package chat_handler

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "MessageService/proto/message_service/chats/v1"
)

func (s *Service) AddChat(ctx context.Context, req *v1.AddChatRequest) (*emptypb.Empty, error) {
	fmt.Println("implement me")
	return &emptypb.Empty{}, nil
}

func (s *Service) GetChats(context.Context, *v1.GetChatsRequest) (*v1.GetChatsResponse, error) {
	fmt.Println("implement me")
	return &v1.GetChatsResponse{}, nil
}
