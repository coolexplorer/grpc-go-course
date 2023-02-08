package main

import (
	"io"
	"log"

	pb "github.com/coolexplorer/grpc-go-course/calculator/proto"
)

func (S *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg Function was invoked.")

	var sum int64 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				AvgResult: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatal("Error while reading client stream: ", err)
		}

		log.Printf("Receiving number: %d\n", req.Value)
		sum += req.Value
		count++
	}

}
