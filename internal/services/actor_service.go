package services

import (
	"context"
	actorservice "film-lib/gen/actor_service"
	"film-lib/internal/repo"
	"film-lib/internal/utils"
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
	return utils.JWTAuth(ctx, token, scheme, s.logger)
}

// AddActor implements addActor.
func (s *actorServicesrvc) AddActor(ctx context.Context, p *actorservice.AddActorPayload) (res uint64, err error) {
	s.logger.Print("actorService.addActor")

	res = repo.PgInstance.AddActor(ctx, *p.ActorInfo)
	err = nil

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
