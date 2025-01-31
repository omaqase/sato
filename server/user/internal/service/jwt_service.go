package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/omaqase/sato/user/internal/config"
	"time"
)

type IJWTService interface {
	NewJWT(userID string, ttl time.Duration) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type JWTService struct {
	Config *config.SecurityConfig
}

func NewJWTService(config *config.SecurityConfig) IJWTService {
	return &JWTService{Config: config}
}

func (s JWTService) NewJWT(userID string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.Config.JWTExpiration)),
		Subject:   userID,
	})

	return token.SignedString([]byte(s.Config.JWTSigningKey))
}

func (s JWTService) Parse(accessToken string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s JWTService) NewRefreshToken() (string, error) {
	//TODO implement me
	panic("implement me")
}
