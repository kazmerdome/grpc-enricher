package category

import (
	"github.com/google/uuid"
	category_grpc "github.com/kazmerdome/grpc-enricher/internal/module/category/category-grpc"
	"github.com/kazmerdome/grpc-enricher/internal/module/tag"
	tag_grpc "github.com/kazmerdome/grpc-enricher/internal/module/tag/tag-grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type categoryEnricher struct {
	tagEnricher        tag.TagEnricher
	categoryDataloader CategoryDataloader
}

func NewCategoryEnricher(tagEnricher tag.TagEnricher, categoryDataloader CategoryDataloader) *categoryEnricher {
	return &categoryEnricher{
		tagEnricher:        tagEnricher,
		categoryDataloader: categoryDataloader,
	}
}

func (r *categoryEnricher) Enrich(loadEntity bool, category *category_grpc.Category, params *category_grpc.CategoryEnrichParams) (*category_grpc.Category, error) {
	// If category is nil, return nil
	if category == nil {
		return nil, nil
	}

	// Handle self enrichment
	if loadEntity {
		// validate id
		categoryId, err := uuid.Parse(category.Id)
		if err != nil {
			return nil, err
		}
		categoryData, err := r.categoryDataloader.LoadCategory(categoryId)
		if err != nil {
			return nil, err
		}

		if categoryData == nil {
			return nil, nil
		}

		createdAt := timestamppb.New(categoryData.CreatedAt)
		updatedAt := timestamppb.New(categoryData.UpdatedAt)
		category = &category_grpc.Category{
			Id:        categoryData.Id.String(),
			Name:      categoryData.Name,
			Slug:      categoryData.Slug,
			Status:    1,
			Tags:      []*tag_grpc.Tag{},
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		for _, postTag := range categoryData.Tags {
			tagId := postTag.String()
			tag := &tag_grpc.Tag{
				Id: tagId,
			}
			category.Tags = append(category.Tags, tag)
		}
	}

	// If no params, return category
	if params == nil {
		return category, nil
	}

	// Enrich fields
	if !params.GetEnrichAllFields() {
		c := &category_grpc.Category{}
		if params.GetId() {
			c.Id = category.Id
		}
		if params.GetName() {
			c.Name = category.Name
		}
		if params.GetSlug() {
			c.Slug = category.Slug
		}
		if params.GetStatus() {
			c.Status = category.Status
		}
		if params.GetTags() != nil {
			c.Tags = category.Tags
		}
		if params.GetCreatedAt() {
			c.CreatedAt = category.CreatedAt
		}
		if params.GetUpdatedAt() {
			c.UpdatedAt = category.UpdatedAt
		}
		category = c
	}

	// Enrich relations
	//

	tagsEnrichParams := params.Tags
	if params.GetEnrichAllRelations() {
		t := true
		tagsEnrichParams = &tag_grpc.TagEnrichParams{
			EnrichAllRelations: &t,
		}
		if params.GetEnrichAllFields() {
			tagsEnrichParams.EnrichAllFields = &t
		}
	}
	if tagsEnrichParams != nil && len(category.Tags) > 0 {
		tagsResultChan := make(chan []*tag_grpc.Tag)
		tagsErrorChan := make(chan error)
		go func() {
			tags, err := r.tagEnricher.EnrichBulk(true, category.Tags, tagsEnrichParams)
			if err != nil {
				tagsErrorChan <- err
				return
			}
			tagsResultChan <- tags
		}()

		select {
		case tags := <-tagsResultChan:
			category.Tags = tags
		case err := <-tagsErrorChan:
			return nil, err
		}
	}

	return category, nil
}
