package services

import (
	"context"
	searchservice "film-lib/gen/search_service"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"

	"goa.design/goa/v3/security"
)

// SearchService service example implementation.
// The example methods log the requests and return zero values.
type searchService struct {
	logger *log.Logger
}

// NewSearchService returns the SearchService service implementation.
func NewSearchService(logger *log.Logger) searchservice.Service {
	return &searchService{logger}
}

// JWTAuth implements the authorization logic for service "SearchService" for
// the "jwt" security scheme.
// TODO: DRY
func (s *searchService) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token,
		claims, func(_ *jwt.Token) (interface{}, error) { return JWTKey, nil })
	if err != nil {
		s.logger.Print("Unable to obtain claim from token, it's invalid")
		return ctx, searchservice.Unauthorized("invalid token")
	}

	s.logger.Print("claims retrieved, validating against scope")
	s.logger.Print(claims)

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		s.logger.Print("Unable to get scope since the scope is empty")
		return ctx, searchservice.InvalidScopes("invalid scopes in token")
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		s.logger.Print("An error occurred when retrieving the scopes")
		s.logger.Print(ok)
		return ctx, searchservice.InvalidScopes("invalid scopes in token")
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := scheme.Validate(scopesInToken); err != nil {
		s.logger.Print("Unable to parse token, check error below")
		return ctx, searchservice.InvalidScopes(err.Error())
	}

	return ctx, fmt.Errorf("not implemented")
}

// SearchLibrary implements searchLibrary.
func (s *searchService) SearchLibrary(ctx context.Context, p *searchservice.SearchLibraryPayload) (res *searchservice.Film, err error) {
	res = &searchservice.Film{}
	s.logger.Print("searchService.searchLibrary")
	return
}
