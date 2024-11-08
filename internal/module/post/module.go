package post

import (
	"github.com/kazmerdome/grpc-enricher/internal/module/category"
	post_grpc "github.com/kazmerdome/grpc-enricher/internal/module/post/post-grpc"
	"github.com/kazmerdome/grpc-enricher/internal/module/tag"
)

type postModule struct {
	service    PostService
	controller post_grpc.PostServiceServer
}

func NewPostModule(categoryEnricher category.CategoryEnricher, tagEnricher tag.TagEnricher) *postModule {
	service := NewPostService()
	postDataloader := NewPostDataloader()
	postEnricher := NewPostEnricher(categoryEnricher, tagEnricher, postDataloader)
	controller := NewPostController(service, postEnricher)
	return &postModule{service: service, controller: controller}
}

func (m *postModule) GetController() post_grpc.PostServiceServer {
	return m.controller
}

func (m *postModule) GetService() PostService {
	return m.service
}
