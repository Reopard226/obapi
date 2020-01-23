package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Port represents port model
type Port struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PortName    string             `json:"port_name" bson:"port_name"`
	CountryCode string             `json:"country_code" bson:"country_code"`
	Unlocode    string             `json:"unlocode" bson:"unlocode"`
}
