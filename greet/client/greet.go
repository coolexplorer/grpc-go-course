package main

import (
	"context"
	"log"

	pb "github.com/coolexplorer/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked.")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Allen",
	})

	if err != nil {
		log.Fatal("Could not greet: ", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
