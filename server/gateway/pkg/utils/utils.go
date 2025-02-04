package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func ParseJWT(signingKey, accessToken string) (string, error) {
	if accessToken == "" || !strings.HasPrefix(accessToken, "Bearer ") {
		return "", errors.New("invalid token prefix")
	}
	accessToken = strings.TrimPrefix(accessToken, "Bearer ")

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	sub, ok := claims["sub"].(string)
	if !ok || sub == "" {
		return "", errors.New("subject (sub) not found in token claims")
	}

	return sub, nil
}
