package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Company represents company model
type Company struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Active    bool               `json:"active"`
	Locations []Location         `json:"locations,omitempty"`
	Owner     User               `json:"owner"`
}
