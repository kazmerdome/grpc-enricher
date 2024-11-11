package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/kazmerdome/grpc-enricher/internal/module/category"
	"github.com/kazmerdome/grpc-enricher/internal/module/post"
	post_grpc "github.com/kazmerdome/grpc-enricher/internal/module/post/post-grpc"
	"github.com/kazmerdome/grpc-enricher/internal/module/tag"
	"google.golang.org/grpc"
)

func main() {
	// Initialize Domain Modules
	//
	tagModule := tag.NewTagModule()
	categoryModule := category.NewCategoryModule(tagModule.GetTagEnricher())
	postModule := post.NewPostModule(categoryModule.GetCategoryEnricher(), tagModule.GetTagEnricher())

	// Set up a TCP listener on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the PostService with the gRPC server
	post_grpc.RegisterPostServiceServer(grpcServer, postModule.GetController())

	// Run our server in a goroutine so that it doesn't block listening for shutdown
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC server over port 50051: %v", err)
		}
	}()

	// Block until a shutdown signal received (CTRL+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Gracefully shut down gRPC server
	grpcServer.GracefulStop()

	// Close TCP listener
	err = listener.Close()
	if err != nil {
		return
	}
}
