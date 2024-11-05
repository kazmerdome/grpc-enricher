package tag

import (
	"fmt"

	"github.com/google/uuid"
)

type tagDataloader struct {
}

func NewTagDataloader() *tagDataloader {
	return &tagDataloader{}
}

// Mock Data
func (r *tagDataloader) LoadTag(id uuid.UUID) (*Tag, error) {
	fmt.Println("LoadTag called with id:", id)

	for i, v := range TagData {
		if v.Id == id {
			return &TagData[i], nil
		}
	}
	return nil, nil
}

// Mock Data
func (r *tagDataloader) LoadTags(ids []uuid.UUID) ([]*Tag, error) {
	fmt.Println("LoadTags called with ids:", ids)
	tags := map[uuid.UUID]*Tag{}
	for i := range TagData {
		tags[TagData[i].Id] = &TagData[i]
	}
	var result []*Tag
	for _, v := range ids {
		result = append(result, tags[v])
	}
	return result, nil
}
