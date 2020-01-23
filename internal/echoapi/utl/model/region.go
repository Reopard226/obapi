package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Region represents region model
type Region struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	RegionName string             `json:"region_name" bson:"region_name"`
}
