package dao

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"oceanbolt.com/iamservice/rpc/iam"
)



// COLLNAME Collection name
const COLLNAME = "apikeys"

type IamDAO struct {
	Ctx context.Context
	Db *mongo.Database
}

func NewMongoDatabase(connString string, database string) (*mongo.Database, error) {
	var dbInit *mongo.Database
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		return nil,err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil,err
	}
	// Collection types can be used to access the database
	dbInit = client.Database(database)

	return dbInit,nil

}

func (m *IamDAO) ListKeys(user *iam.User) (*iam.UserKeys, error) {

	var structResult []*iam.UserKey

	collection := m.Db.Collection(COLLNAME)
	cur, err := collection.Find(context.TODO(), bson.M{"user_id": user.Auth0UserId})
	if err != nil {
		return nil, err
	}
	defer cur.Close(m.Ctx)

	err = cur.All(m.Ctx, &structResult)
	if err != nil {
		log.Fatal(err)
	}

	return &iam.UserKeys{NumberOfKeys: int64(len(structResult)), Keys: structResult},nil
}

func (m *IamDAO) InsertKey(key *iam.UserKey) error {

	collection := m.Db.Collection(COLLNAME)
	_, err := collection.InsertOne(m.Ctx, key)
	if err != nil {
		return err
	}
	return nil
}

func (m *IamDAO) DeleteKey(key *iam.DeleteKeyRequest) error {

	collection := m.Db.Collection(COLLNAME)
	resp, err := collection.DeleteOne(m.Ctx, bson.M{"apikey_id":key.ApikeyId,"user_id":key.UserId})
	if err != nil {
		return err
	}
	if resp.DeletedCount == 0 {
		return errors.New("No key exists with apikey_id '"+key.ApikeyId+"'")
	}
	return nil
}
