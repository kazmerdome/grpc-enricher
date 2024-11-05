package main

import (
	"context"
	"fmt"

	"github.com/kazmerdome/grpc-enricher/internal/module/category"
	category_grpc "github.com/kazmerdome/grpc-enricher/internal/module/category/category-grpc"
	"github.com/kazmerdome/grpc-enricher/internal/module/post"
	post_grpc "github.com/kazmerdome/grpc-enricher/internal/module/post/post-grpc"
)

func main() {
	// Initialize Domain Modules
	//
	categoryModule := category.NewCategoryModule()
	postModule := post.NewPostModule(categoryModule.GetCategoryEnricher())

	t := true
	posts, err := postModule.GetController().ListPost(context.Background(), &post_grpc.ListPostRequest{
		EnrichParams: &post_grpc.PostEnrichParams{
			Id:    &t,
			Title: &t,
			Category: &category_grpc.CategoryEnrichParams{
				Id:   &t,
				Name: &t,
			},
		},
	})

	if err != nil {
		panic(err)
	}

	for _, post := range posts.Posts {
		fmt.Println(post)
	}
}
