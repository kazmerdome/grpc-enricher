package category

import (
	"fmt"

	"github.com/google/uuid"
)

type categoryDataloader struct {
}

func NewCategoryDataloader() *categoryDataloader {
	return &categoryDataloader{}
}

// Mock Data
func (r *categoryDataloader) LoadCategory(id uuid.UUID) (*Category, error) {
	fmt.Println("LoadCategory called with id:", id)

	for i, v := range CategoryData {
		if v.Id == id {
			return &CategoryData[i], nil
		}
	}
	return nil, nil
}
