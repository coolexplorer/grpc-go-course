package main

import (
	"log"
	"net"

	pb "github.com/coolexplorer/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50052"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal("Failed to listen on: ", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
