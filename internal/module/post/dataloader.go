package post

import (
	"fmt"

	"github.com/google/uuid"
)

type postDataloader struct {
}

func NewPostDataloader() *postDataloader {
	return &postDataloader{}
}

// Mock Data
func (r *postDataloader) LoadPost(id uuid.UUID) (*Post, error) {
	fmt.Println("LoadPost called with id:", id)
	for i, v := range PostData {
		if v.Id == id {
			return &PostData[i], nil
		}
	}
	return nil, nil
}
