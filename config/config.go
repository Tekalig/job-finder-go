// config/config.go
package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HasuraEndpoint    string
	HasuraAdminSecret string
	JWTSecret         string
	SMTPHost          string
	SMTPPort          string
	SMTPUser          string
	SMTPPassword      string
}

func Load() (*Config, error) {
	godotenv.Load()

	return &Config{
		HasuraEndpoint:    os.Getenv("HASURA_ENDPOINT"),
		HasuraAdminSecret: os.Getenv("HASURA_ADMIN_SECRET"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		SMTPHost:          os.Getenv("SMTP_HOST"),
		SMTPPort:          os.Getenv("SMTP_PORT"),
		SMTPUser:          os.Getenv("SMTP_USER"),
		SMTPPassword:      os.Getenv("SMTP_PASSWORD"),
	}, nil
}
