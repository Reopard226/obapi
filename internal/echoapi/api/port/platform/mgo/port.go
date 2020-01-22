package mgo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
)

// COLLNAME Collection name
const COLLNAME = "ports"

// Port represents the client for Port table
type Port struct{}

// NewPort returns a new port database instance
func NewPort() *Port {
	return &Port{}
}

// View returns single port
func (u *Port) View(db *mongo.Database, portID, segment string) (*model.Port, error) {
	var port = new(model.Port)
	collection := db.Collection(COLLNAME)
	ctx := context.Background()

	err := collection.FindOne(ctx, bson.M{"_id": portID}).Decode(&port)
	return port, err
}

// List returns list of all ports
func (u *Port) List(db *mongo.Database, portID, segment string) ([]model.Port, error) {
	var ports []model.Port
	collection := db.Collection(COLLNAME)
	ctx := context.Background()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	err = cur.All(ctx, &ports)
	return ports, err
}
