package main

import (
	"log"
	"net/http"
	"oceanbolt.com/iamservice/internal/config"
	"oceanbolt.com/iamservice/internal/dao"
	"oceanbolt.com/iamservice/internal/hooks"
	"oceanbolt.com/iamservice/internal/iamserver"
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

	db, err := dao.NewMongoDatabase(cfg.MONGODB_CONNECTION_STRING,cfg.MONGODB_DATABASE_NAME)
	if err != nil {
		panic(err)
	}

	server := &iamserver.Server{
		Db:db,
		Config:&cfg,
	} // implements Haberdasher interface

	twirpHandler := iam.NewApikeyServer(server, hooks.NewLoggerHooks())
	log.Println("Starting server")
	http.ListenAndServe(":8080", twirpHandler)

}
