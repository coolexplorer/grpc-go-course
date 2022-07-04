package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/coolexplorer/grpc-go-course/calculator/proto"
)

var addr string = "localhost:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	doSum(c, 10, 20)
	doPrimes(c, 100)
}
