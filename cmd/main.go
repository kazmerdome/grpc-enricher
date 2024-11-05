package main

import (
	"context"
	"fmt"

	"github.com/kazmerdome/grpc-enricher/internal/module/category"
	category_grpc "github.com/kazmerdome/grpc-enricher/internal/module/category/category-grpc"
	"github.com/kazmerdome/grpc-enricher/internal/module/post"
	post_grpc "github.com/kazmerdome/grpc-enricher/internal/module/post/post-grpc"
	"github.com/kazmerdome/grpc-enricher/internal/module/tag"
	tag_grpc "github.com/kazmerdome/grpc-enricher/internal/module/tag/tag-grpc"
)

func main() {
	// Initialize Domain Modules
	//
	tagModule := tag.NewTagModule()
	categoryModule := category.NewCategoryModule(tagModule.GetTagEnricher())
	postModule := post.NewPostModule(categoryModule.GetCategoryEnricher(), tagModule.GetTagEnricher())

	t := true
	posts, err := postModule.GetController().ListPost(context.Background(), &post_grpc.ListPostRequest{
		EnrichParams: &post_grpc.PostEnrichParams{
			Id:      &t,
			Title:   &t,
			Content: &t,
			Tags: &tag_grpc.TagEnrichParams{
				Name: &t,
			},
			Category: &category_grpc.CategoryEnrichParams{
				Id:   &t,
				Name: &t,
				Tags: &tag_grpc.TagEnrichParams{
					Name: &t,
				},
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
