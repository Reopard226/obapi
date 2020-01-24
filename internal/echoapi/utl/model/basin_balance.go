package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BasinBalanceFullRow represents BasinBalanceFullRow model
type BasinBalanceFullRow struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Basin       string             `json:"basin" bson:"basin"`
	Segment     string             `json:"segment" bson:"segment"`
	DateCol     string             `json:"date_col" bson:"date_col"`
	VesselCount int                `json:"vessel_count" bson:"vessel_count"`
}

// BasinBalanceCount represents BasinBalanceCount model
type BasinBalanceCount struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	DateCol     string             `json:"date_col" bson:"date_col"`
	VesselCount int                `json:"vessel_count" bson:"vessel_count"`
}

// TimeCountObject represents TimeCountObject model
type TimeCountObject struct {
	Date        []string `json:"date"`
	VesselCount []int    `json:"vessel_count"`
	Year        []int    `json:"year"`
	UnifiedDate []string `json:"unified_date"`
}

// ConvertBasinTimeCount converts basinbalance array to timecount object
func ConvertBasinTimeCount(rows []BasinBalanceFullRow) *TimeCountObject {
	result := &TimeCountObject{
		Date:        []string{},
		VesselCount: []int{},
		Year:        []int{},
		UnifiedDate: []string{},
	}
	dc := ""
	thisYear := time.Now().Year()
	for _, v := range rows {
		if dc == v.DateCol {
			result.VesselCount[len(result.VesselCount)-1] += v.VesselCount
		} else {
			t, _ := time.Parse("2006-01-02", v.DateCol)
			result.Date = append(result.Date, v.DateCol)
			result.VesselCount = append(result.VesselCount, v.VesselCount)
			result.Year = append(result.Year, t.Year())
			result.UnifiedDate = append(result.UnifiedDate, t.AddDate(thisYear-t.Year(), 0, 0).Format("2006-01-02"))
			dc = v.DateCol
		}
	}
	return result
}
