package category

import (
	"github.com/google/uuid"
	category_grpc "github.com/kazmerdome/grpc-enricher/internal/module/category/category-grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type categoryEnricher struct {
	categoryDataloader CategoryDataloader
}

func NewCategoryEnricher(categoryDataloader CategoryDataloader) *categoryEnricher {
	return &categoryEnricher{
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
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
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

	if params.GetEnrichAllRelations() {
		// TODO
	}

	return category, nil
}
