package mongo_helper

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Location struct {
	Database   string
	Collection string
}

type LocationQuery struct {
	Location
	Selector bson.M
}

type AggregationPipeline struct {
	Location
	Pipeline bson.D
	Options  []*options.AggregateOptions
}

type DocumentCounter struct {
	LocationQuery
	Options []*options.CountOptions
}

type DocumentRemover struct {
	LocationQuery
	Options []*options.DeleteOptions
}

type SingleDocumentFinder struct {
	LocationQuery
	Options []*options.FindOneOptions
}

type MultiDocumentFinder struct {
	LocationQuery
	Options []*options.FindOptions
}

type SingleDocumentInserter struct {
	Location
	Document any
	Options  []*options.InsertOneOptions
}

type MultiDocumentInserter struct {
	Location
	Documents []any
	Options   []*options.InsertManyOptions
}

type DocumentUpdater struct {
	LocationQuery
	Update  any
	Options []*options.UpdateOptions
}

type DistinctFieldFinder struct {
	LocationQuery
	Field   string
	Options []*options.DistinctOptions
}
