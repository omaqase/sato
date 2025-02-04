package config

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Security SecurityConfig
}

type AppConfig struct {
	Environment string `envconfig:"APP_ENV" required:"true"`
	Port        string `envconfig:"APP_PORT" default:"8080" required:"true"`
}

type DatabaseConfig struct {
	Host         string        `envconfig:"DB_HOST" required:"true"`
	Port         int           `envconfig:"DB_PORT" default:"5432" required:"true"`
	User         string        `envconfig:"DB_USER" required:"true"`
	Password     string        `envconfig:"DB_PASSWORD" required:"true"`
	Database     string        `envconfig:"DB_DATABASE" required:"true"`
	MaxConns     int32         `envconfig:"DB_MAXCONNS" default:"10"`
	MinConns     int32         `envconfig:"DB_MINCONNS" default:"5"`
	MaxConnIdle  time.Duration `envconfig:"DB_MAXCONN_IDLE" default:"5m"`
	MaxConnLife  time.Duration `envconfig:"DB_MAXCONNLIFE" default:"30m"`
	QueryTimeout time.Duration `envconfig:"DB_QUERYTIMEOUT" default:"10s"`
}

type SecurityConfig struct {
	JWTSigningKey          string        `envconfig:"JWT_SIGNING_KEY" required:"true"`
	JWTExpiration          time.Duration `envconfig:"JWT_EXPIRATION" default:"24h"`
	RefreshTokenExpiration time.Duration `envconfig:"REFRESH_TOKEN_EXPIRATION" default:"720h"` // 30 days
}

type RedisConfig struct {
	Addr     string `envconfig:"REDIS_ADDR" required:"true"`
	Password string `envconfig:"REDIS_PASSWORD" default:""`
	DB       int    `envconfig:"REDIS_DB" default:"0"`
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

func ProvidePGXConfig(c Config) (*pgxpool.Config, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=user_service",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.Database)
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pgxpool config: %w; connection string: %s", err, connString)
	}
	poolConfig.MaxConns = c.Database.MaxConns
	poolConfig.MinConns = c.Database.MinConns
	poolConfig.MaxConnIdleTime = c.Database.MaxConnIdle
	poolConfig.MaxConnLifetime = c.Database.MaxConnLife
	return poolConfig, nil
}

func ProvideRedisConfig(c Config) *redis.Options {
	return &redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	}
}
