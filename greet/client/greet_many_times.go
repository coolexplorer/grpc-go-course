package main

import (
	"context"
	"io"
	"log"

	pb "github.com/coolexplorer/grpc-go-course/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked.")

	req := &pb.GreetRequest{
		FirstName: "Client",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatal("Error while calling GreetManyTimes: ", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Error while reading the stream: ", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
