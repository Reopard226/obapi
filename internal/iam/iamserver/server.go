package iamserver

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/auth0.v3/management"
	"oceanbolt.com/iamservice/internal/iam/config"
)

type Server struct {
	Db              *mongo.Database
	Config          *config.Config
	Auth0           *management.Management
	PermissionCache *map[string]PermissionCache
}
