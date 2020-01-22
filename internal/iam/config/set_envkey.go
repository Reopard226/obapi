package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/envkey/envkeygo"
	"gopkg.in/yaml.v2"
)

// Load returns Configuration struct
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

// Configuration holds data necessery for configuring application
type Configuration struct {
	Server *Server      `yaml:"server,omitempty"`
	JWT    *JWT         `yaml:"jwt,omitempty"`
	App    *Application `yaml:"application,omitempty"`
}

// Server holds data necessery for server configuration
type Server struct {
	Port         string `yaml:"port,omitempty"`
	Debug        bool   `yaml:"debug,omitempty"`
	ReadTimeout  int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeout int    `yaml:"write_timeout_seconds,omitempty"`
}

// JWT holds data necessery for JWT configuration
type JWT struct {
	Secret           string `yaml:"secret,omitempty"`
	Duration         int    `yaml:"duration_minutes,omitempty"`
	RefreshDuration  int    `yaml:"refresh_duration_minutes,omitempty"`
	MaxRefresh       int    `yaml:"max_refresh_minutes,omitempty"`
	SigningAlgorithm string `yaml:"signing_algorithm,omitempty"`
}

// Application holds application configuration details
type Application struct {
	MinPasswordStr int    `yaml:"min_password_strength,omitempty"`
	SwaggerUIPath  string `yaml:"swagger_ui_path,omitempty"`
}

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
}

func (c *Config) ParseEnv() error {
	mustMapEnv(&c.MONGODB_CONNECTION_STRING, "MONGODB_CONNECTION_STRING")
	mustMapEnv(&c.MONGODB_DATABASE_NAME, "MONGODB_DATABASE_NAME")
	mustMapEnv(&c.MONGODB_COLLECTION_NAME, "MONGODB_COLLECTION_NAME")
	mustMapEnv(&c.JWKS_RS256_PRIVATE_KEY, "JWKS_RS256_PRIVATE_KEY")
	mustMapEnv(&c.AUTH0_MGMT_CLIENT_ID, "AUTH0_MGMT_CLIENT_ID")
	mustMapEnv(&c.AUTH0_MGMT_CLIENT_SECRET, "AUTH0_MGMT_CLIENT_SECRET")
	mustMapEnv(&c.AUTH0_DOMAIN, "AUTH0_DOMAIN")

	return nil
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}
