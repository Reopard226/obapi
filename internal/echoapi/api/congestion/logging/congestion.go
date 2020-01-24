package congestion

import (
	"time"

	"github.com/labstack/echo"
	"oceanbolt.com/obapi/internal/echoapi/api/congestion"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
)

// New creates new congestion logging service
func New(svc congestion.Service, logger model.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents congestion logging service
type LogService struct {
	congestion.Service
	logger model.Logger
}

const name = "congestion"

// ReqPort is port request params type
type ReqPort struct {
	PortID  string
	Segment string
}

// ReqRegion is region request params type
type ReqRegion struct {
	RegionID string
	Segment  string
}

// CongestionPort logging
func (ls *LogService) CongestionPort(c echo.Context, portID, segment string) (resp *model.AnchoragePortArray, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "AnchoragePortArray request", err,
			map[string]interface{}{
				"req":  ReqPort{portID, segment},
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.CongestionPort(c, portID, segment)
}

// CongestionRegion logging
func (ls *LogService) CongestionRegion(c echo.Context, regionID, segment string) (resp *model.AnchorageRegionArray, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "CongestionRegion request", err,
			map[string]interface{}{
				"req":  ReqRegion{regionID, segment},
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.CongestionRegion(c, regionID, segment)
}

// ListPort logging
func (ls *LogService) ListPort(c echo.Context) (resp []model.Port, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "List port request", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ListPort(c)
}

// ListRegion logging
func (ls *LogService) ListRegion(c echo.Context) (resp []model.Region, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "View region request", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ListRegion(c)
}
