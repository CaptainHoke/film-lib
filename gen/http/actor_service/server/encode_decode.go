// Code generated by goa v3.15.2, DO NOT EDIT.
//
// ActorService HTTP server encoders and decoders
//
// Command:
// $ goa gen film-lib/design

package server

import (
	"context"
	"errors"
	actorservice "film-lib/gen/actor_service"
	actorserviceviews "film-lib/gen/actor_service/views"
	"io"
	"net/http"
	"strconv"
	"strings"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetAllActorsResponse returns an encoder for responses returned by the
// ActorService getAllActors endpoint.
func EncodeGetAllActorsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(actorserviceviews.ActorResultCollection)
		enc := encoder(ctx, w)
		body := NewActorResultResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetAllActorsRequest returns a decoder for requests sent to the
// ActorService getAllActors endpoint.
func DecodeGetAllActorsRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			token string
			err   error
		)
		token = r.Header.Get("X-Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetAllActorsPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeGetAllActorsError returns an encoder for errors returned by the
// getAllActors ActorService endpoint.
func EncodeGetAllActorsError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "invalid-scopes":
			var res actorservice.InvalidScopes
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "unauthorized":
			var res actorservice.Unauthorized
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeAddActorResponse returns an encoder for responses returned by the
// ActorService addActor endpoint.
func EncodeAddActorResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(*actorserviceviews.ActorResult)
		enc := encoder(ctx, w)
		body := NewAddActorResponseBody(res.Projected)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeAddActorRequest returns a decoder for requests sent to the
// ActorService addActor endpoint.
func DecodeAddActorRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body AddActorRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAddActorRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			token string
		)
		token = r.Header.Get("X-Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewAddActorPayload(&body, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeAddActorError returns an encoder for errors returned by the addActor
// ActorService endpoint.
func EncodeAddActorError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "already-exists":
			var res *actorservice.AlreadyExists
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewAddActorAlreadyExistsResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "invalid-scopes":
			var res actorservice.InvalidScopes
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "unauthorized":
			var res actorservice.Unauthorized
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateActorInfoResponse returns an encoder for responses returned by
// the ActorService updateActorInfo endpoint.
func EncodeUpdateActorInfoResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

// DecodeUpdateActorInfoRequest returns a decoder for requests sent to the
// ActorService updateActorInfo endpoint.
func DecodeUpdateActorInfoRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body UpdateActorInfoRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateActorInfoRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			actorID uint64
			token   string

			params = mux.Vars(r)
		)
		{
			actorIDRaw := params["ActorID"]
			v, err2 := strconv.ParseUint(actorIDRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("ActorID", actorIDRaw, "unsigned integer"))
			}
			actorID = v
		}
		token = r.Header.Get("X-Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdateActorInfoPayload(&body, actorID, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeUpdateActorInfoError returns an encoder for errors returned by the
// updateActorInfo ActorService endpoint.
func EncodeUpdateActorInfoError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "invalid-scopes":
			var res actorservice.InvalidScopes
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "unauthorized":
			var res actorservice.Unauthorized
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteActorResponse returns an encoder for responses returned by the
// ActorService deleteActor endpoint.
func EncodeDeleteActorResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeDeleteActorRequest returns a decoder for requests sent to the
// ActorService deleteActor endpoint.
func DecodeDeleteActorRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			actorID uint64
			token   string
			err     error

			params = mux.Vars(r)
		)
		{
			actorIDRaw := params["ActorID"]
			v, err2 := strconv.ParseUint(actorIDRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("ActorID", actorIDRaw, "unsigned integer"))
			}
			actorID = v
		}
		token = r.Header.Get("X-Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteActorPayload(actorID, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeDeleteActorError returns an encoder for errors returned by the
// deleteActor ActorService endpoint.
func EncodeDeleteActorError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "invalid-scopes":
			var res actorservice.InvalidScopes
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "unauthorized":
			var res actorservice.Unauthorized
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalActorserviceviewsActorResultViewToActorResultResponse builds a value
// of type *ActorResultResponse from a value of type
// *actorserviceviews.ActorResultView.
func marshalActorserviceviewsActorResultViewToActorResultResponse(v *actorserviceviews.ActorResultView) *ActorResultResponse {
	res := &ActorResultResponse{
		ActorID:        *v.ActorID,
		ActorName:      v.ActorName,
		ActorSex:       v.ActorSex,
		ActorBirthdate: v.ActorBirthdate,
	}

	return res
}

// unmarshalActorInfoRequestBodyToActorserviceActorInfo builds a value of type
// *actorservice.ActorInfo from a value of type *ActorInfoRequestBody.
func unmarshalActorInfoRequestBodyToActorserviceActorInfo(v *ActorInfoRequestBody) *actorservice.ActorInfo {
	res := &actorservice.ActorInfo{
		ActorName:      *v.ActorName,
		ActorSex:       *v.ActorSex,
		ActorBirthdate: *v.ActorBirthdate,
	}

	return res
}