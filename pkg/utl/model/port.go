package model

// Port represents port model
type Port struct {
	Base
	PortName    string `json:"port_name" bson:"port_name"`
	CountryCode string `json:"country_code" bson:"country_code"`
	Unlocode    string `json:"unlocode" bson:"unlocode"`
}
