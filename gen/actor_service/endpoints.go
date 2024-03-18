// Code generated by goa v3.15.2, DO NOT EDIT.
//
// ActorService endpoints
//
// Command:
// $ goa gen film-lib/design

package actorservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "ActorService" service endpoints.
type Endpoints struct {
	GetAllActors    goa.Endpoint
	AddActor        goa.Endpoint
	UpdateActorInfo goa.Endpoint
	DeleteActor     goa.Endpoint
}

// NewEndpoints wraps the methods of the "ActorService" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		GetAllActors:    NewGetAllActorsEndpoint(s, a.JWTAuth),
		AddActor:        NewAddActorEndpoint(s, a.JWTAuth),
		UpdateActorInfo: NewUpdateActorInfoEndpoint(s, a.JWTAuth),
		DeleteActor:     NewDeleteActorEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "ActorService" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetAllActors = m(e.GetAllActors)
	e.AddActor = m(e.AddActor)
	e.UpdateActorInfo = m(e.UpdateActorInfo)
	e.DeleteActor = m(e.DeleteActor)
}

// NewGetAllActorsEndpoint returns an endpoint function that calls the method
// "getAllActors" of service "ActorService".
func NewGetAllActorsEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetAllActorsPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{"api:read"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.GetAllActors(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedActorResultCollection(res, "default")
		return vres, nil
	}
}

// NewAddActorEndpoint returns an endpoint function that calls the method
// "addActor" of service "ActorService".
func NewAddActorEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*AddActorPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{"api:write"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.AddActor(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedActorResult(res, "default")
		return vres, nil
	}
}

// NewUpdateActorInfoEndpoint returns an endpoint function that calls the
// method "updateActorInfo" of service "ActorService".
func NewUpdateActorInfoEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdateActorInfoPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{"api:write"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.UpdateActorInfo(ctx, p)
	}
}

// NewDeleteActorEndpoint returns an endpoint function that calls the method
// "deleteActor" of service "ActorService".
func NewDeleteActorEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeleteActorPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{"api:write"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.DeleteActor(ctx, p)
	}
}
