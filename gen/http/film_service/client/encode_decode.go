// Code generated by goa v3.15.2, DO NOT EDIT.
//
// FilmService HTTP client encoders and decoders
//
// Command:
// $ goa gen film-lib/design

package client

import (
	"bytes"
	"context"
	filmservice "film-lib/gen/film_service"
	filmserviceviews "film-lib/gen/film_service/views"
	"io"
	"net/http"
	"net/url"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildAddFilmRequest instantiates a HTTP request object with method and path
// set to call the "FilmService" service "addFilm" endpoint
func (c *Client) BuildAddFilmRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddFilmFilmServicePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("FilmService", "addFilm", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddFilmRequest returns an encoder for requests sent to the FilmService
// addFilm server.
func EncodeAddFilmRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*filmservice.AddFilmPayload)
		if !ok {
			return goahttp.ErrInvalidType("FilmService", "addFilm", "*filmservice.AddFilmPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewAddFilmRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("FilmService", "addFilm", err)
		}
		return nil
	}
}

// DecodeAddFilmResponse returns a decoder for responses returned by the
// FilmService addFilm endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeAddFilmResponse may return the following errors:
//   - "already-exists" (type *filmservice.AlreadyExists): http.StatusBadRequest
//   - "invalid-scopes" (type filmservice.InvalidScopes): http.StatusForbidden
//   - "unauthorized" (type filmservice.Unauthorized): http.StatusUnauthorized
//   - error: internal error
func DecodeAddFilmResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body AddFilmResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("FilmService", "addFilm", err)
			}
			p := NewAddFilmFilmResultCreated(&body)
			view := "default"
			vres := &filmserviceviews.FilmResult{Projected: p, View: view}
			if err = filmserviceviews.ValidateFilmResult(vres); err != nil {
				return nil, goahttp.ErrValidationError("FilmService", "addFilm", err)
			}
			res := filmservice.NewFilmResult(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body AddFilmAlreadyExistsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("FilmService", "addFilm", err)
			}
			err = ValidateAddFilmAlreadyExistsResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("FilmService", "addFilm", err)
			}
			return nil, NewAddFilmAlreadyExists(&body)
		case http.StatusForbidden:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("FilmService", "addFilm", err)
			}
			return nil, NewAddFilmInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("FilmService", "addFilm", err)
			}
			return nil, NewAddFilmUnauthorized(body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("FilmService", "addFilm", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateFilmInfoRequest instantiates a HTTP request object with method
// and path set to call the "FilmService" service "updateFilmInfo" endpoint
func (c *Client) BuildUpdateFilmInfoRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		filmID uint64
	)
	{
		p, ok := v.(*filmservice.UpdateFilmInfoPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("FilmService", "updateFilmInfo", "*filmservice.UpdateFilmInfoPayload", v)
		}
		if p.FilmID != nil {
			filmID = *p.FilmID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateFilmInfoFilmServicePath(filmID)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("FilmService", "updateFilmInfo", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateFilmInfoRequest returns an encoder for requests sent to the
// FilmService updateFilmInfo server.
func EncodeUpdateFilmInfoRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*filmservice.UpdateFilmInfoPayload)
		if !ok {
			return goahttp.ErrInvalidType("FilmService", "updateFilmInfo", "*filmservice.UpdateFilmInfoPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewUpdateFilmInfoRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("FilmService", "updateFilmInfo", err)
		}
		return nil
	}
}

// DecodeUpdateFilmInfoResponse returns a decoder for responses returned by the
// FilmService updateFilmInfo endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeUpdateFilmInfoResponse may return the following errors:
//   - "not-found" (type *filmservice.NotFound): http.StatusNoContent
//   - "invalid-scopes" (type filmservice.InvalidScopes): http.StatusForbidden
//   - "unauthorized" (type filmservice.Unauthorized): http.StatusUnauthorized
//   - error: internal error
func DecodeUpdateFilmInfoResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			return nil, nil
		case http.StatusNoContent:
			var (
				body UpdateFilmInfoNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("FilmService", "updateFilmInfo", err)
			}
			err = ValidateUpdateFilmInfoNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("FilmService", "updateFilmInfo", err)
			}
			return nil, NewUpdateFilmInfoNotFound(&body)
		case http.StatusForbidden:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("FilmService", "updateFilmInfo", err)
			}
			return nil, NewUpdateFilmInfoInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("FilmService", "updateFilmInfo", err)
			}
			return nil, NewUpdateFilmInfoUnauthorized(body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("FilmService", "updateFilmInfo", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteFilmRequest instantiates a HTTP request object with method and
// path set to call the "FilmService" service "deleteFilm" endpoint
func (c *Client) BuildDeleteFilmRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		filmID uint64
	)
	{
		p, ok := v.(*filmservice.DeleteFilmPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("FilmService", "deleteFilm", "*filmservice.DeleteFilmPayload", v)
		}
		filmID = p.FilmID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteFilmFilmServicePath(filmID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("FilmService", "deleteFilm", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteFilmRequest returns an encoder for requests sent to the
// FilmService deleteFilm server.
func EncodeDeleteFilmRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*filmservice.DeleteFilmPayload)
		if !ok {
			return goahttp.ErrInvalidType("FilmService", "deleteFilm", "*filmservice.DeleteFilmPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeDeleteFilmResponse returns a decoder for responses returned by the
// FilmService deleteFilm endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeDeleteFilmResponse may return the following errors:
//   - "invalid-scopes" (type filmservice.InvalidScopes): http.StatusForbidden
//   - "unauthorized" (type filmservice.Unauthorized): http.StatusUnauthorized
//   - error: internal error
func DecodeDeleteFilmResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusForbidden:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("FilmService", "deleteFilm", err)
			}
			return nil, NewDeleteFilmInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("FilmService", "deleteFilm", err)
			}
			return nil, NewDeleteFilmUnauthorized(body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("FilmService", "deleteFilm", resp.StatusCode, string(body))
		}
	}
}

// marshalFilmserviceFilmInfoToFilmInfoRequestBody builds a value of type
// *FilmInfoRequestBody from a value of type *filmservice.FilmInfo.
func marshalFilmserviceFilmInfoToFilmInfoRequestBody(v *filmservice.FilmInfo) *FilmInfoRequestBody {
	res := &FilmInfoRequestBody{
		Title:       v.Title,
		Description: v.Description,
		ReleaseDate: v.ReleaseDate,
		Rating:      v.Rating,
	}
	if v.Actors != nil {
		res.Actors = make([]uint64, len(v.Actors))
		for i, val := range v.Actors {
			res.Actors[i] = val
		}
	} else {
		res.Actors = []uint64{}
	}

	return res
}

// marshalFilmInfoRequestBodyToFilmserviceFilmInfo builds a value of type
// *filmservice.FilmInfo from a value of type *FilmInfoRequestBody.
func marshalFilmInfoRequestBodyToFilmserviceFilmInfo(v *FilmInfoRequestBody) *filmservice.FilmInfo {
	res := &filmservice.FilmInfo{
		Title:       v.Title,
		Description: v.Description,
		ReleaseDate: v.ReleaseDate,
		Rating:      v.Rating,
	}
	if v.Actors != nil {
		res.Actors = make([]uint64, len(v.Actors))
		for i, val := range v.Actors {
			res.Actors[i] = val
		}
	} else {
		res.Actors = []uint64{}
	}

	return res
}
