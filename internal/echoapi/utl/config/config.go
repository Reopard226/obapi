package config

import (
	"fmt"
	_ "github.com/envkey/envkeygo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
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
	JWKS_RS256_PRIVATE_KEY    string
	AUTH0_MGMT_APIKEY_ID      string
	AUTH0_MGMT_APIKEY_SECRET  string
	IAM_SERVICE_URL 		  string
}

func (c *Config) ParseEnv() error {
	mustMapEnv(&c.MONGODB_CONNECTION_STRING, "MONGODB_CONNECTION_STRING")
	mustMapEnv(&c.JWKS_RS256_PRIVATE_KEY, "JWKS_RS256_PRIVATE_KEY")
	mustMapEnv(&c.AUTH0_MGMT_APIKEY_ID, "AUTH0_MGMT_APIKEY_ID")
	mustMapEnv(&c.AUTH0_MGMT_APIKEY_SECRET, "AUTH0_MGMT_APIKEY_SECRET")
	mustMapEnv(&c.IAM_SERVICE_URL, "IAM_SERVICE_URL")

	return nil
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}
