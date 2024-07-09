package main

import (
	"io"
	"log"

	pb "github.com/theinvincible/grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv() //request made from the client will be received and possible errors will be gracefully handled.
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)

		res := &pb.HelloResponse{ //response as defined in the greet.proto file with &pb.HelloResponse entering into the memory location of HelloResponse within the greet.pb.go file (the messaging file) and initializing it directly.
			Message: "Hello" + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
