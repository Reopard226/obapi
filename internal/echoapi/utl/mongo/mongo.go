package mongo

import (
	"cloud.google.com/go/firestore"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// APIKEY_COLLECTION_NAME Collection name

type IamDAO struct {
	Ctx context.Context
	Db  *mongo.Database
	Fs  *firestore.Client
}

func NewMongoDatabase(connString string, database string) (*mongo.Database, error) {
	var dbInit *mongo.Database
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	// Collection types can be used to access the database
	dbInit = client.Database(database)

	return dbInit, nil

}