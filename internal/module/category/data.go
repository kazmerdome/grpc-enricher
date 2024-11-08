package category

import (
	"time"

	"github.com/google/uuid"
	"github.com/kazmerdome/grpc-enricher/internal/module/tag"
)

var CategoryData = []Category{
	{
		Id:     uuid.New(),
		Name:   "Category 1",
		Slug:   "category-1",
		Status: StatusActive,
		Tags: []uuid.UUID{
			tag.TagData[0].Id,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:   uuid.New(),
		Name: "Category 2",
		Slug: "category-2",
		Tags: []uuid.UUID{
			tag.TagData[2].Id,
			tag.TagData[1].Id,
			tag.TagData[0].Id,
		},
		Status:    StatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        uuid.New(),
		Name:      "Category 3",
		Slug:      "category-3",
		Status:    StatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        uuid.New(),
		Name:      "Category 4",
		Slug:      "category-4",
		Status:    StatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
