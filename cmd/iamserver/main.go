package main

import (
	"context"
	"gopkg.in/auth0.v3/management"
	"log"
	"net/http"
	"oceanbolt.com/obapi/internal/iam/config"
	"oceanbolt.com/obapi/internal/iam/dao"
	"oceanbolt.com/obapi/internal/iam/hooks"
	"oceanbolt.com/obapi/internal/iam/iamserver"
	"oceanbolt.com/obapi/rpc/iam"
	"os"
)

var cfg config.Config

func init() {
	if os.Getenv("ENVKEY_IS_SET") != "TRUE" {
		config.SetEnvKey()
		os.Setenv("ENVKEY_IS_SET", "TRUE")
	}

	err := cfg.ParseEnv()
	if err != nil {
		panic(err)
	}

}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	auth0, err := management.New(cfg.AUTH0_DOMAIN, cfg.AUTH0_MGMT_CLIENT_ID, cfg.AUTH0_MGMT_CLIENT_SECRET)
	if err != nil {
		panic(err)
	}

	/*fs, err := dao.NewFireStoreDatabase(context.Background(), cfg.GCP_PROJECT)
	if err != nil {
		panic(err)
	}
	defer fs.Close()
	*/
	ds, err := dao.NewDataStoreDatabase(context.Background(), cfg.GCP_PROJECT)
	if err != nil {
		panic(err)
	}
	defer ds.Close()

	if ds == nil {
		panic("ERROR NO CLIENT!!!!!!")
	}

	server := &iamserver.Server{
		Config: &cfg,
		Auth0:  auth0,
		Ds:     ds,
	} // implements Haberdasher interface

	twirpHandler := iam.NewApikeyServer(server, hooks.NewLoggerHooks())
	log.Println("Starting server")
	http.ListenAndServe(":"+port, twirpHandler)

}
