package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// CurrentCongestion represents CurrentCongestion model
type CurrentCongestion struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Imo              int                `json:"imo" bson:"imo"`
	Name             string             `json:"name" bson:"name"`
	Dwt              string             `json:"dwt" bson:"dwt"`
	Segment          string             `json:"segment" bson:"segment"`
	PortID           string             `json:"port_id" bson:"port_id"`
	PortName         string             `json:"port_name" bson:"port_name"`
	TimestampArrival string             `json:"timestamp_arrival" bson:"timestamp_arrival"`
	WaitingTime      int                `json:"waiting_time" bson:"waiting_time"`
}
