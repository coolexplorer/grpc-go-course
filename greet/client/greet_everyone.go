package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/coolexplorer/grpc-go-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked.")

	reqs := []*pb.GreetRequest{
		{FirstName: "Kim"},
		{FirstName: "Lee"},
		{FirstName: "Jeong"},
	}

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatal("Error while calling GreetEveryone: ", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receving: %v\n", err)
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
