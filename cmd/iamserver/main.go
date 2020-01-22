package main

import (
	"context"
	"gopkg.in/auth0.v3/management"
	"log"
	"net/http"
	"oceanbolt.com/iamservice/internal/iam/config"
	"oceanbolt.com/iamservice/internal/iam/dao"
	"oceanbolt.com/iamservice/internal/iam/hooks"
	"oceanbolt.com/iamservice/internal/iam/iamserver"
	"oceanbolt.com/iamservice/rpc/iam"
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

	db, err := dao.NewMongoDatabase(cfg.MONGODB_CONNECTION_STRING, cfg.MONGODB_DATABASE_NAME)
	if err != nil {
		panic(err)
	}

	fs, err := dao.NewFireStoreDatabase(context.Background(), cfg.GCP_PROJECT)
	if err != nil {
		panic(err)
	}
	defer fs.Close()

	server := &iamserver.Server{
		Db:     db,
		Config: &cfg,
		Auth0:  auth0,
		Fs:     fs,
	} // implements Haberdasher interface

	twirpHandler := iam.NewApikeyServer(server, hooks.NewLoggerHooks())
	log.Println("Starting server")
	http.ListenAndServe(":"+port, twirpHandler)

}
