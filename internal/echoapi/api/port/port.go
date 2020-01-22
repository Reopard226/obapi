// Package user contains user application services
package port

import (
	"github.com/labstack/echo"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
)

// List returns list of ports
func (u *Port) List(c echo.Context, rid string, seg string) ([]model.Port, error) {
	// au := u.rbac.User(c)
	// q, err := query.List(au)
	// if err != nil {
	// 	return nil, err
	// }
	return u.udb.List(u.db, rid, seg)
}

// View returns single port
func (u *Port) View(c echo.Context, rid string, seg string) (*model.Port, error) {
	// if err := u.rbac.EnforceUser(c, id); err != nil {
	// 	return nil, err
	// }
	return u.udb.View(u.db, rid, seg)
}
