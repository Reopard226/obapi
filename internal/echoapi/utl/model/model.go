package model

const (
	ColPort            = "ports"
	ColRegion          = "regions"
	ColAnchoragePort   = "anchorage_port"
	ColAnchorageRegion = "anchorage_region"
	ColBasinBalance    = "basin_balance"
	ColDryDock         = "drydocks"
	ColDryDockStays    = "drydock_stays"
)

type StringEnum string

func (s StringEnum) String() string {
	return string(s)
}

const (
	SegmentCapesize  StringEnum = "capesize"
	SegmentPanamax   StringEnum = "panamax"
	SegmentSupramax  StringEnum = "supramax"
	SegmentHandysize StringEnum = "handysize"
	SegmentTotal     StringEnum = "total"

	BasinAtlantic    StringEnum = "atlantic"
	BasinIndianOcean StringEnum = "indian_ocean"
	BasinPacificEast StringEnum = "pacific_east"
	BasinPacificWest StringEnum = "pacific_west"

	MetricVesselCount StringEnum = "vessel_count"
	MetricDwt         StringEnum = "dwt"
)

// CheckValidSegment checks the segment validity
func CheckValidSegment(str string) bool {
	if str == SegmentCapesize.String() ||
		str == SegmentHandysize.String() ||
		str == SegmentPanamax.String() ||
		str == SegmentTotal.String() ||
		str == SegmentSupramax.String() {
		return true
	}
	return false
}

// CheckValidBasin checks the basin validity
func CheckValidBasin(str string) bool {
	if str == BasinAtlantic.String() ||
		str == BasinIndianOcean.String() ||
		str == BasinPacificEast.String() ||
		str == BasinPacificWest.String() {
		return true
	}
	return false
}

// CheckValidMetric checks the metric validity
func CheckValidMetric(str string) bool {
	if str == MetricVesselCount.String() ||
		str == MetricDwt.String() {
		return true
	}
	return false
}
