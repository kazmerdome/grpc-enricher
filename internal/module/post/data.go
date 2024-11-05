package post

import (
	"time"

	"github.com/google/uuid"
	"github.com/kazmerdome/grpc-enricher/internal/module/category"
	"github.com/kazmerdome/grpc-enricher/internal/module/tag"
)

var PostData = []Post{
	{
		Id:    uuid.New(),
		Title: "Post 1",
		Slug:  "post-1",
		Tags: []uuid.UUID{
			tag.TagData[0].Id,
			tag.TagData[1].Id,
			tag.TagData[3].Id,
		},
		Category:  category.CategoryData[0].Id,
		Status:    StatusActive,
		Content:   "Post 1 content w tag 1, 2, 4",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        uuid.New(),
		Title:     "Post 2",
		Slug:      "post-2",
		Tags:      []uuid.UUID{},
		Category:  category.CategoryData[1].Id,
		Status:    StatusActive,
		Content:   "Post 2 content w no tags",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:    uuid.New(),
		Title: "Post 3 w category 1",
		Slug:  "post-3-w-category-1",
		Tags: []uuid.UUID{
			tag.TagData[3].Id,
			tag.TagData[1].Id,
		},
		Category:  category.CategoryData[0].Id,
		Status:    StatusActive,
		Content:   "Post 3 content w tag 4, 2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
