package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DryDockCount represents DryDockCount model
type DryDockCount struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	DateCol     string             `json:"date_col" bson:"date_col"`
	VesselCount int                `json:"vessel_count" bson:"vessel_count"`
}

// DryDockFullRow represents DryDockFullrow model
type DryDockFullRow struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	DateCol          string             `json:"date_col" bson:"date_col"`
	Segment          string             `json:"segment" bson:"segment"`
	VesselCount      int                `json:"vessel_count" bson:"vessel_count"`
	VesselCountFleet float64            `json:"vessel_count_fleet" bson:"vessel_count_fleet"`
	Dwt              int                `json:"dwt" bson:"dwt"`
	DwtFleet         float64            `json:"dwt_fleet" bson:"dwt_fleet"`
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
	ID                 primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Month              string             `json:"month" bson:"month"`
	Segment            string             `json:"segment" bson:"segment"`
	AverageDaysInDock  int                `json:"average_days_in_dock" bson:"average_days_in_dock"`
	CompletedDockStays int                `json:"completed_dock_stays" bson:"completed_dock_stays"`
}

// DryDockSummaryStatsObject represents DryDockSummaryStatsObject model
type DryDockSummaryStatsObject struct {
	Month              []string `json:"month" bson:"month"`
	Segment            string   `json:"segment" bson:"segment"`
	AverageDaysInDock  []int    `json:"average_days_in_dock" bson:"average_days_in_dock"`
	CompletedDockStays []int    `json:"completed_dock_stays" bson:"completed_dock_stays"`
}

// ConvertDryDockTimeseries converts drydoc time series
func ConvertDryDockTimeseries(rows []DryDockFullRow, metric string, absolute bool) *TimeseriesObject {
	result := &TimeseriesObject{
		Date:        []string{},
		Value:       []float64{},
		Year:        []int{},
		UnifiedDate: []string{},
	}
	dc := ""
	thisYear := time.Now().Year()
	for _, v := range rows {
		if dc == v.DateCol {
			result.Value[len(result.Value)-1] += v.getValueFromDryDockFullRow(metric, absolute)
		} else {
			t, _ := time.Parse("2006-01-02", v.DateCol)
			result.Date = append(result.Date, v.DateCol)
			result.Value = append(result.Value, v.getValueFromDryDockFullRow(metric, absolute))
			result.Year = append(result.Year, t.Year())
			result.UnifiedDate = append(result.UnifiedDate, t.AddDate(thisYear-t.Year(), 0, 0).Format("2006-01-02"))
			dc = v.DateCol
		}
	}
	return result
}

// ConvertDryDockSummaryStats converts drydock summary stats data
func ConvertDryDockSummaryStats(rows []DryDockSummaryStatsFullRow, seg string) *DryDockSummaryStatsObject {
	result := &DryDockSummaryStatsObject{
		Month:              []string{},
		Segment:            seg,
		AverageDaysInDock:  []int{},
		CompletedDockStays: []int{},
	}
	// startDate := "2017-01-01"
	for _, v := range rows {
		result.Month = append(result.Month, v.Month)
		result.AverageDaysInDock = append(result.AverageDaysInDock, v.AverageDaysInDock)
		result.CompletedDockStays = append(result.CompletedDockStays, v.CompletedDockStays)
	}
	return result
}

func (u DryDockFullRow) getValueFromDryDockFullRow(m string, ab bool) float64 {
	if m == MetricVesselCount.String() || ab {
		return float64(u.VesselCount)
	} else if m == MetricVesselCount.String() || !ab {
		return u.VesselCountFleet
	} else if m == MetricDwt.String() || ab {
		return float64(u.Dwt)
	} else if m == MetricDwt.String() || !ab {
		return u.DwtFleet
	}
	return 0
}
