package region

import (
	"time"

	"github.com/labstack/echo"
	"oceanbolt.com/obapi/internal/echoapi/api/region"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
)

// New creates new region logging service
func New(svc region.Service, logger model.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents region logging service
type LogService struct {
	region.Service
	logger model.Logger
}

const name = "region"

// Req is request params type
type Req struct {
	RegionID string
	Segment  string
}

// List logging
func (ls *LogService) List(c echo.Context, regionID, segment string) (resp []model.Region, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "List region request", err,
			map[string]interface{}{
				"req":  Req{regionID, segment},
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List(c, regionID, segment)
}

// View logging
func (ls *LogService) View(c echo.Context, regionID, segment string) (resp *model.Region, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "View region request", err,
			map[string]interface{}{
				"req":  Req{regionID, segment},
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.View(c, regionID, segment)
}
