package main

import (
	"log"

	pb "github.com/coolexplorer/grpc-go-course/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %s", in)

	var k int32 = 2
	var n int32 = in.Number

	for {
		if n <= 1 {
			break
		}

		if n%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Prime: k,
			})
			n = n / k
		} else {
			k++
		}
	}

	return nil
}
