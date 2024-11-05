package tag

import (
	"time"

	"github.com/google/uuid"
)

var TagData = []Tag{
	{
		Id:        uuid.New(),
		Name:      "Tag 1",
		Slug:      "tag-1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        uuid.New(),
		Name:      "Tag 2",
		Slug:      "tag-2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        uuid.New(),
		Name:      "Tag 3",
		Slug:      "tag-3",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        uuid.New(),
		Name:      "Tag 4",
		Slug:      "tag-4",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
