package post

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type postRepository struct {
}

func NewPostRepository() *postRepository {
	return &postRepository{}
}

func (r *postRepository) GetMany(ctx context.Context) ([]Post, error) {
	fmt.Println("\033[33m postRepository.GetMany called with")
	return PostData, nil
}

func (r *postRepository) GetManyByIds(ctx context.Context, ids []uuid.UUID) ([]*Post, error) {
	fmt.Println("\033[33m postRepository.GetManyByIds called with ids:", ids, "\033[0m")

	result := []*Post{}
	for _, id := range ids {
		var foundPost *Post
		for i := range PostData {
			if PostData[i].Id == id {
				foundPost = &PostData[i]
				break
			}
		}
		result = append(result, foundPost)
	}
	return result, nil
}
