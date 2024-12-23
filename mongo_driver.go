package mongo_helper

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoHelper struct {
	client *mongo.Client
}

func (helper *MongoHelper) Aggregate(ctx context.Context, pipeline AggregationPipeline, result []any) (err error) {
	cursor, err := helper.client.Database(pipeline.Database).Collection(pipeline.Collection).Aggregate(ctx, pipeline.Pipeline, pipeline.Options...)
	if err != nil {
		return err
	}

	defer func() {
		err = cursor.Close(ctx)
		if err != nil {
			err = cursor.Err()
		}
	}()

	return cursor.All(ctx, &result)
}

func (helper *MongoHelper) CountDocuments(ctx context.Context, counter DocumentCounter) (int64, error) {
	return helper.client.Database(counter.Database).Collection(counter.Collection).CountDocuments(ctx, counter.Selector, counter.Options...)
}

func (helper *MongoHelper) DeleteMany(ctx context.Context, remover DocumentRemover) (int64, error) {
	result, err := helper.client.Database(remover.Database).Collection(remover.Collection).DeleteMany(ctx, remover.Selector, remover.Options...)
	if err != nil || result == nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (helper *MongoHelper) DeleteOne(ctx context.Context, remover DocumentRemover) error {
	_, err := helper.client.Database(remover.Database).Collection(remover.Collection).DeleteOne(ctx, remover.Selector, remover.Options...)
	return err
}

func (helper *MongoHelper) Distinct(ctx context.Context, finder DistinctFieldFinder) ([]any, error) {
	return helper.client.Database(finder.Database).Collection(finder.Collection).Distinct(ctx, finder.Field, finder.Selector, finder.Options...)
}

func (helper *MongoHelper) FindMany(ctx context.Context, finder MultiDocumentFinder, result []any) (err error) {
	cursor, err := helper.client.Database(finder.Database).Collection(finder.Collection).Find(ctx, finder.Selector, finder.Options...)
	if err != nil {
		return err
	}

	defer func() {
		err = cursor.Close(ctx)
		if err != nil {
			err = cursor.Err()
		}
	}()

	return cursor.All(ctx, &result)
}

func (helper *MongoHelper) FindOne(ctx context.Context, finder SingleDocumentFinder, result any) error {
	single := helper.client.Database(finder.Database).Collection(finder.Collection).FindOne(ctx, finder.Selector, finder.Options...)
	return single.Decode(&result)
}

func (helper *MongoHelper) InsertMany(ctx context.Context, inserter MultiDocumentInserter) (int64, error) {
	result, err := helper.client.Database(inserter.Database).Collection(inserter.Collection).InsertMany(ctx, inserter.Documents, inserter.Options...)
	if err != nil || result == nil {
		return 0, err
	}

	return int64(len(result.InsertedIDs)), nil
}

func (helper *MongoHelper) InsertOne(ctx context.Context, inserter SingleDocumentInserter) error {
	_, err := helper.client.Database(inserter.Database).Collection(inserter.Collection).InsertOne(ctx, inserter.Document, inserter.Options...)
	return err
}

func (helper *MongoHelper) UpdateMany(ctx context.Context, updater DocumentUpdater) (int64, error) {
	result, err := helper.client.Database(updater.Database).Collection(updater.Collection).UpdateMany(ctx, updater.Selector, bson.M{"$set": updater.Update}, updater.Options...)
	if err != nil || result == nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (helper *MongoHelper) UpdateOne(ctx context.Context, updater DocumentUpdater) error {
	_, err := helper.client.Database(updater.Database).Collection(updater.Collection).UpdateOne(ctx, updater.Selector, bson.M{"$set": updater.Update}, updater.Options...)
	return err
}

func (helper *MongoHelper) Close(ctx context.Context) {
	if helper.client != nil {
		err := helper.client.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}
}

func NewMongoHelper(ctx context.Context, uri string) (*MongoHelper, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &MongoHelper{client: client}, nil
}
