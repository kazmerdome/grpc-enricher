package tag

import (
	"context"
	"time"

	"github.com/google/uuid"
	tag_grpc "github.com/kazmerdome/grpc-enricher/internal/module/tag/tag-grpc"
)

// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/module/tag/tag-grpc/tag.proto

// Definitions
//

type TagRepository interface {
	GetManyByIds(ctx context.Context, ids []uuid.UUID) ([]*Tag, error)
}

type TagEnricher interface {
	Enrich(loadEntity bool, tag *tag_grpc.Tag, params *tag_grpc.TagEnrichParams) (*tag_grpc.Tag, error)
	EnrichBulk(loadEntity bool, tags []*tag_grpc.Tag, params *tag_grpc.TagEnrichParams) ([]*tag_grpc.Tag, error)
}

type TagDataloader interface {
	ItemLoader(ctx context.Context, id uuid.UUID) (*Tag, error)
	ItemsLoader(ctx context.Context, ids []uuid.UUID) ([]*Tag, error)
}

// Entity & Enum
//

type Tag struct {
	Id        uuid.UUID `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Slug      string    `json:"slug" bson:"slug"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
