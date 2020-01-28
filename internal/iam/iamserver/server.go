package iamserver

import (
	"cloud.google.com/go/datastore"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/auth0.v3/management"
	"oceanbolt.com/obapi/internal/iam/config"
)

type Server struct {
	Db              *mongo.Database
	Config          *config.Config
	Auth0           *management.Management
	Ds              *datastore.Client
	PermissionCache *map[string]PermissionCache
}
