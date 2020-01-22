package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/envkey/envkeygo"
)

func SetEnvKey() {
	log.Println("Envkey is set")
}

type Config struct {
	MONGODB_CONNECTION_STRING string
	MONGODB_DATABASE_NAME     string
	MONGODB_COLLECTION_NAME   string
	JWKS_RS256_PRIVATE_KEY    string
	AUTH0_MGMT_CLIENT_ID      string
	AUTH0_MGMT_CLIENT_SECRET  string
	AUTH0_DOMAIN              string
	GCP_PROJECT               string
}

func (c *Config) ParseEnv() error {
	mustMapEnv(&c.MONGODB_CONNECTION_STRING, "MONGODB_CONNECTION_STRING")
	mustMapEnv(&c.MONGODB_DATABASE_NAME, "MONGODB_DATABASE_NAME")
	mustMapEnv(&c.MONGODB_COLLECTION_NAME, "MONGODB_COLLECTION_NAME")
	mustMapEnv(&c.JWKS_RS256_PRIVATE_KEY, "JWKS_RS256_PRIVATE_KEY")
	mustMapEnv(&c.AUTH0_MGMT_CLIENT_ID, "AUTH0_MGMT_CLIENT_ID")
	mustMapEnv(&c.AUTH0_MGMT_CLIENT_SECRET, "AUTH0_MGMT_CLIENT_SECRET")
	mustMapEnv(&c.AUTH0_DOMAIN, "AUTH0_DOMAIN")
	mustMapEnv(&c.GCP_PROJECT, "GCP_PROJECT")

	return nil
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}
