package api

import (
	"crypto/sha1"
	"fmt"
	"oceanbolt.com/obapi/internal/iam/iamclient"

	"go.mongodb.org/mongo-driver/mongo"
	"oceanbolt.com/obapi/internal/echoapi/utl/zlog"

	"oceanbolt.com/obapi/internal/echoapi/api/apiaccess"
	al "oceanbolt.com/obapi/internal/echoapi/api/apiaccess/logging"
	at "oceanbolt.com/obapi/internal/echoapi/api/apiaccess/transport"

	"oceanbolt.com/obapi/internal/echoapi/api/tonnage"
	tl "oceanbolt.com/obapi/internal/echoapi/api/tonnage/logging"
	tt "oceanbolt.com/obapi/internal/echoapi/api/tonnage/transport"

	"oceanbolt.com/obapi/internal/echoapi/api/congestion"
	cl "oceanbolt.com/obapi/internal/echoapi/api/congestion/logging"
	ct "oceanbolt.com/obapi/internal/echoapi/api/congestion/transport"

	"oceanbolt.com/obapi/internal/echoapi/utl/config"
	"oceanbolt.com/obapi/internal/echoapi/utl/middleware/jwt"
	"oceanbolt.com/obapi/internal/echoapi/utl/rbac"
	"oceanbolt.com/obapi/internal/echoapi/utl/secure"
	"oceanbolt.com/obapi/internal/echoapi/utl/server"
)

// Start starts the API service
func Start(db *mongo.Database, cfg *config.Configuration, envkeyCfg *config.Config) error {
	sec := secure.New(cfg.App.MinPasswordStr, sha1.New())
	rbac := rbac.New()
	iam := iamclient.GetDefaultIamClient(envkeyCfg.IAM_SERVICE_URL)

	jwtService := jwt.New(cfg.JWT.Secret, cfg.JWT.SigningAlgorithm, cfg.JWT.Duration, iam)
	log := zlog.New()

	e := server.New()

	v1 := e.Group("/v1")

	useAUTH := false
	if useAUTH {
		v1.Use(jwtService.MWFunc())
	}

	at.NewHTTP(al.New(apiaccess.Initialize(iam, rbac, sec), log), v1)
	tt.NewHTTP(tl.New(tonnage.Initialize(db, rbac, sec), log), v1)
	ct.NewHTTP(cl.New(congestion.Initialize(db, rbac, sec), log), v1)

	fmt.Printf("Starting server")
	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
