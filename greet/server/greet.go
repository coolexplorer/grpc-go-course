package main

import (
	"context"
	"log"

	pb "github.com/coolexplorer/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet fucntion was invoked with %s", in)
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
