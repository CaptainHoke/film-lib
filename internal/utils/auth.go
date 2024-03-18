package utils

import (
	"context"
	searchservice "film-lib/gen/search_service"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"goa.design/goa/v3/security"
	"log"
)

var (
	// JWTKey is the key used in JWT authentication
	JWTKey = []byte("TODO: FIXME")
)

func JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme, logger *log.Logger) (context.Context, error) {
	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token,
		claims, func(_ *jwt.Token) (interface{}, error) { return JWTKey, nil })
	if err != nil {
		logger.Print("Unable to obtain claim from token, it's invalid")
		return ctx, searchservice.Unauthorized("invalid token")
	}

	logger.Print("claims retrieved, validating against scope")
	logger.Print(claims)

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		logger.Print("Unable to get scope since the scope is empty")
		return ctx, searchservice.InvalidScopes("invalid scopes in token")
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		logger.Print("An error occurred when retrieving the scopes")
		logger.Print(ok)
		return ctx, searchservice.InvalidScopes("invalid scopes in token")
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := scheme.Validate(scopesInToken); err != nil {
		logger.Print("Unable to parse token, check error below")
		return ctx, searchservice.InvalidScopes(err.Error())
	}

	return ctx, fmt.Errorf("not implemented")
}
