package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App                 AppConfig
	NotificationService NotificationServiceConfig
	Security            SecurityConfig
}

type AppConfig struct {
	Environment string `envconfig:"APP_ENV" required:"true"`
	Port        string `envconfig:"APP_PORT" default:"8080" required:"true"`
}

type NotificationServiceConfig struct {
	Host string `envconfig:"NOTIFICATION_SERVICE_HOST"`
	Port string `envconfig:"NOTIFICATION_SERVICE_PORT" default:"5050"`
}

type SecurityConfig struct {
	SigningKey string `envconfig:"SECURITY_SIGNING_KEY"`
}

func Load() (dest *Config, err error) {
	root, err := os.Getwd()
	if err != nil {
		return dest, err
	}

	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		return dest, err
	}

	if err = envconfig.Process("APP", &dest.App); err != nil {
		return dest, err
	}

	if err = envconfig.Process("NOTIFICATION_SERVICE", &dest.NotificationService); err != nil {
		return dest, err
	}

	if err = envconfig.Process("SECURITY", &dest.Security); err != nil {
		return dest, err
	}

	return dest, err
}
