package tag

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/graph-gophers/dataloader"
	dl "github.com/kazmerdome/grpc-enricher/internal/util/dataloader"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type tagDataloader struct {
	repository  TagRepository
	itemLoader  *dataloader.Loader
	itemsLoader *dataloader.Loader
	logger      zerolog.Logger
}

func NewTagDataloader(repository TagRepository) *tagDataloader {
	loader := &tagDataloader{
		repository: repository,
		logger: log.
			With().
			Str("module", "tag").
			Str("provider", "dataloader").
			Logger(),
	}
	loader.itemLoader = dataloader.NewBatchedLoader(
		loader.batchItemLoader,
		dataloader.WithCache(&dataloader.NoCache{}),
	)
	loader.itemsLoader = dataloader.NewBatchedLoader(
		loader.batchItemsLoader,
		dataloader.WithCache(&dataloader.NoCache{}),
	)
	return loader
}

func (r *tagDataloader) ItemLoader(ctx context.Context, id uuid.UUID) (*Tag, error) {
	fmt.Println("\033[33m tagDataloader.ItemLoader called with id:", id, "\033[0m")

	thunk := r.itemLoader.Load(ctx, dl.UuidKey(id))
	result, err := thunk()
	if err != nil {
		r.logger.
			Debug().
			Err(err).
			Str("method", "ItemLoader").
			Str("tag", "call itemLoader.Load.thunk").
			Str("id", id.String()).
			Send()
		return nil, err
	}
	org, ok := result.(*Tag)
	if !ok {
		return nil, nil
	}
	return org, nil
}

func (r *tagDataloader) batchItemLoader(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// convert keys to uuids and create placeholders
	uuids := make([]uuid.UUID, len(keys))
	bucket := make(map[uuid.UUID]*dataloader.Result, len(keys))
	for i, key := range keys {
		uid := key.Raw().(uuid.UUID)
		uuids[i] = uid
		bucket[uid] = &dataloader.Result{Data: nil, Error: nil}
	}
	// call repository and add the values to the bucket
	orgs, err := r.repository.GetManyByIds(ctx, uuids)
	if err != nil {
		r.logger.
			Debug().
			Err(err).
			Str("method", "batchItemLoader").
			Str("tag", "repository.GetManyByIds call is failed").
			Send()

		return []*dataloader.Result{{Data: nil, Error: err}}
	}
	for _, org := range orgs {
		if org != nil {
			bucket[org.Id] = &dataloader.Result{Data: org, Error: nil}
		}
	}
	// create result array
	results := make([]*dataloader.Result, len(keys))
	for i, key := range keys {
		uid := key.Raw().(uuid.UUID)
		results[i] = bucket[uid]
	}
	return results
}

func (r *tagDataloader) ItemsLoader(ctx context.Context, ids []uuid.UUID) ([]*Tag, error) {
	fmt.Println("tagDataloader.ItemsLoader called with ids:", ids)

	uuidKeys := []dataloader.Key{}
	for _, id := range ids {
		key := dl.UuidKey(id)
		uuidKeys = append(uuidKeys, key)
	}

	thunk := r.itemsLoader.LoadMany(ctx, uuidKeys)
	resultsAny, errs := thunk()
	for _, err := range errs {
		r.logger.
			Debug().
			Err(err).
			Str("method", "ItemLoader").
			Str("tag", "call itemLoader.Load.thunk").
			Send()
		return nil, err
	}

	results := []*Tag{}
	for _, result := range resultsAny {
		var res *Tag
		r, ok := result.(*Tag)
		if ok {
			res = r
		}
		results = append(results, res)
	}
	return results, nil
}

func (r *tagDataloader) batchItemsLoader(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// convert keys to uuids and create placeholders
	uuids := make([]uuid.UUID, len(keys))
	bucket := make(map[uuid.UUID]*dataloader.Result, len(keys))
	for i, key := range keys {
		uid := key.Raw().(uuid.UUID)
		uuids[i] = uid
		bucket[uid] = &dataloader.Result{Data: nil, Error: nil}
	}
	// call repository and add the values to the bucket
	orgs, err := r.repository.GetManyByIds(ctx, uuids)
	if err != nil {
		r.logger.
			Debug().
			Err(err).
			Str("method", "batchItemLoader").
			Str("tag", "repository.GetManyByIds call is failed").
			Send()

		return []*dataloader.Result{{Data: nil, Error: err}}
	}
	for _, org := range orgs {
		if org != nil {
			bucket[org.Id] = &dataloader.Result{Data: org, Error: nil}
		}
	}
	// create result array
	results := make([]*dataloader.Result, len(keys))
	for i, key := range keys {
		uid := key.Raw().(uuid.UUID)
		results[i] = bucket[uid]
	}
	return results
}
