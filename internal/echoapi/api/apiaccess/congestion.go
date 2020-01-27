package apiaccess

import (
	"github.com/labstack/echo"
	"oceanbolt.com/obapi/rpc/iam"
)

// ListKey returns all api access keys
func (u *Apiaccess) ListKey(c echo.Context) (*iam.UserKeys, error) {
	return u.udb.ListKey(u.client)
}

// CreateKey creates new api access key
func (u *Apiaccess) CreateKey(c echo.Context, tag string, exp int64) (*iam.UserKeyWithSecret, error) {
	return u.udb.CreateKey(u.client, tag, exp)
}

// DeleteKey deletes api access key
func (u *Apiaccess) DeleteKey(c echo.Context, apikeyID string) (*iam.KeyDeletedResponse, error) {
	return u.udb.DeleteKey(u.client, apikeyID)
}
