package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Location represents company location model
type Location struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name"`
	Active  bool               `json:"active"`
	Address string             `json:"address"`

	CompanyID int `json:"company_id"`
}
