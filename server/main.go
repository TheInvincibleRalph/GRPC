package main

import (
	"log"
	"net"

	pb "github.com/theinvincible/grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct { //GreetServiceServer is the server API for GreetService service
	pb.GreetServiceServer //By embedding this interface, helloServer becomes a gRPC server that implements the methods defined in GreetServiceServer.
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
