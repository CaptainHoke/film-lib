// Code generated by goa v3.15.2, DO NOT EDIT.
//
// FilmService HTTP server types
//
// Command:
// $ goa gen film-lib/design

package server

import (
	filmservice "film-lib/gen/film_service"
	filmserviceviews "film-lib/gen/film_service/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// GetAllFilmsRequestBody is the type of the "FilmService" service
// "getAllFilms" endpoint HTTP request body.
type GetAllFilmsRequestBody struct {
	SortBy *SortByRequestBody `form:"SortBy,omitempty" json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

// AddFilmRequestBody is the type of the "FilmService" service "addFilm"
// endpoint HTTP request body.
type AddFilmRequestBody struct {
	FilmInfo *FilmInfoRequestBody `form:"FilmInfo,omitempty" json:"FilmInfo,omitempty" xml:"FilmInfo,omitempty"`
}

// UpdateFilmInfoRequestBody is the type of the "FilmService" service
// "updateFilmInfo" endpoint HTTP request body.
type UpdateFilmInfoRequestBody struct {
	FilmInfo *FilmInfoRequestBody `form:"FilmInfo,omitempty" json:"FilmInfo,omitempty" xml:"FilmInfo,omitempty"`
}

// FilmResultResponseCollection is the type of the "FilmService" service
// "getAllFilms" endpoint HTTP response body.
type FilmResultResponseCollection []*FilmResultResponse

// AddFilmResponseBody is the type of the "FilmService" service "addFilm"
// endpoint HTTP response body.
type AddFilmResponseBody struct {
	// Unique ID of a Film
	FilmID      uint64  `form:"FilmID" json:"FilmID" xml:"FilmID"`
	Title       *string `form:"Title,omitempty" json:"Title,omitempty" xml:"Title,omitempty"`
	Description *string `form:"Description,omitempty" json:"Description,omitempty" xml:"Description,omitempty"`
	ReleaseDate *string `form:"ReleaseDate,omitempty" json:"ReleaseDate,omitempty" xml:"ReleaseDate,omitempty"`
	Rating      *string `form:"Rating,omitempty" json:"Rating,omitempty" xml:"Rating,omitempty"`
	Actors      *string `form:"Actors,omitempty" json:"Actors,omitempty" xml:"Actors,omitempty"`
}

// AddFilmAlreadyExistsResponseBody is the type of the "FilmService" service
// "addFilm" endpoint HTTP response body for the "already-exists" error.
type AddFilmAlreadyExistsResponseBody struct {
	// Error message
	Message string `form:"message" json:"message" xml:"message"`
	// ID of existing data
	ID string `form:"id" json:"id" xml:"id"`
}

// FilmResultResponse is used to define fields on response body types.
type FilmResultResponse struct {
	// Unique ID of a Film
	FilmID      uint64  `form:"FilmID" json:"FilmID" xml:"FilmID"`
	Title       *string `form:"Title,omitempty" json:"Title,omitempty" xml:"Title,omitempty"`
	Description *string `form:"Description,omitempty" json:"Description,omitempty" xml:"Description,omitempty"`
	ReleaseDate *string `form:"ReleaseDate,omitempty" json:"ReleaseDate,omitempty" xml:"ReleaseDate,omitempty"`
	Rating      *string `form:"Rating,omitempty" json:"Rating,omitempty" xml:"Rating,omitempty"`
	Actors      *string `form:"Actors,omitempty" json:"Actors,omitempty" xml:"Actors,omitempty"`
}

// SortByRequestBody is used to define fields on request body types.
type SortByRequestBody struct {
	// Field to sort by (Rating (default) | Title | Release Date)
	Field *string `form:"Field,omitempty" json:"Field,omitempty" xml:"Field,omitempty"`
	// Ascending / Descending
	Ordering *string `form:"Ordering,omitempty" json:"Ordering,omitempty" xml:"Ordering,omitempty"`
}

// FilmInfoRequestBody is used to define fields on request body types.
type FilmInfoRequestBody struct {
	// Title of a film
	Title *string `form:"Title,omitempty" json:"Title,omitempty" xml:"Title,omitempty"`
	// Description of a film
	Description *string `form:"Description,omitempty" json:"Description,omitempty" xml:"Description,omitempty"`
	// YYYY-MM-DD
	ReleaseDate *string `form:"ReleaseDate,omitempty" json:"ReleaseDate,omitempty" xml:"ReleaseDate,omitempty"`
	// Rating (0.0 - 10.0)
	Rating *float32 `form:"Rating,omitempty" json:"Rating,omitempty" xml:"Rating,omitempty"`
	// List of Actors involved in Film
	Actors []*ActorRequestBody `form:"Actors,omitempty" json:"Actors,omitempty" xml:"Actors,omitempty"`
}

// ActorRequestBody is used to define fields on request body types.
type ActorRequestBody struct {
	// Unique ID of an Actor
	ActorID *uint64 `form:"ActorID,omitempty" json:"ActorID,omitempty" xml:"ActorID,omitempty"`
	// Actor Info
	ActorInfo *ActorInfoRequestBody `form:"ActorInfo,omitempty" json:"ActorInfo,omitempty" xml:"ActorInfo,omitempty"`
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

// NewFilmResultResponseCollection builds the HTTP response body from the
// result of the "getAllFilms" endpoint of the "FilmService" service.
func NewFilmResultResponseCollection(res filmserviceviews.FilmResultCollectionView) FilmResultResponseCollection {
	body := make([]*FilmResultResponse, len(res))
	for i, val := range res {
		body[i] = marshalFilmserviceviewsFilmResultViewToFilmResultResponse(val)
	}
	return body
}

// NewAddFilmResponseBody builds the HTTP response body from the result of the
// "addFilm" endpoint of the "FilmService" service.
func NewAddFilmResponseBody(res *filmserviceviews.FilmResultView) *AddFilmResponseBody {
	body := &AddFilmResponseBody{
		FilmID:      *res.FilmID,
		Title:       res.Title,
		Description: res.Description,
		ReleaseDate: res.ReleaseDate,
		Rating:      res.Rating,
		Actors:      res.Actors,
	}
	return body
}

// NewAddFilmAlreadyExistsResponseBody builds the HTTP response body from the
// result of the "addFilm" endpoint of the "FilmService" service.
func NewAddFilmAlreadyExistsResponseBody(res *filmservice.AlreadyExists) *AddFilmAlreadyExistsResponseBody {
	body := &AddFilmAlreadyExistsResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewGetAllFilmsPayload builds a FilmService service getAllFilms endpoint
// payload.
func NewGetAllFilmsPayload(body *GetAllFilmsRequestBody, token string) *filmservice.GetAllFilmsPayload {
	v := &filmservice.GetAllFilmsPayload{}
	v.SortBy = unmarshalSortByRequestBodyToFilmserviceSortBy(body.SortBy)
	v.Token = token

	return v
}

// NewAddFilmPayload builds a FilmService service addFilm endpoint payload.
func NewAddFilmPayload(body *AddFilmRequestBody, token string) *filmservice.AddFilmPayload {
	v := &filmservice.AddFilmPayload{}
	v.FilmInfo = unmarshalFilmInfoRequestBodyToFilmserviceFilmInfo(body.FilmInfo)
	v.Token = token

	return v
}

// NewUpdateFilmInfoPayload builds a FilmService service updateFilmInfo
// endpoint payload.
func NewUpdateFilmInfoPayload(body *UpdateFilmInfoRequestBody, filmID uint64, token string) *filmservice.UpdateFilmInfoPayload {
	v := &filmservice.UpdateFilmInfoPayload{}
	v.FilmInfo = unmarshalFilmInfoRequestBodyToFilmserviceFilmInfo(body.FilmInfo)
	v.FilmID = filmID
	v.Token = token

	return v
}

// NewDeleteFilmPayload builds a FilmService service deleteFilm endpoint
// payload.
func NewDeleteFilmPayload(filmID string, token string) *filmservice.DeleteFilmPayload {
	v := &filmservice.DeleteFilmPayload{}
	v.FilmID = filmID
	v.Token = token

	return v
}

// ValidateGetAllFilmsRequestBody runs the validations defined on
// GetAllFilmsRequestBody
func ValidateGetAllFilmsRequestBody(body *GetAllFilmsRequestBody) (err error) {
	if body.SortBy == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("SortBy", "body"))
	}
	if body.SortBy != nil {
		if err2 := ValidateSortByRequestBody(body.SortBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateAddFilmRequestBody runs the validations defined on AddFilmRequestBody
func ValidateAddFilmRequestBody(body *AddFilmRequestBody) (err error) {
	if body.FilmInfo == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("FilmInfo", "body"))
	}
	if body.FilmInfo != nil {
		if err2 := ValidateFilmInfoRequestBody(body.FilmInfo); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdateFilmInfoRequestBody runs the validations defined on
// UpdateFilmInfoRequestBody
func ValidateUpdateFilmInfoRequestBody(body *UpdateFilmInfoRequestBody) (err error) {
	if body.FilmInfo == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("FilmInfo", "body"))
	}
	if body.FilmInfo != nil {
		if err2 := ValidateFilmInfoRequestBody(body.FilmInfo); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSortByRequestBody runs the validations defined on SortByRequestBody
func ValidateSortByRequestBody(body *SortByRequestBody) (err error) {
	if body.Field == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Field", "body"))
	}
	if body.Ordering == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Ordering", "body"))
	}
	return
}

// ValidateFilmInfoRequestBody runs the validations defined on
// FilmInfoRequestBody
func ValidateFilmInfoRequestBody(body *FilmInfoRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Title", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Description", "body"))
	}
	if body.ReleaseDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ReleaseDate", "body"))
	}
	if body.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Rating", "body"))
	}
	if body.Actors == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Actors", "body"))
	}
	if body.Title != nil {
		if utf8.RuneCountInString(*body.Title) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.Title", *body.Title, utf8.RuneCountInString(*body.Title), 1, true))
		}
	}
	if body.Title != nil {
		if utf8.RuneCountInString(*body.Title) > 150 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.Title", *body.Title, utf8.RuneCountInString(*body.Title), 150, false))
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 1000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.Description", *body.Description, utf8.RuneCountInString(*body.Description), 1000, false))
		}
	}
	if body.ReleaseDate != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.ReleaseDate", *body.ReleaseDate, "^\\d{4}\\-(0[1-9]|1[012])\\-(0[1-9]|[12][0-9]|3[01])$"))
	}
	if body.Rating != nil {
		if *body.Rating < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.Rating", *body.Rating, 0, true))
		}
	}
	if body.Rating != nil {
		if *body.Rating > 10 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.Rating", *body.Rating, 10, false))
		}
	}
	for _, e := range body.Actors {
		if e != nil {
			if err2 := ValidateActorRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateActorRequestBody runs the validations defined on ActorRequestBody
func ValidateActorRequestBody(body *ActorRequestBody) (err error) {
	if body.ActorID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ActorID", "body"))
	}
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