package main

import (
	"context"
	pb "github.com/theinvincible/grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
