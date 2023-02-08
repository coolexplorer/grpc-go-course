package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/coolexplorer/grpc-go-course/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet Function was invoked.")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatal("Error while reading client strem: ", err)
		}

		res += fmt.Sprintf("Hello, %s!\n", req.FirstName)
	}
}
