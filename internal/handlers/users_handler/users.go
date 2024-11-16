package users_handler

import (
	pb "MessageService/proto/message_service/users/v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetStatusInfo(ctx context.Context, request *pb.GetStatusInfoRequest) (*pb.GetStatusInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateStatus(ctx context.Context, request *pb.UpdateStatusRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
