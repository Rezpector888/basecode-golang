package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfigType struct {
	Port    string `envconfig:"APP_PORT" default:"8080"`
	AppName string `envconfig:"APP_NAME" default:"basecode"`
	AppMode string `envconfig:"APP_MODE" default:"debug" `
}

type DatabaseConfigType struct {
	Host     string `envconfig:"DATABASE_HOST" required:"true"`
	Port     string `envconfig:"DATABASE_PORT" required:"true" default:"5432"`
	User     string `envconfig:"DATABASE_USER" required:"true"`
	Password string `envconfig:"DATABASE_PASSWORD" required:"true"`
	Name     string `envconfig:"DATABASE_NAME" required:"true"`
}

type JwtConfigType struct {
	Secret string `envconfig:"JWT_SECRET" required:"true"`
}

var AppConfig AppConfigType
var DatabaseConfig DatabaseConfigType
var JwtConfig JwtConfigType

func LoadConfig() {
	_ = godotenv.Load(".env")

	if err := envconfig.Process("", &AppConfig); err != nil {
		log.Fatalf("Failed to load AppConfig: %v", err)
	}

	if err := envconfig.Process("", &DatabaseConfig); err != nil {
		log.Fatalf("Failed to load DatabaseConfig: %v", err)
	}

	if err := envconfig.Process("", &JwtConfig); err != nil {
		log.Fatalf("Failed to load JwtConfig: %v", err)
	}

	switch AppConfig.AppMode {
	case "debug", "production", "testing":
		// valid
	default:
		log.Fatalf("Invalid APP_MODE: %s. Must be one of: debug, production, testing", AppConfig.AppMode)
	}
}
