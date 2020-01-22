package model

// Region represents region model
type Region struct {
	Base
	RegionName string `json:"region_name" bson:"region_name"`
}
