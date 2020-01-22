package port

import (
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"oceanbolt.com/iamservice/pkg/api/port/platform/mgo"
	"oceanbolt.com/iamservice/pkg/utl/model"
)

// Service represents port application interface
type Service interface {
	List(echo.Context, string, string) ([]model.Port, error)
	View(echo.Context, string, string) (*model.Port, error)
}

// New creates new port application service
func New(db *mongo.Database, udb UDB, rbac RBAC, sec Securer) *Port {
	return &Port{db, udb, rbac, sec}
}

// Initialize initalizes port application service with defaults
func Initialize(db *mongo.Database, rbac RBAC, sec Securer) *Port {
	return New(db, mgo.NewPort(), rbac, sec)
}

// Port represents port application service
type Port struct {
	db   *mongo.Database
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents port repository interface
type UDB interface {
	List(*mongo.Database, string, string) ([]model.Port, error)
	View(*mongo.Database, string, string) (*model.Port, error)
}

type RBAC interface {
	User(echo.Context) *model.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, model.AccessRole, int, int) error
	IsLowerRole(echo.Context, model.AccessRole) error
}
