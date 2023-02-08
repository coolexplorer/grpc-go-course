package main

import (
	"context"
	"log"

	pb "github.com/coolexplorer/grpc-go-course/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked.")

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatal("Error while openling the stream: ", err)
	}

	numbers := []int64{3, 5, 9, 54, 23}

	for _, number := range numbers {
		log.Printf("Sending number: %d\n", number)

		stream.Send(&pb.AvgRequest{
			Value: number,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal("Error while receiving response: ", err)
	}

	log.Printf("Avg: %f\n", res.AvgResult)
}
