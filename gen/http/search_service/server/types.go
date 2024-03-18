// Code generated by goa v3.15.2, DO NOT EDIT.
//
// SearchService HTTP server types
//
// Command:
// $ goa gen film-lib/design

package server

import (
	searchservice "film-lib/gen/search_service"
	searchserviceviews "film-lib/gen/search_service/views"

	goa "goa.design/goa/v3/pkg"
)

// GetAllFilmsRequestBody is the type of the "SearchService" service
// "getAllFilms" endpoint HTTP request body.
type GetAllFilmsRequestBody struct {
	SortBy *SortByRequestBody `form:"SortBy,omitempty" json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

// SearchLibraryResponseBody is the type of the "SearchService" service
// "searchLibrary" endpoint HTTP response body.
type SearchLibraryResponseBody struct {
	// Unique ID of a Film
	FilmID uint64 `form:"FilmID" json:"FilmID" xml:"FilmID"`
	// Film Info
	FilmInfo *FilmInfoResponseBody `form:"FilmInfo" json:"FilmInfo" xml:"FilmInfo"`
}

// ActorResultResponseCollection is the type of the "SearchService" service
// "getAllActors" endpoint HTTP response body.
type ActorResultResponseCollection []*ActorResultResponse

// FilmResultResponseCollection is the type of the "SearchService" service
// "getAllFilms" endpoint HTTP response body.
type FilmResultResponseCollection []*FilmResultResponse

// FilmInfoResponseBody is used to define fields on response body types.
type FilmInfoResponseBody struct {
	// Title of a film
	Title string `form:"Title" json:"Title" xml:"Title"`
	// Description of a film
	Description string `form:"Description" json:"Description" xml:"Description"`
	// YYYY-MM-DD
	ReleaseDate string `form:"ReleaseDate" json:"ReleaseDate" xml:"ReleaseDate"`
	// Rating (0.0 - 10.0)
	Rating float32 `form:"Rating" json:"Rating" xml:"Rating"`
	// Actors' Ids
	Actors []uint64 `form:"Actors" json:"Actors" xml:"Actors"`
}

// ActorResultResponse is used to define fields on response body types.
type ActorResultResponse struct {
	// Unique ID of an Actor
	ActorID        uint64  `form:"ActorID" json:"ActorID" xml:"ActorID"`
	ActorName      *string `form:"ActorName,omitempty" json:"ActorName,omitempty" xml:"ActorName,omitempty"`
	ActorSex       *string `form:"ActorSex,omitempty" json:"ActorSex,omitempty" xml:"ActorSex,omitempty"`
	ActorBirthdate *string `form:"ActorBirthdate,omitempty" json:"ActorBirthdate,omitempty" xml:"ActorBirthdate,omitempty"`
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

// NewSearchLibraryResponseBody builds the HTTP response body from the result
// of the "searchLibrary" endpoint of the "SearchService" service.
func NewSearchLibraryResponseBody(res *searchservice.Film) *SearchLibraryResponseBody {
	body := &SearchLibraryResponseBody{
		FilmID: res.FilmID,
	}
	if res.FilmInfo != nil {
		body.FilmInfo = marshalSearchserviceFilmInfoToFilmInfoResponseBody(res.FilmInfo)
	}
	return body
}

// NewActorResultResponseCollection builds the HTTP response body from the
// result of the "getAllActors" endpoint of the "SearchService" service.
func NewActorResultResponseCollection(res searchserviceviews.ActorResultCollectionView) ActorResultResponseCollection {
	body := make([]*ActorResultResponse, len(res))
	for i, val := range res {
		body[i] = marshalSearchserviceviewsActorResultViewToActorResultResponse(val)
	}
	return body
}

// NewFilmResultResponseCollection builds the HTTP response body from the
// result of the "getAllFilms" endpoint of the "SearchService" service.
func NewFilmResultResponseCollection(res searchserviceviews.FilmResultCollectionView) FilmResultResponseCollection {
	body := make([]*FilmResultResponse, len(res))
	for i, val := range res {
		body[i] = marshalSearchserviceviewsFilmResultViewToFilmResultResponse(val)
	}
	return body
}

// NewSearchLibraryPayload builds a SearchService service searchLibrary
// endpoint payload.
func NewSearchLibraryPayload(queryString *string, token *string) *searchservice.SearchLibraryPayload {
	v := &searchservice.SearchLibraryPayload{}
	v.QueryString = queryString
	v.Token = token

	return v
}

// NewGetAllActorsPayload builds a SearchService service getAllActors endpoint
// payload.
func NewGetAllActorsPayload(token *string) *searchservice.GetAllActorsPayload {
	v := &searchservice.GetAllActorsPayload{}
	v.Token = token

	return v
}

// NewGetAllFilmsPayload builds a SearchService service getAllFilms endpoint
// payload.
func NewGetAllFilmsPayload(body *GetAllFilmsRequestBody, token *string) *searchservice.GetAllFilmsPayload {
	v := &searchservice.GetAllFilmsPayload{}
	v.SortBy = unmarshalSortByRequestBodyToSearchserviceSortBy(body.SortBy)
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
