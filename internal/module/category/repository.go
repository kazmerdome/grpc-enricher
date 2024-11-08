package category

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type categoryRepository struct {
}

func NewCategoryRepository() *categoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) GetManyByIds(ctx context.Context, ids []uuid.UUID) ([]*Category, error) {
	fmt.Println("categoryRepository.GetManyByIds called with ids:", ids)

	result := []*Category{}
	for _, id := range ids {
		var foundCategory *Category
		for i := range CategoryData {
			if CategoryData[i].Id == id {
				foundCategory = &CategoryData[i]
				break
			}
		}
		result = append(result, foundCategory)
	}
	return result, nil
}
