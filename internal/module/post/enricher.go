package post

import (
	"github.com/google/uuid"
	"github.com/kazmerdome/grpc-enricher/internal/module/category"
	category_grpc "github.com/kazmerdome/grpc-enricher/internal/module/category/category-grpc"
	post_grpc "github.com/kazmerdome/grpc-enricher/internal/module/post/post-grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type postEnricher struct {
	categoryEnricher category.CategoryEnricher
	postDataloader   PostDataloader
}

func NewPostEnricher(categoryEnricher category.CategoryEnricher, postDataloader PostDataloader) *postEnricher {
	return &postEnricher{
		categoryEnricher: categoryEnricher,
		postDataloader:   postDataloader,
	}
}

func (r *postEnricher) Enrich(loadEntity bool, post *post_grpc.Post, params *post_grpc.PostEnrichParams) (*post_grpc.Post, error) {
	// If post is nil, return nil
	if post == nil {
		return nil, nil
	}

	// Handle self enrichment
	if loadEntity {
		// validate id
		postId, err := uuid.Parse(post.Id)
		if err != nil {
			return nil, err
		}
		postData, err := r.postDataloader.LoadPost(postId)
		if err != nil {
			return nil, err
		}

		if postData == nil {
			return nil, nil
		}

		createdAt := timestamppb.New(postData.CreatedAt)
		updatedAt := timestamppb.New(postData.UpdatedAt)
		post = &post_grpc.Post{
			Id:    postData.Id.String(),
			Title: postData.Title,
			Slug:  postData.Slug,
			Category: &category_grpc.Category{
				Id: postData.Category.String(),
			},
			Content:   postData.Content,
			Status:    1, // TODO - fix this
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
	}

	// If no params, return post
	if params == nil {
		return post, nil
	}

	// Enrich fields
	if !params.GetEnrichAllFields() {
		p := &post_grpc.Post{}
		if params.GetId() {
			p.Id = post.Id
		}
		if params.GetTitle() {
			p.Title = post.Title
		}
		if params.GetSlug() {
			p.Slug = post.Slug
		}
		if params.GetCategory() != nil {
			p.Category = post.Category
		}
		if params.GetStatus() {
			p.Status = post.Status
		}
		if params.GetContent() {
			p.Content = post.Content
		}
		if params.GetCreatedAt() {
			p.CreatedAt = post.CreatedAt
		}
		if params.GetUpdatedAt() {
			p.UpdatedAt = post.UpdatedAt
		}
		post = p
	}

	// Enrich relations
	//

	// Enrich category
	categoryEnrichParams := params.Category
	if params.GetEnrichAllRelations() {
		t := true
		categoryEnrichParams = &category_grpc.CategoryEnrichParams{
			EnrichAllFields: &t,
		}
		if params.GetEnrichAllFields() {
			categoryEnrichParams.EnrichAllFields = &t
		}
	}

	// Enrich category
	if categoryEnrichParams != nil {
		category, err := r.categoryEnricher.Enrich(true, post.Category, categoryEnrichParams)
		if err != nil {
			return nil, err
		}
		post.Category = category
	}

	return post, nil
}
