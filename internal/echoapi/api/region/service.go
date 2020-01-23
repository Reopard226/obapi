package region

import (
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"oceanbolt.com/obapi/internal/echoapi/api/region/platform/mgo"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
)

// Service represents region application interface
type Service interface {
	List(echo.Context, string, string) ([]model.Region, error)
	View(echo.Context, string, string) (*model.Region, error)
}

// New creates new region application service
func New(db *mongo.Database, udb UDB, rbac RBAC, sec Securer) *Region {
	return &Region{db, udb, rbac, sec}
}

// Initialize initalizes region application service with defaults
func Initialize(db *mongo.Database, rbac RBAC, sec Securer) *Region {
	return New(db, mgo.NewRegion(), rbac, sec)
}

// Region represents region application service
type Region struct {
	db   *mongo.Database
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents region repository interface
type UDB interface {
	List(*mongo.Database, string, string) ([]model.Region, error)
	View(*mongo.Database, string, string) (*model.Region, error)
}

type RBAC interface {
	User(echo.Context) *model.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, model.AccessRole, int, int) error
	IsLowerRole(echo.Context, model.AccessRole) error
}
