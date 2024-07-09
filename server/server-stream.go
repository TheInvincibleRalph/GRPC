package main

import (
	pb "github.com/theinvincible/grpc/proto"

	"log"
)

func (s *helloServer) callSayHelloServerStreaming(req *pb.NamesList, steam pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request with names : %v", req.Names)
	for _, name := range req.Names {

		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
