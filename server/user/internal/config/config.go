package config

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Security SecurityConfig
}

type AppConfig struct {
	Environment string `envconfig:"APP_ENV" required:"true"`
	Port        string `envconfig:"APP_PORT" default:"8080" required:"true"`
}

type DatabaseConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	Database     string
	MaxConns     int32
	MinConns     int32
	MaxConnIdle  time.Duration
	MaxConnLife  time.Duration
	QueryTimeout time.Duration
}

type SecurityConfig struct {
	JWTSigningKey string
	JWTExpiration time.Duration
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

	if err = envconfig.Process("DB", &dest.Database); err != nil {
		return dest, err
	}

	return dest, err
}

func ProvidePGXConfig(c Config) (*pgxpool.Config, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=user_service",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.Database)

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("failed to parse pgxpool config: %s; connection string: %s", err.Error(), connString)
		return nil, err
	}

	poolConfig.MaxConns = c.Database.MaxConns
	poolConfig.MinConns = c.Database.MinConns
	poolConfig.MaxConnIdleTime = c.Database.MaxConnIdle
	poolConfig.MaxConnLifetime = c.Database.MaxConnLife

	return poolConfig, nil
}
