package model

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AnchorageDataRegion represents AnchorageDataRegion model
type AnchorageDataRegion struct {
	ID                       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	RegionID                 string             `json:"region_id" bson:"region_id"`
	DateActual               string             `json:"date_actual" bson:"date_actual"`
	Segment                  string             `json:"segment" bson:"segment"`
	NCongestedVessels        int                `json:"n_congested_vessels" bson:"n_congested_vessels"`
	AvgWaitingTime           int                `json:"avg_waiting_time" bson:"avg_waiting_time"`
	FirstQuartileWaitingTime int                `json:"first_quartile_waiting_time" bson:"first_quartile_waiting_time"`
	MedianWaitingTime        int                `json:"median_waiting_time" bson:"median_waiting_time"`
	ThirdQuartileWaitingTime int                `json:"third_quartile_waiting_time" bson:"third_quartile_waiting_time"`
	MaxWaitingTime           int                `json:"max_waiting_time" bson:"max_waiting_time"`
	DayOfYear                int                `json:"day_of_year" bson:"day_of_year"`
	Year                     int                `json:"year" bson:"year"`
	UnifiedDate              string             `json:"unified_date" bson:"unified_date"`
}

// AnchorageDataRegionArray represents AnchorageDataRegionArray model
type AnchorageDataRegionArray struct {
	RegionID                 string   `json:"region_id" bson:"region_id"`
	DateActual               []string `json:"date_actual" bson:"date_actual"`
	Segment                  string   `json:"segment" bson:"segment"`
	NCongestedVessels        []int    `json:"n_congested_vessels" bson:"n_congested_vessels"`
	AvgWaitingTime           []int    `json:"avg_waiting_time" bson:"avg_waiting_time"`
	FirstQuartileWaitingTime []int    `json:"first_quartile_waiting_time" bson:"first_quartile_waiting_time"`
	MedianWaitingTime        []int    `json:"median_waiting_time" bson:"median_waiting_time"`
	ThirdQuartileWaitingTime []int    `json:"third_quartile_waiting_time" bson:"third_quartile_waiting_time"`
	MaxWaitingTime           []int    `json:"max_waiting_time" bson:"max_waiting_time"`
	DayOfYear                []int    `json:"day_of_year" bson:"day_of_year"`
	Year                     []int    `json:"year" bson:"year"`
	UnifiedDate              []string `json:"unified_date" bson:"unified_date"`
}

func (a AnchorageDataRegion) String() string {
	return fmt.Sprintf("AnchorageDataRegion<%s \t| %s \t| %d\t>\n", a.DateActual, a.Segment, a.NCongestedVessels)
}

// ConvertAnchorageRegionOutput converts AnchorageRegion output
func ConvertAnchorageRegionOutput(a []AnchorageDataRegion) AnchorageDataRegionArray {
	var regionID = a[0].RegionID
	var dateActual = make([]string, len(a))
	var segment = a[0].Segment
	var nCongestedVessels = make([]int, len(a))
	var avgWaitingTime = make([]int, len(a))
	var firstQuartileWaitingTime = make([]int, len(a))
	var medianWaitingTime = make([]int, len(a))
	var thirdQuartileWaitingTime = make([]int, len(a))
	var maxWaitingTime = make([]int, len(a))
	var dayOfYear = make([]int, len(a))
	var year = make([]int, len(a))
	var unifiedDate = make([]string, len(a))

	for k, v := range a {
		dateActual[k] = v.DateActual
		nCongestedVessels[k] = v.NCongestedVessels
		avgWaitingTime[k] = v.AvgWaitingTime
		firstQuartileWaitingTime[k] = v.FirstQuartileWaitingTime
		medianWaitingTime[k] = v.MedianWaitingTime
		thirdQuartileWaitingTime[k] = v.ThirdQuartileWaitingTime
		maxWaitingTime[k] = v.MaxWaitingTime
		dayOfYear[k] = v.DayOfYear
		year[k] = v.Year
		unifiedDate[k] = v.UnifiedDate
	}

	return AnchorageDataRegionArray{
		RegionID:                 regionID,
		DateActual:               dateActual,
		Segment:                  segment,
		NCongestedVessels:        nCongestedVessels,
		AvgWaitingTime:           avgWaitingTime,
		FirstQuartileWaitingTime: firstQuartileWaitingTime,
		MedianWaitingTime:        medianWaitingTime,
		ThirdQuartileWaitingTime: thirdQuartileWaitingTime,
		MaxWaitingTime:           maxWaitingTime,
		DayOfYear:                dayOfYear,
		Year:                     year,
		UnifiedDate:              unifiedDate,
	}
}
