package mongo_helper

import (
	"context"
)

type Helper interface {
	Aggregate(ctx context.Context, pipeline AggregationPipeline, result []any) (err error)
	CountDocuments(ctx context.Context, counter DocumentCounter) (int64, error)
	DeleteMany(ctx context.Context, remover DocumentRemover) (int64, error)
	DeleteOne(ctx context.Context, remover DocumentRemover) error
	Distinct(ctx context.Context, finder DistinctFieldFinder) ([]any, error)
	FindMany(ctx context.Context, finder MultiDocumentFinder, result []any) (err error)
	FindOne(ctx context.Context, finder SingleDocumentFinder, result any) error
	InsertMany(ctx context.Context, inserter MultiDocumentInserter) (int64, error)
	InsertOne(ctx context.Context, inserter SingleDocumentInserter) error
	UpdateMany(ctx context.Context, updater DocumentUpdater) (int64, error)
	UpdateOne(ctx context.Context, updater DocumentUpdater) error
	Close(ctx context.Context)
}
