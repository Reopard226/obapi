package congestion

import (
	"github.com/labstack/echo"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
)

// CongestionPort returns AnchoragePortArray
func (u *Congestion) CongestionPort(c echo.Context, pid, seg string) (*model.AnchoragePortArray, error) {
	return u.udb.CongestionPort(u.db, pid, seg)
}

// CongestionRegion returns AnchorageRegionArray
func (u *Congestion) CongestionRegion(c echo.Context, rid, seg string) (*model.AnchorageRegionArray, error) {
	return u.udb.CongestionRegion(u.db, rid, seg)
}

// ListPort returns list of ports
func (u *Congestion) ListPort(c echo.Context) ([]model.Port, error) {
	return u.udb.ListPort(u.db)
}

// ListRegion returns list regions
func (u *Congestion) ListRegion(c echo.Context) ([]model.Region, error) {
	return u.udb.ListRegion(u.db)
}
