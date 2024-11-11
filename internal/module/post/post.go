package post

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kazmerdome/grpc-enricher/internal/module/category"
	post_grpc "github.com/kazmerdome/grpc-enricher/internal/module/post/post-grpc"
)

// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/module/post/post-grpc/post.proto

// Definitions
//

type PostRepository interface {
	GetManyByIds(ctx context.Context, ids []uuid.UUID) ([]*Post, error)
	GetMany(ctx context.Context) ([]Post, error)
}

type PostDataloader interface {
	ItemLoader(ctx context.Context, id uuid.UUID) (*Post, error)
}

type PostResolver interface {
	GetCategory(id uuid.UUID) (category.Category, error)
}

type PostEnricher interface {
	Enrich(loadEntity bool, post *post_grpc.Post, params *post_grpc.PostEnrichParams) (*post_grpc.Post, error)
}

// Entity & Enum
//

type Post struct {
	Id        uuid.UUID   `json:"id" bson:"_id"`
	Title     string      `json:"title" bson:"title"`
	Slug      string      `json:"slug" bson:"slug"`
	Tags      []uuid.UUID `json:"tags" bson:"tags"`
	Category  uuid.UUID   `json:"category" bson:"category"`
	Status    PostStatus  `json:"status" bson:"status"`
	Content   string      `json:"content" bson:"content"`
	CreatedAt time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt" bson:"updatedAt"`
}

type PostStatus string

const (
	StatusActive   PostStatus = "ACTIVE"
	StatusPending  PostStatus = "PENDING"
	StatusArchived PostStatus = "ARCHIVED"
)
