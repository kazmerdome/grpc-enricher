package tag

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type tagRepository struct {
}

func NewTagRepository() *tagRepository {
	return &tagRepository{}
}

func (r *tagRepository) GetManyByIds(ctx context.Context, ids []uuid.UUID) ([]*Tag, error) {
	fmt.Println("\033[33m tagRepository.GetManyByIds called with ids:", ids, "\033[0m")

	result := []*Tag{}
	for _, id := range ids {
		var foundTag *Tag
		for i := range TagData {
			if TagData[i].Id == id {
				foundTag = &TagData[i]
				break
			}
		}
		result = append(result, foundTag)
	}
	return result, nil
}
