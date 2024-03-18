package services

import (
	"context"
	filmservice "film-lib/gen/film_service"
	"github.com/golang-jwt/jwt/v5"
	"log"

	"goa.design/goa/v3/security"
)

// FilmService service example implementation.
// The example methods log the requests and return zero values.
type filmService struct {
	logger *log.Logger
}

// NewFilmService returns the FilmService service implementation.
func NewFilmService(logger *log.Logger) filmservice.Service {
	return &filmService{logger}
}

// JWTAuth implements the authorization logic for service "FilmService" for the
// "jwt" security scheme.
// TODO: DRY
func (s *filmService) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token,
		claims, func(_ *jwt.Token) (interface{}, error) { return JWTKey, nil })
	if err != nil {
		s.logger.Print("Unable to obtain claim from token, it's invalid")
		return ctx, filmservice.Unauthorized("invalid token")
	}

	s.logger.Print("claims retrieved, validating against scope")
	s.logger.Print(claims)

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		s.logger.Print("Unable to get scope since the scope is empty")
		return ctx, filmservice.InvalidScopes("invalid scopes in token")
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		s.logger.Print("An error occurred when retrieving the scopes")
		s.logger.Print(ok)
		return ctx, filmservice.InvalidScopes("invalid scopes in token")
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := scheme.Validate(scopesInToken); err != nil {
		s.logger.Print("Unable to parse token, check error below")
		return ctx, filmservice.InvalidScopes(err.Error())
	}

	return ctx, nil
}

// GetAllFilms implements getAllFilms.
func (s *filmService) GetAllFilms(ctx context.Context, p *filmservice.GetAllFilmsPayload) (res filmservice.FilmResultCollection, err error) {
	s.logger.Print("filmService.getAllFilms")
	return
}

// AddFilm implements addFilm.
func (s *filmService) AddFilm(ctx context.Context, p *filmservice.AddFilmPayload) (res *filmservice.FilmResult, err error) {
	res = &filmservice.FilmResult{}
	s.logger.Print("filmService.addFilm")
	return
}

// UpdateFilmInfo implements updateFilmInfo.
func (s *filmService) UpdateFilmInfo(ctx context.Context, p *filmservice.UpdateFilmInfoPayload) (err error) {
	s.logger.Print("filmService.updateFilmInfo")
	return
}

// DeleteFilm implements deleteFilm.
func (s *filmService) DeleteFilm(ctx context.Context, p *filmservice.DeleteFilmPayload) (err error) {
	s.logger.Print("filmService.deleteFilm")
	return
}
