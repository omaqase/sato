package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/omaqase/sato/user/internal/config"
	"time"
)

type IJWTService interface {
	NewJWT(userID, role string, ttl time.Duration) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type JWTService struct {
	Config config.Config
}

func NewJWTService(config config.Config) IJWTService {
	return &JWTService{Config: config}
}

type CustomClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func (s JWTService) NewJWT(userID, role string, ttl time.Duration) (string, error) {
	claims := CustomClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			Subject:   userID,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.Config.Security.JWTSigningKey))
}

func (s JWTService) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.Config.Security.JWTSigningKey), nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid or expired token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}
	subject, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("missing subject in claims")
	}
	return subject, nil
}

func (s JWTService) NewRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.New("failed to generate refresh token")
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
