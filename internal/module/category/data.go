package category

import (
	"time"

	"github.com/google/uuid"
)

var CategoryData = []Category{
	{
		Id:        uuid.New(),
		Name:      "Category 1",
		Slug:      "category-1",
		Status:    StatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        uuid.New(),
		Name:      "Category 2",
		Slug:      "category-2",
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
