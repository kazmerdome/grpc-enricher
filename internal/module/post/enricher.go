package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/kazmerdome/grpc-enricher/internal/module/category"
	category_grpc "github.com/kazmerdome/grpc-enricher/internal/module/category/category-grpc"
	post_grpc "github.com/kazmerdome/grpc-enricher/internal/module/post/post-grpc"
	"github.com/kazmerdome/grpc-enricher/internal/module/tag"
	tag_grpc "github.com/kazmerdome/grpc-enricher/internal/module/tag/tag-grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type postEnricher struct {
	categoryEnricher category.CategoryEnricher
	tagEnricher      tag.TagEnricher
	postDataloader   PostDataloader
}

func NewPostEnricher(
	categoryEnricher category.CategoryEnricher,
	tagEnricher tag.TagEnricher,
	postDataloader PostDataloader,
) *postEnricher {
	return &postEnricher{
		categoryEnricher: categoryEnricher,
		tagEnricher:      tagEnricher,
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
		postData, err := r.postDataloader.ItemLoader(context.Background(), postId)
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
			Tags:  []*tag_grpc.Tag{},
			Category: &category_grpc.Category{
				Id: postData.Category.String(),
			},
			Content:   postData.Content,
			Status:    1, // TODO - fix this
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		for _, postTag := range postData.Tags {
			tagId := postTag.String()
			tag := &tag_grpc.Tag{
				Id: tagId,
			}
			post.Tags = append(post.Tags, tag)
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
		if params.GetTags() != nil {
			p.Tags = post.Tags
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

	tagsEnrichParams := params.Tags
	categoryEnrichParams := params.Category
	if params.GetEnrichAllRelations() {
		t := true
		categoryEnrichParams = &category_grpc.CategoryEnrichParams{
			EnrichAllRelations: &t,
		}
		tagsEnrichParams = &tag_grpc.TagEnrichParams{
			EnrichAllRelations: &t,
		}
		if params.GetEnrichAllFields() {
			categoryEnrichParams.EnrichAllFields = &t
			tagsEnrichParams.EnrichAllFields = &t
		}
	}
	if categoryEnrichParams != nil {
		categoryResultChan := make(chan *category_grpc.Category)
		categoryErrorChan := make(chan error)
		go func() {
			category, err := r.categoryEnricher.Enrich(true, post.Category, categoryEnrichParams)
			if err != nil {
				categoryErrorChan <- err
				return
			}
			categoryResultChan <- category
		}()

		select {
		case category := <-categoryResultChan:
			post.Category = category
		case err := <-categoryErrorChan:
			return nil, err
		}
	}
	if tagsEnrichParams != nil && len(post.Tags) > 0 {
		tagsResultChan := make(chan []*tag_grpc.Tag)
		tagsErrorChan := make(chan error)
		go func() {
			tags, err := r.tagEnricher.EnrichBulk(true, post.Tags, tagsEnrichParams)
			if err != nil {
				tagsErrorChan <- err
				return
			}
			tagsResultChan <- tags
		}()

		select {
		case tags := <-tagsResultChan:
			post.Tags = tags
		case err := <-tagsErrorChan:
			return nil, err
		}
	}

	return post, nil
}
