package model

// DryDockCount represents DryDockCount model
type DryDockCount struct {
	Base
	DateCol     string `json:"date_col" bson:"date_col"`
	VesselCount int    `json:"vessel_count" bson:"vessel_count"`
}

// DryDockFullrow represents DryDockFullrow model
type DryDockFullrow struct {
	Base
	DateCol          string  `json:"date_col" bson:"date_col"`
	Segment          string  `json:"segment" bson:"segment"`
	VesselCount      int     `json:"vessel_count" bson:"vessel_count"`
	VesselCountFleet float64 `json:"vessel_count_fleet" bson:"vessel_count_fleet"`
	Dwt              int     `json:"dwt" bson:"dwt"`
	DwtFleet         float64 `json:"dwt_fleet" bson:"dwt_fleet"`
}

// TimeseriesArray represents TimeseriesArray model
type TimeseriesArray struct {
	DateCol string  `json:"date_col" bson:"date_col"`
	Value   float64 `json:"value" bson:"value"`
}

// TimeseriesObject represents TimeseriesObject model
type TimeseriesObject struct {
	Date        []string  `json:"date"`
	Value       []float64 `json:"value"`
	Year        []int     `json:"year"`
	UnifiedDate []string  `json:"unified_date"`
}

// DryDockSummaryStatsFullRow represents DryDockSummaryStatsFullRow model
type DryDockSummaryStatsFullRow struct {
	Base
	Month              string `json:"month" bson:"month"`
	Segment            string `json:"segment" bson:"segment"`
	AverageDaysInDock  int    `json:"average_days_in_dock" bson:"average_days_in_dock"`
	CompletedDockStays int    `json:"completed_dock_stays" bson:"completed_dock_stays"`
}

// DryDockSummaryStatsObject represents DryDockSummaryStatsObject model
type DryDockSummaryStatsObject struct {
	Month              []string `json:"month" bson:"month"`
	Segment            string   `json:"segment" bson:"segment"`
	AverageDaysInDock  []int    `json:"average_days_in_dock" bson:"average_days_in_dock"`
	CompletedDockStays []int    `json:"completed_dock_stays" bson:"completed_dock_stays"`
}
