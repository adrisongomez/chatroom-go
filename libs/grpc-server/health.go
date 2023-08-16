package grpcserver

import (
	"context"
	"grpc-server/gen/proto"
)

type HealthServer struct {
	proto.UnimplementedHealthServer
}

func (s HealthServer) Check(_ context.Context, req *proto.Request) (*proto.Response, error) {
	return &proto.Response{
		Content: req.Content,
	}, nil
}
