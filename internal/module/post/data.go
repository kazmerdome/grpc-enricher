package post

import (
	"time"

	"github.com/google/uuid"
	"github.com/kazmerdome/grpc-enricher/internal/module/category"
)

var PostData = []Post{
	{
		Id:        uuid.New(),
		Title:     "Post 1",
		Slug:      "post-1",
		Category:  category.CategoryData[0].Id,
		Status:    StatusActive,
		Content:   "Post 1 content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        uuid.New(),
		Title:     "Post 2",
		Slug:      "post-2",
		Category:  category.CategoryData[1].Id,
		Status:    StatusActive,
		Content:   "Post 2 content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        uuid.New(),
		Title:     "Post 3 w category 1",
		Slug:      "post-3-w-category-1",
		Category:  category.CategoryData[0].Id,
		Status:    StatusActive,
		Content:   "Post 3 content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
