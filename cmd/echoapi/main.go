package main

import (
	"flag"
	"os"

	"oceanbolt.com/obapi/internal/iam/dao"

	"oceanbolt.com/obapi/internal/echoapi/api"

	"oceanbolt.com/obapi/internal/echoapi/utl/config"
)

var configStore config.Config

const localMongoString = "mongodb://localhost:27017"

func init() {
	if os.Getenv("ENVKEY_IS_SET") != "TRUE" {
		config.SetEnvKey()
		os.Setenv("ENVKEY_IS_SET", "TRUE")
	}
	err := configStore.ParseEnv()
	if err != nil {
		panic(err)
	}
}

func main() {

	cfgPath := flag.String("p", "./cmd/echoapi/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfgs, err := config.Load(*cfgPath)
	checkErr(err)

	// db, err := dao.NewMongoDatabase(configStore.MONGODB_CONNECTION_STRING, "apikeys")
	db, err := dao.NewMongoDatabase(localMongoString, "test")

	if err != nil {
		panic(err.Error())
	}

	checkErr(api.Start(db, cfgs,&configStore))

}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
