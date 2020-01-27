package apiaccess

import (
	"github.com/labstack/echo"
	"oceanbolt.com/obapi/internal/echoapi/api/apiaccess/platform/mgo"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
	"oceanbolt.com/obapi/internal/iam/iamclient"
	"oceanbolt.com/obapi/rpc/iam"
)

// Service represents apiaccess application interface
type Service interface {
	ListKey(echo.Context) (*iam.UserKeys, error)
	CreateKey(echo.Context, string, int64) (*iam.UserKeyWithSecret, error)
	DeleteKey(echo.Context, string) (*iam.KeyDeletedResponse, error)
}

// New creates new apiaccess application service
func New(client iamclient.OceanboltIAMClient, udb UDB, rbac RBAC, sec Securer) *Apiaccess {
	return &Apiaccess{client, udb, rbac, sec}
}

// Initialize initalizes apiaccess application service with defaults
func Initialize(client iamclient.OceanboltIAMClient, rbac RBAC, sec Securer) *Apiaccess {
	return New(client, mgo.NewApiaccess(), rbac, sec)
}

// Apiaccess represents apiaccess application service
type Apiaccess struct {
	client iamclient.OceanboltIAMClient
	udb    UDB
	rbac   RBAC
	sec    Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents apiaccess repository interface
type UDB interface {
	ListKey(iamclient.OceanboltIAMClient) (*iam.UserKeys, error)
	CreateKey(iamclient.OceanboltIAMClient, string, int64) (*iam.UserKeyWithSecret, error)
	DeleteKey(iamclient.OceanboltIAMClient, string) (*iam.KeyDeletedResponse, error)
}

type RBAC interface {
	User(echo.Context) *model.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, model.AccessRole, int, int) error
	IsLowerRole(echo.Context, model.AccessRole) error
}
