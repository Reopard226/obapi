package congestion

import (
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"oceanbolt.com/obapi/internal/echoapi/api/congestion/platform/mgo"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
)

// Service represents congestion application interface
type Service interface {
	CongestionPort(echo.Context, string, string) (*model.AnchoragePortArray, error)
	CongestionRegion(echo.Context, string, string) (*model.AnchorageRegionArray, error)
	ListPort(echo.Context) ([]model.Port, error)
	ListRegion(echo.Context) ([]model.Region, error)
}

// New creates new congestion application service
func New(db *mongo.Database, udb UDB, rbac RBAC, sec Securer) *Congestion {
	return &Congestion{db, udb, rbac, sec}
}

// Initialize initalizes congestion application service with defaults
func Initialize(db *mongo.Database, rbac RBAC, sec Securer) *Congestion {
	return New(db, mgo.NewCongestion(), rbac, sec)
}

// Congestion represents congestion application service
type Congestion struct {
	db   *mongo.Database
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents congestion repository interface
type UDB interface {
	CongestionPort(*mongo.Database, string, string) (*model.AnchoragePortArray, error)
	CongestionRegion(*mongo.Database, string, string) (*model.AnchorageRegionArray, error)
	ListPort(*mongo.Database) ([]model.Port, error)
	ListRegion(*mongo.Database) ([]model.Region, error)
}

type RBAC interface {
	User(echo.Context) *model.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, model.AccessRole, int, int) error
	IsLowerRole(echo.Context, model.AccessRole) error
}
