package filmlib

import (
	"context"
	searchservice "film-lib/gen/search_service"
	"fmt"
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

// SearchLibrary implements searchLibrary.
func (s *searchServicesrvc) SearchLibrary(ctx context.Context, p *searchservice.SearchLibraryPayload) (res *searchservice.Film, err error) {
	res = &searchservice.Film{}
	s.logger.Print("searchService.searchLibrary")
	return
}

// GetAllActors implements getAllActors.
func (s *searchServicesrvc) GetAllActors(ctx context.Context, p *searchservice.GetAllActorsPayload) (res searchservice.ActorResultCollection, err error) {
	s.logger.Print("searchService.getAllActors")
	return
}

// GetAllFilms implements getAllFilms.
func (s *searchServicesrvc) GetAllFilms(ctx context.Context, p *searchservice.GetAllFilmsPayload) (res searchservice.FilmResultCollection, err error) {
	s.logger.Print("searchService.getAllFilms")
	return
}
