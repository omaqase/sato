package session

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/omaqase/sato/user/internal/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

type IStore interface {
	SetSession(ctx context.Context, userID string, session domain.Session, ttl time.Duration) error
	GetSession(ctx context.Context, userID string) (*domain.Session, error)
	UpdateSession(ctx context.Context, userID string, session domain.Session, ttl time.Duration) error
	DeleteSession(ctx context.Context, userID string) error
	GetSessionByRefreshToken(ctx context.Context, refreshToken string) (*domain.Session, error)
}

type Store struct {
	rdb *redis.Client
}

func NewSessionStore(rdb *redis.Client) IStore {
	return &Store{rdb: rdb}
}

func (s Store) SetSession(ctx context.Context, userID string, session domain.Session, ttl time.Duration) error {
	userKey := fmt.Sprintf("session:%s", userID)
	refreshKey := fmt.Sprintf("refresh_token:%s", session.RefreshToken)
	sessionData, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to marshal session data: %w", err)
	}
	if err := s.rdb.Set(ctx, userKey, sessionData, ttl).Err(); err != nil {
		return fmt.Errorf("failed to set session for userID: %s: %w", userID, err)
	}
	if err := s.rdb.Set(ctx, refreshKey, sessionData, ttl).Err(); err != nil {
		return fmt.Errorf("failed to set refresh token for userID: %s: %w", userID, err)
	}
	return nil
}

func (s Store) GetSession(ctx context.Context, userID string) (*domain.Session, error) {
	key := fmt.Sprintf("session:%s", userID)
	sessionData, err := s.rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("session for userID: %s not found", userID)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get session for userID: %s: %w", userID, err)
	}
	var session domain.Session
	if err := json.Unmarshal([]byte(sessionData), &session); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session data for userID: %s: %w", userID, err)
	}
	return &session, nil
}

func (s Store) UpdateSession(ctx context.Context, userID string, session domain.Session, ttl time.Duration) error {
	key := fmt.Sprintf("session:%s", userID)
	if _, err := s.GetSession(ctx, userID); err != nil {
		if errors.Is(err, redis.Nil) {
			return fmt.Errorf("session for userID: %s not found", userID)
		}
		return fmt.Errorf("failed to retrieve session for userID: %s: %w", userID, err)
	}
	sessionData, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to marshal session data: %w", err)
	}
	if err := s.rdb.Set(ctx, key, sessionData, ttl).Err(); err != nil {
		return fmt.Errorf("failed to update session for userID: %s: %w", userID, err)
	}
	return nil
}

func (s Store) DeleteSession(ctx context.Context, userID string) error {
	userKey := fmt.Sprintf("session:%s", userID)
	sessionData, err := s.rdb.Get(ctx, userKey).Result()
	if errors.Is(err, redis.Nil) {
		return redis.Nil
	}
	if err != nil {
		return fmt.Errorf("failed to get session for userID: %s: %w", userID, err)
	}
	var session domain.Session
	if err := json.Unmarshal([]byte(sessionData), &session); err != nil {
		return fmt.Errorf("failed to unmarshal session data for userID: %s: %w", userID, err)
	}
	refreshKey := fmt.Sprintf("refresh_token:%s", session.RefreshToken)
	pipe := s.rdb.Pipeline()
	pipe.Del(ctx, userKey)
	pipe.Del(ctx, refreshKey)
	_, err = pipe.Exec(ctx)
	return err
}

func (s Store) GetSessionByRefreshToken(ctx context.Context, refreshToken string) (*domain.Session, error) {
	key := fmt.Sprintf("refresh_token:%s", refreshToken)
	sessionData, err := s.rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, redis.Nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get session by refresh token: %w", err)
	}
	var session domain.Session
	if err := json.Unmarshal([]byte(sessionData), &session); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session data for refresh token: %s: %w", refreshToken, err)
	}
	return &session, nil
}
