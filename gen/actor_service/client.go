// Code generated by goa v3.15.2, DO NOT EDIT.
//
// ActorService client
//
// Command:
// $ goa gen film-lib/design

package actorservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "ActorService" service client.
type Client struct {
	AddActorEndpoint        goa.Endpoint
	UpdateActorInfoEndpoint goa.Endpoint
	DeleteActorEndpoint     goa.Endpoint
}

// NewClient initializes a "ActorService" service client given the endpoints.
func NewClient(addActor, updateActorInfo, deleteActor goa.Endpoint) *Client {
	return &Client{
		AddActorEndpoint:        addActor,
		UpdateActorInfoEndpoint: updateActorInfo,
		DeleteActorEndpoint:     deleteActor,
	}
}

// AddActor calls the "addActor" endpoint of the "ActorService" service.
// AddActor may return the following errors:
//   - "invalid-scopes" (type InvalidScopes)
//   - "already-exists" (type *AlreadyExists): Actor already exists
//   - error: internal error
func (c *Client) AddActor(ctx context.Context, p *AddActorPayload) (res uint64, err error) {
	var ires any
	ires, err = c.AddActorEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(uint64), nil
}

// UpdateActorInfo calls the "updateActorInfo" endpoint of the "ActorService"
// service.
// UpdateActorInfo may return the following errors:
//   - "invalid-scopes" (type InvalidScopes)
//   - "not-found" (type *NotFound): Actor not found
//   - error: internal error
func (c *Client) UpdateActorInfo(ctx context.Context, p *UpdateActorInfoPayload) (err error) {
	_, err = c.UpdateActorInfoEndpoint(ctx, p)
	return
}

// DeleteActor calls the "deleteActor" endpoint of the "ActorService" service.
// DeleteActor may return the following errors:
//   - "invalid-scopes" (type InvalidScopes)
//   - "not-found" (type *NotFound): Actor not found
//   - error: internal error
func (c *Client) DeleteActor(ctx context.Context, p *DeleteActorPayload) (err error) {
	_, err = c.DeleteActorEndpoint(ctx, p)
	return
}
