// Code generated by goa v3.15.2, DO NOT EDIT.
//
// FilmService client HTTP transport
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

// Client lists the FilmService service endpoint HTTP clients.
type Client struct {
	// GetAllFilms Doer is the HTTP client used to make requests to the getAllFilms
	// endpoint.
	GetAllFilmsDoer goahttp.Doer

	// AddFilm Doer is the HTTP client used to make requests to the addFilm
	// endpoint.
	AddFilmDoer goahttp.Doer

	// UpdateFilmInfo Doer is the HTTP client used to make requests to the
	// updateFilmInfo endpoint.
	UpdateFilmInfoDoer goahttp.Doer

	// DeleteFilm Doer is the HTTP client used to make requests to the deleteFilm
	// endpoint.
	DeleteFilmDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the FilmService service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetAllFilmsDoer:     doer,
		AddFilmDoer:         doer,
		UpdateFilmInfoDoer:  doer,
		DeleteFilmDoer:      doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// GetAllFilms returns an endpoint that makes HTTP requests to the FilmService
// service getAllFilms server.
func (c *Client) GetAllFilms() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetAllFilmsRequest(c.encoder)
		decodeResponse = DecodeGetAllFilmsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildGetAllFilmsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetAllFilmsDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("FilmService", "getAllFilms", err)
		}
		return decodeResponse(resp)
	}
}

// AddFilm returns an endpoint that makes HTTP requests to the FilmService
// service addFilm server.
func (c *Client) AddFilm() goa.Endpoint {
	var (
		encodeRequest  = EncodeAddFilmRequest(c.encoder)
		decodeResponse = DecodeAddFilmResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildAddFilmRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddFilmDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("FilmService", "addFilm", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateFilmInfo returns an endpoint that makes HTTP requests to the
// FilmService service updateFilmInfo server.
func (c *Client) UpdateFilmInfo() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateFilmInfoRequest(c.encoder)
		decodeResponse = DecodeUpdateFilmInfoResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildUpdateFilmInfoRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateFilmInfoDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("FilmService", "updateFilmInfo", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteFilm returns an endpoint that makes HTTP requests to the FilmService
// service deleteFilm server.
func (c *Client) DeleteFilm() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteFilmRequest(c.encoder)
		decodeResponse = DecodeDeleteFilmResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildDeleteFilmRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteFilmDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("FilmService", "deleteFilm", err)
		}
		return decodeResponse(resp)
	}
}
