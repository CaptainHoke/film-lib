package services

import (
	"context"
	searchservice "film-lib/gen/search_service"
	"film-lib/internal/repo"
	"film-lib/internal/utils"
	"log"

	"goa.design/goa/v3/security"
)

// SearchService service example implementation.
// The example methods log the requests and return zero values.
type searchServicesrvc struct {
	logger *log.Logger
}

// NewSearchService returns the SearchService service implementation.
func NewSearchService(logger *log.Logger) searchservice.Service {
	return &searchServicesrvc{logger}
}

// JWTAuth implements the authorization logic for service "SearchService" for
// the "jwt" security scheme.
func (s *searchServicesrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return utils.JWTAuth(ctx, token, scheme, s.logger)
}

// SearchLibrary implements searchLibrary.
func (s *searchServicesrvc) SearchLibrary(ctx context.Context, p *searchservice.SearchLibraryPayload) (res *searchservice.Film, err error) {
	res = &searchservice.Film{}
	s.logger.Print("searchService.searchLibrary")
	return
}

// GetAllActors implements getAllActors.
func (s *searchServicesrvc) GetAllActors(ctx context.Context, p *searchservice.GetAllActorsPayload) (res searchservice.ActorWithFilmsResultCollection, err error) {
	s.logger.Print("searchService.getAllActors")

	actors := repo.PgInstance.GetAllActors(ctx, s.logger)

	for _, actor := range actors {
		s.logger.Printf("Actor: %v\n", actor)
		res = append(res, &searchservice.ActorWithFilmsResult{
			ActorID:        actor.ActorID,
			ActorName:      actor.ActorName,
			ActorSex:       actor.ActorSex,
			ActorBirthdate: actor.ActorBirthdate,
			FilmIDs:        actor.FilmIDs,
		})
	}
	err = nil

	return
}

// GetAllFilms implements getAllFilms.
func (s *searchServicesrvc) GetAllFilms(ctx context.Context, p *searchservice.GetAllFilmsPayload) (res searchservice.FilmResultCollection, err error) {
	s.logger.Print("searchService.getAllFilms")
	return
}
