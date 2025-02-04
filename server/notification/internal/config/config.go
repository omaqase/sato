package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	App         AppConfig
	Resend      ResendConfig
	UserService UserServiceConfig
}

type AppConfig struct {
	Environment string `envconfig:"APP_ENV" required:"true"`
	Port        string `envconfig:"APP_PORT" default:"8080" required:"true"`
}

type ResendConfig struct {
	APIKey string `envconfig:"RESEND_API_KEY" required:"true"`
}

type UserServiceConfig struct {
	Host string `envconfig:"USER_SERVICE_HOST" required:"true"`
	Port string `envconfig:"USER_SERVICE_PORT" required:"true"`
}

func Load() (Config, error) {
	root, err := os.Getwd()
	if err != nil {
		return Config{}, fmt.Errorf("failed to get working directory: %w", err)
	}
	if err := godotenv.Load(filepath.Join(root, ".env")); err != nil {
		log.Printf("failed to load .env file: %v", err)
	}
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to process environment variables: %w", err)
	}
	return cfg, nil
}
