package post

import (
	"context"

	category_grpc "github.com/kazmerdome/grpc-enricher/internal/module/category/category-grpc"
	post_grpc "github.com/kazmerdome/grpc-enricher/internal/module/post/post-grpc"
	tag_grpc "github.com/kazmerdome/grpc-enricher/internal/module/tag/tag-grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// implements post_grpc.PostServiceServer

type postController struct {
	post_grpc.UnimplementedPostServiceServer
	service      PostService
	postEnricher PostEnricher
}

func NewPostController(service PostService, postEnricher PostEnricher) *postController {
	return &postController{
		service:      service,
		postEnricher: postEnricher,
	}
}

func (r *postController) ListPost(ctx context.Context, in *post_grpc.ListPostRequest) (*post_grpc.ListPostResponse, error) {
	// Get list of posts
	posts, err := r.service.ListPosts()
	if err != nil {
		return nil, err
	}

	postsResponse := []*post_grpc.Post{}
	for _, post := range posts {
		createdAt := timestamppb.New(post.CreatedAt)
		updatedAt := timestamppb.New(post.UpdatedAt)
		grpcPost := &post_grpc.Post{
			Id:    post.Id.String(),
			Title: post.Title,
			Slug:  post.Slug,
			Category: &category_grpc.Category{
				Id: post.Category.String(),
			},
			Content:   post.Content,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		for i := range post.Tags {
			grpcPost.Tags = append(grpcPost.Tags, &tag_grpc.Tag{
				Id: post.Tags[i].String(),
			})
		}

		// Enrich post
		grpcPostResultChan := make(chan *post_grpc.Post)
		grpcPostErrChan := make(chan error)

		go func() {
			grpcPost, err := r.postEnricher.Enrich(true, grpcPost, in.GetEnrichParams())
			if err != nil {
				grpcPostErrChan <- err
				return
			}
			grpcPostResultChan <- grpcPost
		}()

		select {
		case grpcPost = <-grpcPostResultChan:
		case err = <-grpcPostErrChan:
			return nil, err
		}

		postsResponse = append(postsResponse, grpcPost)
	}

	return &post_grpc.ListPostResponse{
		Posts: postsResponse,
	}, nil
}
