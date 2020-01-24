package tonnage

import (
	"time"

	"github.com/labstack/echo"
	"oceanbolt.com/obapi/internal/echoapi/api/tonnage"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
)

// New creates new tonnage logging service
func New(svc tonnage.Service, logger model.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents tonnage logging service
type LogService struct {
	tonnage.Service
	logger model.Logger
}

const name = "tonnage"

// ReqBasinBalance is BasinBalance request params type
type ReqBasinBalance struct {
	Basin   string
	Segment string
}

// ReqDryDockTimeseries is DryDockTimeseries request params type
type ReqDryDockTimeseries struct {
	Segment  string
	Metric   string
	Absolute bool
}

// ReqDryDockSummaryStats is DryDockSummaryStats request params type
type ReqDryDockSummaryStats struct {
	Segment string
}

// BasinBalance logging
func (ls *LogService) BasinBalance(c echo.Context, basin, seg string) (resp *model.TimeCountObject, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "BasinBalance request", err,
			map[string]interface{}{
				"req":  ReqBasinBalance{basin, seg},
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.BasinBalance(c, basin, seg)
}

// DryDockTimeseries logging
func (ls *LogService) DryDockTimeseries(c echo.Context, seg, metric string, absolute bool) (resp *model.TimeseriesObject, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "DryDockTimeseries request", err,
			map[string]interface{}{
				"req":  ReqDryDockTimeseries{seg, metric, absolute},
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.DryDockTimeseries(c, seg, metric, absolute)
}

// DryDockSummaryStats logging
func (ls *LogService) DryDockSummaryStats(c echo.Context, seg string) (resp *model.DryDockSummaryStatsObject, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "DryDockSummaryStats request", err,
			map[string]interface{}{
				"req":  ReqDryDockSummaryStats{seg},
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.DryDockSummaryStats(c, seg)
}
