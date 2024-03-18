// Code generated by goa v3.15.2, DO NOT EDIT.
//
// IsAlive HTTP client encoders and decoders
//
// Command:
// $ goa gen film-lib/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildCheckRequest instantiates a HTTP request object with method and path
// set to call the "IsAlive" service "check" endpoint
func (c *Client) BuildCheckRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CheckIsAlivePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("IsAlive", "check", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeCheckResponse returns a decoder for responses returned by the IsAlive
// check endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeCheckResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("IsAlive", "check", resp.StatusCode, string(body))
		}
	}
}