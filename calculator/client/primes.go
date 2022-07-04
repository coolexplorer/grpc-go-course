package main

import (
	"context"
	"io"
	"log"

	pb "github.com/coolexplorer/grpc-go-course/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient, num int32) {
	log.Println("doPrimes was invoked.")

	req := &pb.PrimeRequest{
		Number: num,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatal("Error while calling doPrimes: ", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Error while reading the stream: ", err)
		}

		log.Printf("Primes: %v\n", msg.Prime)
	}
}
