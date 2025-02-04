package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type ICache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (interface{}, error)
	Del(ctx context.Context, key string) error
}

type Cache struct {
	redis *redis.Client
}

func NewCache(redis *redis.Client) ICache {
	return &Cache{redis: redis}
}

func (c Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	marshaled, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("error while marshaling payload: %v", err)
	}

	if err := c.redis.Set(ctx, key, string(marshaled), expiration).Err(); err != nil {
		return fmt.Errorf("error while setting cache: %v", err)
	}

	return nil
}

func (c Cache) Get(ctx context.Context, key string) (interface{}, error) {
	marshaled, err := c.redis.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("cache with key %s not found: %v", key, err)
		}
		return nil, fmt.Errorf("error while getting cache: %v", err)
	}

	var result interface{}
	if err := json.Unmarshal([]byte(marshaled), &result); err != nil {
		return nil, fmt.Errorf("error while unmarshaling cache: %v", err)
	}

	return result, nil
}

func (c Cache) Del(ctx context.Context, key string) error {
	if err := c.redis.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("error while deleting cache: %v", err)
	}

	return nil
}
