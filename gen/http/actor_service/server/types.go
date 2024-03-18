// Code generated by goa v3.15.2, DO NOT EDIT.
//
// ActorService HTTP server types
//
// Command:
// $ goa gen film-lib/design

package server

import (
	actorservice "film-lib/gen/actor_service"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// AddActorRequestBody is the type of the "ActorService" service "addActor"
// endpoint HTTP request body.
type AddActorRequestBody struct {
	ActorInfo *ActorInfoRequestBody `form:"ActorInfo,omitempty" json:"ActorInfo,omitempty" xml:"ActorInfo,omitempty"`
}

// UpdateActorInfoRequestBody is the type of the "ActorService" service
// "updateActorInfo" endpoint HTTP request body.
type UpdateActorInfoRequestBody struct {
	ActorInfo *ActorInfoRequestBody `form:"ActorInfo,omitempty" json:"ActorInfo,omitempty" xml:"ActorInfo,omitempty"`
}

// AddActorAlreadyExistsResponseBody is the type of the "ActorService" service
// "addActor" endpoint HTTP response body for the "already-exists" error.
type AddActorAlreadyExistsResponseBody struct {
	// Error message
	Message string `form:"message" json:"message" xml:"message"`
	// ID of existing data
	ID string `form:"id" json:"id" xml:"id"`
}

// ActorInfoRequestBody is used to define fields on request body types.
type ActorInfoRequestBody struct {
	// Name of an Actor
	ActorName *string `form:"ActorName,omitempty" json:"ActorName,omitempty" xml:"ActorName,omitempty"`
	// Sex of an Actor
	ActorSex *string `form:"ActorSex,omitempty" json:"ActorSex,omitempty" xml:"ActorSex,omitempty"`
	// YYYY-MM-DD
	ActorBirthdate *string `form:"ActorBirthdate,omitempty" json:"ActorBirthdate,omitempty" xml:"ActorBirthdate,omitempty"`
}

// NewAddActorAlreadyExistsResponseBody builds the HTTP response body from the
// result of the "addActor" endpoint of the "ActorService" service.
func NewAddActorAlreadyExistsResponseBody(res *actorservice.AlreadyExists) *AddActorAlreadyExistsResponseBody {
	body := &AddActorAlreadyExistsResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewAddActorPayload builds a ActorService service addActor endpoint payload.
func NewAddActorPayload(body *AddActorRequestBody, token *string) *actorservice.AddActorPayload {
	v := &actorservice.AddActorPayload{}
	v.ActorInfo = unmarshalActorInfoRequestBodyToActorserviceActorInfo(body.ActorInfo)
	v.Token = token

	return v
}

// NewUpdateActorInfoPayload builds a ActorService service updateActorInfo
// endpoint payload.
func NewUpdateActorInfoPayload(body *UpdateActorInfoRequestBody, actorID uint64, token *string) *actorservice.UpdateActorInfoPayload {
	v := &actorservice.UpdateActorInfoPayload{}
	v.ActorInfo = unmarshalActorInfoRequestBodyToActorserviceActorInfo(body.ActorInfo)
	v.ActorID = actorID
	v.Token = token

	return v
}

// NewDeleteActorPayload builds a ActorService service deleteActor endpoint
// payload.
func NewDeleteActorPayload(actorID uint64, token *string) *actorservice.DeleteActorPayload {
	v := &actorservice.DeleteActorPayload{}
	v.ActorID = actorID
	v.Token = token

	return v
}

// ValidateAddActorRequestBody runs the validations defined on
// AddActorRequestBody
func ValidateAddActorRequestBody(body *AddActorRequestBody) (err error) {
	if body.ActorInfo == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ActorInfo", "body"))
	}
	if body.ActorInfo != nil {
		if err2 := ValidateActorInfoRequestBody(body.ActorInfo); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdateActorInfoRequestBody runs the validations defined on
// UpdateActorInfoRequestBody
func ValidateUpdateActorInfoRequestBody(body *UpdateActorInfoRequestBody) (err error) {
	if body.ActorInfo == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ActorInfo", "body"))
	}
	if body.ActorInfo != nil {
		if err2 := ValidateActorInfoRequestBody(body.ActorInfo); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateActorInfoRequestBody runs the validations defined on
// ActorInfoRequestBody
func ValidateActorInfoRequestBody(body *ActorInfoRequestBody) (err error) {
	if body.ActorName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ActorName", "body"))
	}
	if body.ActorSex == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ActorSex", "body"))
	}
	if body.ActorBirthdate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ActorBirthdate", "body"))
	}
	if body.ActorName != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.ActorName", *body.ActorName, "^.*\\S.*$"))
	}
	if body.ActorName != nil {
		if utf8.RuneCountInString(*body.ActorName) > 32 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.ActorName", *body.ActorName, utf8.RuneCountInString(*body.ActorName), 32, false))
		}
	}
	if body.ActorSex != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.ActorSex", *body.ActorSex, "^(M|F)$"))
	}
	if body.ActorBirthdate != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.ActorBirthdate", *body.ActorBirthdate, "^\\d{4}\\-(0[1-9]|1[012])\\-(0[1-9]|[12][0-9]|3[01])$"))
	}
	return
}
