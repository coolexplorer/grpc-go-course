package main

import (
	"context"
	"log"

	pb "github.com/coolexplorer/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoded with %s", in)
	return &pb.SumResponse{
		SumResult: in.Value1 + in.Value2,
	}, nil
}
