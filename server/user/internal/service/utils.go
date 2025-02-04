package service

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
)

const AuthorizationPrefix = "grpcgateway-authorization"

func (s UsersService) ExtractUserIDFromMetadata(ctx context.Context) (string, bool) {
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return "", false
	}

	authHeaders, ok := md[AuthorizationPrefix]
	if !ok || len(authHeaders) == 0 {
		return "", false
	}

	authHeader := authHeaders[0]
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", false
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := s.jwtService.Parse(token)
	if err != nil {
		return "", false
	}

	return userID, true
}
