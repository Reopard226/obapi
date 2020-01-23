package main

import (
	"flag"
	"oceanbolt.com/obapi/internal/iam/dao"
	"os"

	"oceanbolt.com/obapi/internal/echoapi/api"

	"oceanbolt.com/obapi/internal/echoapi/utl/config"
)

var configStore config.Config

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

	db, err := dao.NewMongoDatabase(configStore.MONGODB_CONNECTION_STRING, "apikeys")
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
