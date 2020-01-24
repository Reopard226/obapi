package tonnage

import (
	"github.com/labstack/echo"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
)

// BasinBalance returns TimeCountObject
func (u *Tonnage) BasinBalance(c echo.Context, basin, seg string) (*model.TimeCountObject, error) {
	return u.udb.BasinBalance(u.db, basin, seg)
}

// DryDockTimeseries returns TimeseriesObject
func (u *Tonnage) DryDockTimeseries(c echo.Context, seg, metric string, absolute bool) (*model.TimeseriesObject, error) {
	return u.udb.DryDockTimeseries(u.db, seg, metric, absolute)
}

// DryDockSummaryStats returns DryDockSummaryStatsObject
func (u *Tonnage) DryDockSummaryStats(c echo.Context, seg string) (*model.DryDockSummaryStatsObject, error) {
	return u.udb.DryDockSummaryStats(u.db, seg)
}
