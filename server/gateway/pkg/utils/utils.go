package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

type TokenClaims struct {
	Sub  string `json:"sub"`
	Role string `json:"role"`
}

func ParseJWT(signingKey, accessToken string) (*TokenClaims, error) {
	if accessToken == "" || !strings.HasPrefix(accessToken, "Bearer ") {
		return nil, errors.New("invalid token prefix")
	}
	accessToken = strings.TrimPrefix(accessToken, "Bearer ")

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	tokenClaims := new(TokenClaims)

	sub, ok := claims["sub"].(string)
	if !ok || sub == "" {
		return nil, errors.New("subject (sub) not found in token claims")
	}

	tokenClaims.Sub = sub

	role, ok := claims["role"].(string)
	if !ok || role == "" {
		return nil, errors.New("role (role) not found in token claims")
	}

	tokenClaims.Role = role

	return tokenClaims, nil
}
