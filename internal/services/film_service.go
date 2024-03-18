package services

import (
	"context"
	filmservice "film-lib/gen/film_service"
	"fmt"
	"log"

	"goa.design/goa/v3/security"
)

// FilmService service example implementation.
// The example methods log the requests and return zero values.
type filmServicesrvc struct {
	logger *log.Logger
}

// NewFilmService returns the FilmService service implementation.
func NewFilmService(logger *log.Logger) filmservice.Service {
	return &filmServicesrvc{logger}
}

// JWTAuth implements the authorization logic for service "FilmService" for the
// "jwt" security scheme.
func (s *filmServicesrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// AddFilm implements addFilm.
func (s *filmServicesrvc) AddFilm(ctx context.Context, p *filmservice.AddFilmPayload) (res *filmservice.FilmResult, err error) {
	res = &filmservice.FilmResult{}
	s.logger.Print("filmService.addFilm")
	return
}

// UpdateFilmInfo implements updateFilmInfo.
func (s *filmServicesrvc) UpdateFilmInfo(ctx context.Context, p *filmservice.UpdateFilmInfoPayload) (err error) {
	s.logger.Print("filmService.updateFilmInfo")
	return
}

// DeleteFilm implements deleteFilm.
func (s *filmServicesrvc) DeleteFilm(ctx context.Context, p *filmservice.DeleteFilmPayload) (err error) {
	s.logger.Print("filmService.deleteFilm")
	return
}
