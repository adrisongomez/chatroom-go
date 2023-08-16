package service

import (
	"context"
	"user-services/gen/proto"
)

type Service struct {
	proto.UnimplementedUserServiceServer
}

func New() *Service {
	return &Service{}
}

func (Service) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.User, error) {
	return &proto.User{
		Username: req.Username,
	}, nil
}
