package messages_handler

import (
	pb "MessageService/proto/message_service/message/v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetMessages(ctx context.Context, request *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	messages, err := s.service.GetMessages(s.cvt.GetMessagesToService(request))
	if err != nil {
		return nil, err
	}
	return s.cvt.GetMessagesToHandler(messages), nil
}

func (s *Service) SendMessage(ctx context.Context, request *pb.SendMessageRequest) (*emptypb.Empty, error) {
	err := s.service.SendMessage(s.cvt.SendMessageToService(request))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Service) UpdateMessageStatus(ctx context.Context, request *pb.UpdateMessageStatusRequest) (*emptypb.Empty, error) {
	err := s.service.UpdateMessageStatus(s.cvt.UpdateMessageStatusToService(request))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
