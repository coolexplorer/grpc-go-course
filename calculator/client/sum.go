package main

import (
	"context"
	"log"

	pb "github.com/coolexplorer/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient, value1 int32, value2 int32) {
	log.Println("dosum was invoked.")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Value1: value1, Value2: value2,
	})

	if err != nil {
		log.Fatal("Could not sum: ", err)
	}

	log.Printf("Sum: %d\n", res.SumResult)
}
