package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"mongo-helper"
)

type User struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

var usersLocation = mongo_helper.Location{
	Database:   "tests",
	Collection: "users",
}

func InsertUser(ctx context.Context, client mongo_helper.Helper, user User) error {
	sdi := mongo_helper.SingleDocumentInserter{
		Location: usersLocation,
		Document: user,
	}
	return client.InsertOne(ctx, sdi)
}

func FindUser(ctx context.Context, client mongo_helper.Helper, id string) (User, error) {
	var user User
	sdf := mongo_helper.SingleDocumentFinder{
		LocationQuery: mongo_helper.LocationQuery{
			Location: usersLocation,
			Selector: bson.M{"id": id},
		},
	}
	err := client.FindOne(ctx, sdf, &user)
	return user, err
}

func UpdateUser(ctx context.Context, client mongo_helper.Helper, id string, user User) error {
	du := mongo_helper.DocumentUpdater{
		LocationQuery: mongo_helper.LocationQuery{
			Location: usersLocation,
			Selector: bson.M{"id": id},
		},
		Update: user,
	}
	return client.UpdateOne(ctx, du)
}
