package category

import (
	"time"

	"github.com/google/uuid"
	category_grpc "github.com/kazmerdome/grpc-enricher/internal/module/category/category-grpc"
)

// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/module/category/category-grpc/category.proto

// Definitions
//

type CategoryEnricher interface {
	Enrich(loadEntity bool, category *category_grpc.Category, params *category_grpc.CategoryEnrichParams) (*category_grpc.Category, error)
}

type CategoryDataloader interface {
	LoadCategory(id uuid.UUID) (*Category, error)
}

// Entity & Enum
//

type Category struct {
	Id        uuid.UUID      `json:"id" bson:"_id"`
	Name      string         `json:"name" bson:"name"`
	Slug      string         `json:"slug" bson:"slug"`
	Status    CategoryStatus `json:"status" bson:"status"`
	CreatedAt time.Time      `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt" bson:"updatedAt"`
}

type CategoryStatus string

const (
	StatusActive   CategoryStatus = "ACTIVE"
	StatusPending  CategoryStatus = "PENDING"
	StatusArchived CategoryStatus = "ARCHIVED"
)
