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

type MgoDao struct {
	Ctx context.Context
	Db  *mongo.Database
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

func (m *MgoDao) ListKeys(user *iam.User) (*iam.UserKeys, error) {

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

	return &iam.UserKeys{NumberOfKeys: int64(len(structResult)), Keys: structResult}, nil
}

func (m *MgoDao) InsertKey(key *iam.UserKey) error {

	collection := m.Db.Collection(COLLNAME)
	_, err := collection.InsertOne(m.Ctx, key)
	if err != nil {
		return err
	}
	return nil
}

func (m *MgoDao) DeleteKey(key *iam.DeleteKeyRequest) error {

	collection := m.Db.Collection(COLLNAME)
	resp, err := collection.DeleteOne(m.Ctx, bson.M{"apikey_id": key.ApikeyId, "user_id": key.UserId})
	if err != nil {
		return err
	}
	if resp.DeletedCount == 0 {
		return errors.New("No key exists with apikey_id '" + key.ApikeyId + "'")
	}
	return nil
}

func (m *MgoDao) CheckKey(key *iam.UserKey) (bool, error) {
	log.Println("Checking key...")

	collection := m.Db.Collection(COLLNAME)
	res := collection.FindOne(m.Ctx, bson.M{"apikey_id": key.ApikeyId, "user_id": key.UserId})
	if res.Err() != nil {
		if res.Err().Error() == "mongo: no documents in result" {
			return false, nil
		} else {
			return false, res.Err()
		}
	}

	return true, nil
}
