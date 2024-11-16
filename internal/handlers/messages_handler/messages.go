package messages_handler

import (
	pb "MessageService/proto/message_service/message/v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetMessages(ctx context.Context, request *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) SendMessage(ctx context.Context, request *pb.SendMessageRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateMessageStatus(ctx context.Context, request *pb.UpdateMessageStatusRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
