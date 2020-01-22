package model

// VoyageData represents VoyageData model
type VoyageData struct {
	Base
	Imo          int    `json:"imo" bson:"imo"`
	VesselName   string `json:"vessel_name" bson:"vessel_name"`
	Segment      string `json:"segment" bson:"segment"`
	Commodity    string `json:"commodity" bson:"commodity"`
	Volume       int64  `json:"volume" bson:"volume"`
	FromPortID   int64  `json:"from_port_id" bson:"from_port_id"`
	FromPortName string `json:"from_port_name" bson:"from_port_name"`
	//From_port_coords_wkt string   `json:"from_port_coords_wkt" bson:"from_port_coords_wkt"`
	FromBerthName      string `json:"from_berth_name" bson:"from_berth_name"`
	FromCountryCode    string `json:"from_country_code" bson:"from_country_code"`
	FromCountry        string `json:"from_country" bson:"from_country"`
	FromBerthArrival   string `json:"from_berth_arrival" bson:"from_berth_arrival"`
	FromBerthDeparture string `json:"from_berth_departure" bson:"from_berth_departure"`
	ToPortID           int64  `json:"to_port_id" bson:"to_port_id"`
	ToPortName         string `json:"to_port_name" bson:"to_port_name"`
	//To_port_coords_wkt   string   `json:"to_port_coords_wkt" bson:"to_port_coords_wkt"`
	ToBerthName      string `json:"to_berth_name" bson:"to_berth_name"`
	ToCountryCode    string `json:"to_country_code" bson:"to_country_code"`
	ToCountry        string `json:"to_country" bson:"to_country"`
	ToBerthArrival   string `json:"to_berth_arrival" bson:"to_berth_arrival"`
	ToBerthDeparture string `json:"to_berth_departure" bson:"to_berth_departure"`
}

// VoyageDataObject represents VoyageDataObject model
type VoyageDataObject struct {
	CurrentPage  int          `json:"current_page" bson:"current_page"`
	PreviousPage *int         `json:"previous_page" bson:"previous_page"`
	NextPage     *int         `json:"next_page" bson:"next_page"`
	LastPage     int          `json:"max_page" bson:"max_page"`
	Records      int          `json:"records" bson:"records"`
	TotalRecords int          `json:"total_records" bson:"total_records"`
	Data         []VoyageData `json:"data" bson:"data"`
}

//SegmentVolumeData is segment volume structs
type SegmentVolumeData struct {
	Segment string  `json:"segment"`
	Volume  float64 `json:"volume"`
}

// SegmentVolumeObject represents SegmentVolumeObject model
type SegmentVolumeObject struct {
	Segment []string  `json:"segment"`
	Volume  []float64 `json:"volume"`
}

// CountryVolumeObject represents CountryVolumeObject model
type CountryVolumeObject struct {
	Country     []string  `json:"country" bson:"country"`
	CountryCode []string  `json:"country_code" bson:"country_code"`
	Volume      []float64 `json:"volume" bson:"volume"`
}

// CountryVolumeData is Country volume structs
type CountryVolumeData struct {
	Country     string  `json:"country" bson:"country"`
	CountryCode string  `json:"country_code" bson:"country_code"`
	Volume      float64 `json:"volume" bson:"volume"`
}

//CountryData is Country list struct
type CountryData struct {
	Country     string `json:"country" bson:"country"`
	CountryCode string `json:"country_code" bson:"country_code"`
}

//PortVolumeData is Port Volumes structs
type PortVolumeData struct {
	PortName    string  `json:"port_name" bson:"port_name"`
	CountryCode string  `json:"country_code" bson:"country_code"`
	Coords      string  `json:"coords" bson:"coords"`
	Volume      float64 `json:"volume" bson:"volume"`
}

// PortVolumeObject represents PortVolumeObject model
type PortVolumeObject struct {
	PortName    []string     `json:"port_name" bson:"port_name"`
	CountryCode []string     `json:"country_code" bson:"country_code"`
	Coords      []*[]float64 `json:"coords" bson:"coords"`
	Volume      []float64    `json:"volume" bson:"volume"`
}

// YearMonthVolume represents YearMonthVolume model
type YearMonthVolume struct {
	YearMonth string  `json:"year_month" bson:"year_month"`
	Volume    float64 `json:"volume" bson:"volume"`
}

// YearWeekVolume represents YearWeekVolume model
type YearWeekVolume struct {
	YearWeek string  `json:"year_week" bson:"year_week"`
	Volume   float64 `json:"volume" bson:"volume"`
}

// TimeVolume represents TimeVolume model
type TimeVolume struct {
	Date   string  `json:"date" bson:"date"`
	Volume float64 `json:"volume" bson:"volume"`
}

// TimeVolumeObject represents TimeVolumeObject model
type TimeVolumeObject struct {
	Date   []string  `json:"date" bson:"date"`
	Volume []float64 `json:"volume" bson:"volume"`
}

// CommodityData represents CommodityData model
type CommodityData struct {
	CommodityName  string `json:"commodity" bson:"commodity"`
	CommodityValue string `json:"commodity_value" bson:"commodity_value"`
}

// StatsDataObject represents StatsDataObject model
type StatsDataObject struct {
	Voyages                string  `json:"voyages"`
	Volume                 float64 `json:"volume"`
	UniqueOriginPorts      int     `json:"unique_origin_ports"`
	UniqueDestinationPorts int     `json:"unique_destination_ports"`
	UniqueVessels          int     `json:"unique_vessels"`
}
