package main

import (
	"context"
	"github.com/google/uuid"
	"mongo-helper"
	"time"
)

const uri = "mongodb://root:example@localhost:27017"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	helper, err := mongo_helper.NewMongoHelper(ctx, uri)
	if err != nil {
		panic(err)
	}

	testUser := User{
		ID:   uuid.NewString(),
		Name: "John Doe",
	}

	sdi := mongo_helper.SingleDocumentInserter{
		Location: mongo_helper.Location{},
		Document: testUser,
	}
	err = helper.InsertOne(ctx, sdi)
	if err != nil {
		panic(err)
	}
}
