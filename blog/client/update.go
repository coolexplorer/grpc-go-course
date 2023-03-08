package main

import (
	"context"
	"log"

	pb "github.com/coolexplorer/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Allen",
		Title:    "New Title",
		Content:  "Content of the first blog, with soem awesom additions!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happende while updating: %v\n", err)
	}

	log.Println("Blog was updated!")
}
