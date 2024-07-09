package main

import (
	"context"
	"log"
	"time"

	pb "github.com/theinvincible/grpc/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, err := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
