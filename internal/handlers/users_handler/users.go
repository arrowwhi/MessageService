package users_handler

import (
	pb "MessageService/proto/message_service/users/v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetStatusInfo(ctx context.Context, request *pb.GetStatusInfoRequest) (*pb.GetStatusInfoResponse, error) {
	info, err := s.service.GetStatusInfo(s.cvt.GetStatusInfoToService(request))
	if err != nil {
		return nil, err
	}

	return s.cvt.GetStatusInfoToHandler(info), nil
}

func (s *Service) UpdateStatus(ctx context.Context, request *pb.UpdateStatusRequest) (*emptypb.Empty, error) {
	err := s.service.UpdateStatus(s.cvt.UpdateStatusToService(request))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
