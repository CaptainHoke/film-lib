package services

import (
	"context"
	isalive "film-lib/gen/is_alive"
	"log"
)

// IsAlive service example implementation.
// The example methods log the requests and return zero values.
type isAlivesrvc struct {
	logger *log.Logger
}

// NewIsAlive returns the IsAlive service implementation.
func NewIsAlive(logger *log.Logger) isalive.Service {
	return &isAlivesrvc{logger}
}

// Check implements check.
func (s *isAlivesrvc) Check(ctx context.Context) (err error) {
	s.logger.Print("isAlive.check")
	return
}
