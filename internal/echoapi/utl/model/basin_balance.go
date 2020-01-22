package model

// BasinBalanceFullRow represents BasinBalanceFullRow model
type BasinBalanceFullRow struct {
	Base
	Basin       string `json:"basin" bson:"basin"`
	Segment     string `json:"segment" bson:"segment"`
	DateCol     string `json:"date_col" bson:"date_col"`
	VesselCount int    `json:"vessel_count" bson:"vessel_count"`
}

// BasinBalanceCount represents BasinBalanceCount model
type BasinBalanceCount struct {
	Base
	DateCol     string `json:"date_col" bson:"date_col"`
	VesselCount int    `json:"vessel_count" bson:"vessel_count"`
}

// TimeCountObject represents TimeCountObject model
type TimeCountObject struct {
	Date        []string `json:"date"`
	VesselCount []int    `json:"vessel_count"`
	Year        []int    `json:"year"`
	UnifiedDate []string `json:"unified_date"`
}
