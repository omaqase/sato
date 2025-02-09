package utils

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"strings"
)

func ParseMetadata(ctx context.Context) (string, error) {
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return "", fmt.Errorf("no metadata")
	}

	authHeaders, ok := md["grpcgateway-authorization"]
	if !ok || len(authHeaders) == 0 {
		return "", fmt.Errorf("no authorization header")
	}

	authHeader := authHeaders[0]
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("authorization header is not bearer token")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	return token, nil
}
