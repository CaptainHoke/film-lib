package services

import (
	"context"
	actorservice "film-lib/gen/actor_service"
	"fmt"
	"log"

	"goa.design/goa/v3/security"
)

// ActorService service example implementation.
// The example methods log the requests and return zero values.
type actorServicesrvc struct {
	logger *log.Logger
}

// NewActorService returns the ActorService service implementation.
func NewActorService(logger *log.Logger) actorservice.Service {
	return &actorServicesrvc{logger}
}

// JWTAuth implements the authorization logic for service "ActorService" for
// the "jwt" security scheme.
func (s *actorServicesrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
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

// AddActor implements addActor.
func (s *actorServicesrvc) AddActor(ctx context.Context, p *actorservice.AddActorPayload) (res uint64, err error) {
	s.logger.Print("actorService.addActor")
	return
}

// UpdateActorInfo implements updateActorInfo.
func (s *actorServicesrvc) UpdateActorInfo(ctx context.Context, p *actorservice.UpdateActorInfoPayload) (err error) {
	s.logger.Print("actorService.updateActorInfo")
	return
}

// DeleteActor implements deleteActor.
func (s *actorServicesrvc) DeleteActor(ctx context.Context, p *actorservice.DeleteActorPayload) (err error) {
	s.logger.Print("actorService.deleteActor")
	return
}
