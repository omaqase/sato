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
	UserService         UserServiceConfig
	Security            SecurityConfig
}

type AppConfig struct {
	Environment string
	Port        string
}

type NotificationServiceConfig struct {
	Host string
	Port string
}

type UserServiceConfig struct {
	Host string
	Port string
}

type SecurityConfig struct {
	SigningKey string
}

func Load() (dest Config, err error) {
	root, err := os.Getwd()
	if err != nil {
		return dest, err
	}

	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		return dest, err
	}

	dest = Config{}

	if err = envconfig.Process("APP", &dest.App); err != nil {
		return dest, err
	}

	if err = envconfig.Process("NOTIFICATION_SERVICE", &dest.NotificationService); err != nil {
		return dest, err
	}

	if err = envconfig.Process("USER_SERVICE", &dest.UserService); err != nil {
		return dest, err
	}

	if err = envconfig.Process("SECURITY", &dest.Security); err != nil {
		return dest, err
	}

	return dest, err
}
