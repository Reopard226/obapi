package mgo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
)

// COLLNAME Collection name
const COLLNAME = "regions"

// Region represents the client for region table
type Region struct{}

// NewRegion returns a new region database instance
func NewRegion() *Region {
	return &Region{}
}

// View returns single region
func (u *Region) View(db *mongo.Database, regionID, segment string) (*model.Region, error) {
	region := new(model.Region)
	collection := db.Collection(COLLNAME)
	ctx := context.Background()

	objID, err := primitive.ObjectIDFromHex(regionID)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&region)
	return region, err
}

// List returns list of all regions
func (u *Region) List(db *mongo.Database, regionID, segment string) ([]model.Region, error) {
	var regions []model.Region
	collection := db.Collection(COLLNAME)
	ctx := context.Background()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &regions)
	return regions, err
}
