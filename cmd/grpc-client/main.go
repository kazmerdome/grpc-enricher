package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	category_grpc "github.com/kazmerdome/grpc-enricher/internal/module/category/category-grpc"
	post_grpc "github.com/kazmerdome/grpc-enricher/internal/module/post/post-grpc"
	tag_grpc "github.com/kazmerdome/grpc-enricher/internal/module/tag/tag-grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Dial the gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := post_grpc.NewPostServiceClient(conn)

	t := true
	res, err := client.ListPost(context.Background(), &post_grpc.ListPostRequest{
		EnrichParams: &post_grpc.PostEnrichParams{
			// EnrichAllFields:    &t,
			// EnrichAllRelations: &t,
			Id:      &t,
			Content: &t,
			Tags: &tag_grpc.TagEnrichParams{
				Id:   &t,
				Name: &t,
			},
			Category: &category_grpc.CategoryEnrichParams{
				Name: &t,
				Tags: &tag_grpc.TagEnrichParams{
					Id:   &t,
					Name: &t,
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed to call ListPost: %v", err)
	}

	// Convert the struct to JSON with indentation
	jsonData, err := json.MarshalIndent(res.GetPosts(), "", "  ")
	if err != nil {
		log.Fatalf("Error converting struct to JSON: %v", err)
	}

	// Log the JSON-formatted struct
	fmt.Println(string(jsonData))
}
