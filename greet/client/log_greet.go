package main

import (
	"context"
	"log"
	"time"

	pb "github.com/coolexplorer/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Kim"},
		{FirstName: "Lee"},
		{FirstName: "Jeong"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatal("Error while calling LongGreet ", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal("Error while receiving response from LongGreet: ", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
