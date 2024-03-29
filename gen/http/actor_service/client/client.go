// Code generated by goa v3.15.2, DO NOT EDIT.
//
// ActorService client HTTP transport
//
// Command:
// $ goa gen film-lib/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the ActorService service endpoint HTTP clients.
type Client struct {
	// AddActor Doer is the HTTP client used to make requests to the addActor
	// endpoint.
	AddActorDoer goahttp.Doer

	// UpdateActorInfo Doer is the HTTP client used to make requests to the
	// updateActorInfo endpoint.
	UpdateActorInfoDoer goahttp.Doer

	// DeleteActor Doer is the HTTP client used to make requests to the deleteActor
	// endpoint.
	DeleteActorDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the ActorService service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		AddActorDoer:        doer,
		UpdateActorInfoDoer: doer,
		DeleteActorDoer:     doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// AddActor returns an endpoint that makes HTTP requests to the ActorService
// service addActor server.
func (c *Client) AddActor() goa.Endpoint {
	var (
		encodeRequest  = EncodeAddActorRequest(c.encoder)
		decodeResponse = DecodeAddActorResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildAddActorRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddActorDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("ActorService", "addActor", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateActorInfo returns an endpoint that makes HTTP requests to the
// ActorService service updateActorInfo server.
func (c *Client) UpdateActorInfo() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateActorInfoRequest(c.encoder)
		decodeResponse = DecodeUpdateActorInfoResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildUpdateActorInfoRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateActorInfoDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("ActorService", "updateActorInfo", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteActor returns an endpoint that makes HTTP requests to the ActorService
// service deleteActor server.
func (c *Client) DeleteActor() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteActorRequest(c.encoder)
		decodeResponse = DecodeDeleteActorResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildDeleteActorRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteActorDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("ActorService", "deleteActor", err)
		}
		return decodeResponse(resp)
	}
}
