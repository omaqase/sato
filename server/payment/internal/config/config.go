package config

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Security SecurityConfig
}

type AppConfig struct {
	Environment string `envconfig:"APP_ENV" required:"true"`
	Port        string `envconfig:"APP_PORT" default:"5056" required:"true"`
	Name        string `envconfig:"APP_NAME" default:"payment-service"`
}

type DatabaseConfig struct {
	Host         string        `envconfig:"DB_HOST" required:"true"`
	Port         int           `envconfig:"DB_PORT" default:"5432" required:"true"`
	User         string        `envconfig:"DB_USER" required:"true"`
	Password     string        `envconfig:"DB_PASSWORD" required:"true"`
	Database     string        `envconfig:"DB_NAME" required:"true"`
	SSLMode      string        `envconfig:"DB_SSLMODE" default:"disable"`
	MaxConns     int32         `envconfig:"DB_MAXCONNS" default:"10"`
	MinConns     int32         `envconfig:"DB_MINCONNS" default:"2"`
	MaxConnIdle  time.Duration `envconfig:"DB_MAXCONN_IDLE" default:"30m"`
	MaxConnLife  time.Duration `envconfig:"DB_MAXCONNLIFE" default:"1h"`
	QueryTimeout time.Duration `envconfig:"DB_QUERYTIMEOUT" default:"30s"`
}

type SecurityConfig struct {
	EncryptionKey []byte
}

func Load() (*Config, error) {
	root, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	if err := godotenv.Load(filepath.Join(root, ".env")); err != nil {
		log.Printf("failed to load .env file: %v", err)
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("failed to process environment variables: %w", err)
	}

	encKey := os.Getenv("PAYMENT_ENCRYPTION_KEY")
	if encKey == "" {
		return nil, fmt.Errorf("PAYMENT_ENCRYPTION_KEY is not set")
	}

	key, err := base64.StdEncoding.DecodeString(encKey)
	if err != nil {
		return nil, fmt.Errorf("invalid encryption key format: %w", err)
	}

	if len(key) != 32 {
		return nil, fmt.Errorf("encryption key must be 32 bytes")
	}

	cfg.Security.EncryptionKey = key

	return &cfg, nil
}

func (c *Config) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=payment_management",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.Database, c.Database.SSLMode)
}

func (c *Config) GetEncryptionKey() []byte {
	return c.Security.EncryptionKey
}
