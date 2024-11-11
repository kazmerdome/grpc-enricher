package post

import (
	"github.com/kazmerdome/grpc-enricher/internal/module/category"
	post_grpc "github.com/kazmerdome/grpc-enricher/internal/module/post/post-grpc"
	"github.com/kazmerdome/grpc-enricher/internal/module/tag"
)

type postModule struct {
	controller post_grpc.PostServiceServer
}

func NewPostModule(categoryEnricher category.CategoryEnricher, tagEnricher tag.TagEnricher) *postModule {
	repository := NewPostRepository()
	postDataloader := NewPostDataloader(repository)
	postEnricher := NewPostEnricher(categoryEnricher, tagEnricher, postDataloader)
	controller := NewPostController(repository, postEnricher)
	return &postModule{controller: controller}
}

func (m *postModule) GetController() post_grpc.PostServiceServer {
	return m.controller
}
