package tonnage

import (
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"oceanbolt.com/obapi/internal/echoapi/api/tonnage/platform/mgo"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
)

// Service represents tonnage application interface
type Service interface {
	BasinBalance(echo.Context, string, string) (*model.TimeCountObject, error)
	DryDockTimeseries(echo.Context, string, string, bool) (*model.TimeseriesObject, error)
	DryDockSummaryStats(echo.Context, string) (*model.DryDockSummaryStatsObject, error)
}

// New creates new tonnage application service
func New(db *mongo.Database, udb UDB, rbac RBAC, sec Securer) *Tonnage {
	return &Tonnage{db, udb, rbac, sec}
}

// Initialize initalizes tonnage application service with defaults
func Initialize(db *mongo.Database, rbac RBAC, sec Securer) *Tonnage {
	return New(db, mgo.NewTonnage(), rbac, sec)
}

// Tonnage represents tonnage application service
type Tonnage struct {
	db   *mongo.Database
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents tonnage repository interface
type UDB interface {
	BasinBalance(*mongo.Database, string, string) (*model.TimeCountObject, error)
	DryDockTimeseries(*mongo.Database, string, string, bool) (*model.TimeseriesObject, error)
	DryDockSummaryStats(*mongo.Database, string) (*model.DryDockSummaryStatsObject, error)
}

type RBAC interface {
	User(echo.Context) *model.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, model.AccessRole, int, int) error
	IsLowerRole(echo.Context, model.AccessRole) error
}
