package service

import (
	"context"
	"f-dorm/app/demo/internal/biz"
	"f-dorm/core/const/codename"
	"f-dorm/core/logger"

	pb "f-dorm/api/demo/v1"
)

type DemoService struct {
	pb.UnimplementedDemoServer

	uc *biz.GreeterUsecase
}

func NewDemoService(uc *biz.GreeterUsecase) *DemoService {
	return &DemoService{
		uc: uc,
	}
}

func (s *DemoService) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloReply, error) {
	ctx = logger.NewContextWithTraceID(ctx, codename.ServiceDemo+"-SayHello")
	message, err := s.uc.GenerateHelloMessage(ctx, req.UserName)
	if err != nil {
		return nil, err
	}
	return &pb.SayHelloReply{
		Message: message,
	}, nil
}
