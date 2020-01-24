package mgo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
)

// Congestion represents the client for Congestion table
type Congestion struct{}

// NewCongestion returns a new NewCongestion database instance
func NewCongestion() *Congestion {
	return &Congestion{}
}

// CongestionPort returns CongestionPort
func (u *Congestion) CongestionPort(db *mongo.Database, portID, segment string) (*model.AnchoragePortArray, error) {
	var arrAncPort []model.AnchoragePort
	collection := db.Collection(model.ColAnchoragePort)
	ctx := context.Background()

	cur, err := collection.Find(ctx, bson.M{"port_id": portID, "segment": segment})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &arrAncPort)
	return model.ConvertAnchorageOutput(arrAncPort), err
}

// CongestionRegion returns CongestionRegion
func (u *Congestion) CongestionRegion(db *mongo.Database, regionID, segment string) (*model.AnchorageRegionArray, error) {
	var arrAncRegion []model.AnchorageRegion
	collection := db.Collection(model.ColAnchorageRegion)
	ctx := context.Background()

	cur, err := collection.Find(ctx, bson.M{"region_id": regionID, "segment": segment})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &arrAncRegion)
	return model.ConvertAnchorageRegionOutput(arrAncRegion), err
}

// ListPort returns list of all ports
func (u *Congestion) ListPort(db *mongo.Database) ([]model.Port, error) {
	var ports []model.Port
	collection := db.Collection(model.ColPort)
	ctx := context.Background()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &ports)
	return ports, err
}

// ListRegion returns list of all regions
func (u *Congestion) ListRegion(db *mongo.Database) ([]model.Region, error) {
	var regions []model.Region
	collection := db.Collection(model.ColRegion)
	ctx := context.Background()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &regions)
	return regions, err
}
