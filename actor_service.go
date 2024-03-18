package filmlib

import (
	"context"
	actorservice "film-lib/gen/actor_service"
	"github.com/golang-jwt/jwt/v5"
	"log"

	"goa.design/goa/v3/security"
)

// ActorService service example implementation.
// The example methods log the requests and return zero values.
type actorService struct {
	logger *log.Logger
}

// NewActorService returns the ActorService service implementation.
func NewActorService(logger *log.Logger) actorservice.Service {
	return &actorService{logger}
}

// JWTAuth implements the authorization logic for service "ActorService" for
// the "jwt" security scheme.
// TODO: DRY
func (s *actorService) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token,
		claims, func(_ *jwt.Token) (interface{}, error) { return JWTKey, nil })
	if err != nil {
		s.logger.Print("Unable to obtain claim from token, it's invalid")
		return ctx, actorservice.Unauthorized("invalid token")
	}

	s.logger.Print("claims retrieved, validating against scope")
	s.logger.Print(claims)

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		s.logger.Print("Unable to get scope since the scope is empty")
		return ctx, actorservice.InvalidScopes("invalid scopes in token")
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		s.logger.Print("An error occurred when retrieving the scopes")
		s.logger.Print(ok)
		return ctx, actorservice.InvalidScopes("invalid scopes in token")
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := scheme.Validate(scopesInToken); err != nil {
		s.logger.Print("Unable to parse token, check error below")
		return ctx, actorservice.InvalidScopes(err.Error())
	}

	return ctx, nil
}

// GetAllActors implements getAllActors.
func (s *actorService) GetAllActors(ctx context.Context, p *actorservice.GetAllActorsPayload) (res actorservice.ActorResultCollection, err error) {
	s.logger.Print("actorService.getAllActors")
	return
}

// AddActor implements addActor.
func (s *actorService) AddActor(ctx context.Context, p *actorservice.AddActorPayload) (res *actorservice.ActorResult, err error) {
	res = &actorservice.ActorResult{}
	s.logger.Print("actorService.addActor")
	return
}

// UpdateActorInfo implements updateActorInfo.
func (s *actorService) UpdateActorInfo(ctx context.Context, p *actorservice.UpdateActorInfoPayload) (err error) {
	s.logger.Print("actorService.updateActorInfo")
	return
}

// DeleteActor implements deleteActor.
func (s *actorService) DeleteActor(ctx context.Context, p *actorservice.DeleteActorPayload) (err error) {
	s.logger.Print("actorService.deleteActor")
	return
}
