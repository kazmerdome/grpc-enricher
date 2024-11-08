package tag

import (
	"context"

	"github.com/google/uuid"
	tag_grpc "github.com/kazmerdome/grpc-enricher/internal/module/tag/tag-grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type tagEnricher struct {
	tagDataloader TagDataloader
}

func NewTagEnricher(tagDataloader TagDataloader) *tagEnricher {
	return &tagEnricher{
		tagDataloader: tagDataloader,
	}
}

func (r *tagEnricher) Enrich(loadEntity bool, tag *tag_grpc.Tag, params *tag_grpc.TagEnrichParams) (*tag_grpc.Tag, error) {
	// If tag is nil, return nil
	if tag == nil {
		return nil, nil
	}

	// Handle self enrichment
	if loadEntity {
		// validate id
		tagId, err := uuid.Parse(tag.Id)
		if err != nil {
			return nil, err
		}
		tagData, err := r.tagDataloader.ItemLoader(context.Background(), tagId)
		if err != nil {
			return nil, err
		}

		if tagData == nil {
			return nil, nil
		}

		createdAt := timestamppb.New(tagData.CreatedAt)
		updatedAt := timestamppb.New(tagData.UpdatedAt)
		tag = &tag_grpc.Tag{
			Id:        tagData.Id.String(),
			Name:      tagData.Name,
			Slug:      tagData.Slug,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
	}

	// If no params, return category
	if params == nil {
		return tag, nil
	}

	// Enrich fields
	if !params.GetEnrichAllFields() {
		t := &tag_grpc.Tag{}
		if params.GetId() {
			t.Id = tag.Id
		}
		if params.GetName() {
			t.Name = tag.Name
		}
		if params.GetSlug() {
			t.Slug = tag.Slug
		}
		if params.GetCreatedAt() {
			t.CreatedAt = tag.CreatedAt
		}
		if params.GetUpdatedAt() {
			t.UpdatedAt = tag.UpdatedAt
		}
		tag = t
	}

	return tag, nil
}

func (r *tagEnricher) EnrichBulk(loadEntity bool, tags []*tag_grpc.Tag, params *tag_grpc.TagEnrichParams) ([]*tag_grpc.Tag, error) {
	// If tags is nil, return nil
	if tags == nil {
		return nil, nil
	}

	// Handle self enrichment
	if loadEntity {
		ids := []uuid.UUID{}
		for _, tag := range tags {
			tagId, err := uuid.Parse(tag.Id)
			if err != nil {
				return nil, err
			}
			ids = append(ids, tagId)
		}

		tagsData, err := r.tagDataloader.ItemsLoader(context.TODO(), ids)
		if err != nil {
			return nil, err
		}

		if tagsData == nil {
			return nil, nil
		}

		for i := range tags {
			if tags[i] != nil {
				for _, t := range tagsData {
					if t.Id.String() == tags[i].Id {
						createdAt := timestamppb.New(t.CreatedAt)
						updatedAt := timestamppb.New(t.UpdatedAt)
						tags[i] = &tag_grpc.Tag{
							Id:        t.Id.String(),
							Name:      t.Name,
							Slug:      t.Slug,
							CreatedAt: createdAt,
							UpdatedAt: updatedAt,
						}
					}
				}
			}
		}
	}

	for i := range tags {
		tag, err := r.Enrich(false, tags[i], params)
		if err != nil {
			return nil, err
		}
		tags[i] = tag
	}

	return tags, nil
}
