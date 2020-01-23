// Package user contains user application services
package region

import (
	"github.com/labstack/echo"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
)

// List returns list of regions
func (u *Region) List(c echo.Context, rid string, seg string) ([]model.Region, error) {
	// au := u.rbac.User(c)
	// q, err := query.List(au)
	// if err != nil {
	// 	return nil, err
	// }
	return u.udb.List(u.db, rid, seg)
}

// View returns single region
func (u *Region) View(c echo.Context, rid string, seg string) (*model.Region, error) {
	// if err := u.rbac.EnforceUser(c, id); err != nil {
	// 	return nil, err
	// }
	return u.udb.View(u.db, rid, seg)
}
