package services

import (
	"context"
	isalivesrvc "film-lib/gen/is_alive"
	"log"
)

// SearchService service example implementation.
// The example methods log the requests and return zero values.
type isAliveService struct {
	logger *log.Logger
}

// NewIsAliveService returns the SearchService service implementation.
func NewIsAliveService(logger *log.Logger) isalivesrvc.Service {
	return &isAliveService{logger}
}

func (s *isAliveService) Check(ctx context.Context) error {
	s.logger.Print("isAliveService.check")
	return nil
}
