package main

import (
	"log"

	pb "github.com/coolexplorer/grpc-go-course/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %s", in)

	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Prime: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}

	return nil

	// var k int64 = 2
	// var n int64 = in.Number

	// for {
	// 	if n <= 1 {
	// 		break
	// 	}

	// 	if n%k == 0 {
	// 		stream.Send(&pb.PrimeResponse{
	// 			Prime: k,
	// 		})
	// 		n = n / k
	// 	} else {
	// 		k++
	// 	}
	// }

	// return nil
}
