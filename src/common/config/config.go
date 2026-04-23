package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfigType struct {
	Port     string `envconfig:"APP_PORT" default:"8080"`
	AppName  string `envconfig:"APP_NAME" default:"basecode"`
	AppMode  string `envconfig:"APP_MODE" default:"debug" `
	GrpcPort string `envconfig:"GRPC_PORT" default:"50051"`
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

type MinioConfigType struct {
	AccessKey  string `envconfig:"MINIO_ACCESS_KEY" required:"true"`
	SecretKey  string `envconfig:"MINIO_SECRET_KEY" required:"true"`
	BucketName string `envconfig:"MINIO_BUCKET_NAME" required:"true"`
	EndPoint   string `envconfig:"MINIO_ENDPOINT" required:"true"`
	UseSSL     bool   `envconfig:"MINIO_USE_SSL" required:"true"`
}

var AppConfig AppConfigType
var DatabaseConfig DatabaseConfigType
var JwtConfig JwtConfigType
var MinioConfig MinioConfigType

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

	if err := envconfig.Process("", &MinioConfig); err != nil {
		log.Fatalf("Failed to load MinioConfig: %v", err)

	}

	switch AppConfig.AppMode {
	case "debug", "production", "testing":
		// valid
	default:
		log.Fatalf("Invalid APP_MODE: %s. Must be one of: debug, production, testing", AppConfig.AppMode)
	}
}
