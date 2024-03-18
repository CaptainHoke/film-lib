// Code generated by goa v3.15.2, DO NOT EDIT.
//
// FilmService service
//
// Command:
// $ goa gen film-lib/design

package filmservice

import (
	"context"
	filmserviceviews "film-lib/gen/film_service/views"

	"goa.design/goa/v3/security"
)

// API for film-related requests
type Service interface {
	// GetAllFilms implements getAllFilms.
	GetAllFilms(context.Context, *GetAllFilmsPayload) (res FilmResultCollection, err error)
	// AddFilm implements addFilm.
	AddFilm(context.Context, *AddFilmPayload) (res *FilmResult, err error)
	// UpdateFilmInfo implements updateFilmInfo.
	UpdateFilmInfo(context.Context, *UpdateFilmInfoPayload) (err error)
	// DeleteFilm implements deleteFilm.
	DeleteFilm(context.Context, *DeleteFilmPayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// APIName is the name of the API as defined in the design.
const APIName = "film-lib"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "FilmService"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"getAllFilms", "addFilm", "updateFilmInfo", "deleteFilm"}

// Actor + ID
type Actor struct {
	// Unique ID of an Actor
	ActorID uint64
	// Actor Info
	ActorInfo *ActorInfo
}

// Describes an Actor to be added
type ActorInfo struct {
	// Name of an Actor
	ActorName string
	// Sex of an Actor
	ActorSex string
	// YYYY-MM-DD
	ActorBirthdate string
}

// AddFilmPayload is the payload type of the FilmService service addFilm method.
type AddFilmPayload struct {
	// JWT used for authentication
	Token    string
	FilmInfo *FilmInfo
}

// AlreadyExists is a custom type returned when trying to add an entity that is
// already present in the db
type AlreadyExists struct {
	// Error message
	Message string
	// ID of existing data
	ID string
}

// DeleteFilmPayload is the payload type of the FilmService service deleteFilm
// method.
type DeleteFilmPayload struct {
	// JWT used for authentication
	Token string
	// Film ID
	FilmID string
}

// Describes a Film to be added
type FilmInfo struct {
	// Title of a film
	Title string
	// Description of a film
	Description string
	// YYYY-MM-DD
	ReleaseDate string
	// Rating (0.0 - 10.0)
	Rating float32
	// List of Actors involved in Film
	Actors []*Actor
}

// FilmResult is the result type of the FilmService service addFilm method.
type FilmResult struct {
	// Unique ID of a Film
	FilmID      uint64
	Title       *string
	Description *string
	ReleaseDate *string
	Rating      *string
	Actors      *string
}

// FilmResultCollection is the result type of the FilmService service
// getAllFilms method.
type FilmResultCollection []*FilmResult

// GetAllFilmsPayload is the payload type of the FilmService service
// getAllFilms method.
type GetAllFilmsPayload struct {
	// JWT used for authentication
	Token  string
	SortBy *SortBy
}

// NotFound is the type returned when the requested data that does not exist.
type NotFound struct {
	// Message of error
	Message string
	// ID of missing data
	ID string
}

// Sorting configuration
type SortBy struct {
	// Field to sort by (Rating (default) | Title | Release Date)
	Field string
	// Ascending / Descending
	Ordering string
}

// UpdateFilmInfoPayload is the payload type of the FilmService service
// updateFilmInfo method.
type UpdateFilmInfoPayload struct {
	// JWT used for authentication
	Token    string
	FilmID   uint64
	FilmInfo *FilmInfo
}

// Token scopes are invalid
type InvalidScopes string

// Credentials are invalid
type Unauthorized string

// Error returns an error description.
func (e *AlreadyExists) Error() string {
	return "AlreadyExists is a custom type returned when trying to add an entity that is already present in the db"
}

// ErrorName returns "AlreadyExists".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *AlreadyExists) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "AlreadyExists".
func (e *AlreadyExists) GoaErrorName() string {
	return "already-exists"
}

// Error returns an error description.
func (e *NotFound) Error() string {
	return "NotFound is the type returned when the requested data that does not exist."
}

// ErrorName returns "NotFound".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *NotFound) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "NotFound".
func (e *NotFound) GoaErrorName() string {
	return "not-found"
}

// Error returns an error description.
func (e InvalidScopes) Error() string {
	return "Token scopes are invalid"
}

// ErrorName returns "invalid-scopes".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e InvalidScopes) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "invalid-scopes".
func (e InvalidScopes) GoaErrorName() string {
	return "invalid-scopes"
}

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "Credentials are invalid"
}

// ErrorName returns "unauthorized".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e Unauthorized) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "unauthorized".
func (e Unauthorized) GoaErrorName() string {
	return "unauthorized"
}

// NewFilmResultCollection initializes result type FilmResultCollection from
// viewed result type FilmResultCollection.
func NewFilmResultCollection(vres filmserviceviews.FilmResultCollection) FilmResultCollection {
	return newFilmResultCollection(vres.Projected)
}

// NewViewedFilmResultCollection initializes viewed result type
// FilmResultCollection from result type FilmResultCollection using the given
// view.
func NewViewedFilmResultCollection(res FilmResultCollection, view string) filmserviceviews.FilmResultCollection {
	p := newFilmResultCollectionView(res)
	return filmserviceviews.FilmResultCollection{Projected: p, View: "default"}
}

// NewFilmResult initializes result type FilmResult from viewed result type
// FilmResult.
func NewFilmResult(vres *filmserviceviews.FilmResult) *FilmResult {
	return newFilmResult(vres.Projected)
}

// NewViewedFilmResult initializes viewed result type FilmResult from result
// type FilmResult using the given view.
func NewViewedFilmResult(res *FilmResult, view string) *filmserviceviews.FilmResult {
	p := newFilmResultView(res)
	return &filmserviceviews.FilmResult{Projected: p, View: "default"}
}

// newFilmResultCollection converts projected type FilmResultCollection to
// service type FilmResultCollection.
func newFilmResultCollection(vres filmserviceviews.FilmResultCollectionView) FilmResultCollection {
	res := make(FilmResultCollection, len(vres))
	for i, n := range vres {
		res[i] = newFilmResult(n)
	}
	return res
}

// newFilmResultCollectionView projects result type FilmResultCollection to
// projected type FilmResultCollectionView using the "default" view.
func newFilmResultCollectionView(res FilmResultCollection) filmserviceviews.FilmResultCollectionView {
	vres := make(filmserviceviews.FilmResultCollectionView, len(res))
	for i, n := range res {
		vres[i] = newFilmResultView(n)
	}
	return vres
}

// newFilmResult converts projected type FilmResult to service type FilmResult.
func newFilmResult(vres *filmserviceviews.FilmResultView) *FilmResult {
	res := &FilmResult{
		Title:       vres.Title,
		Description: vres.Description,
		ReleaseDate: vres.ReleaseDate,
		Rating:      vres.Rating,
		Actors:      vres.Actors,
	}
	if vres.FilmID != nil {
		res.FilmID = *vres.FilmID
	}
	return res
}

// newFilmResultView projects result type FilmResult to projected type
// FilmResultView using the "default" view.
func newFilmResultView(res *FilmResult) *filmserviceviews.FilmResultView {
	vres := &filmserviceviews.FilmResultView{
		FilmID:      &res.FilmID,
		Title:       res.Title,
		Description: res.Description,
		ReleaseDate: res.ReleaseDate,
		Rating:      res.Rating,
		Actors:      res.Actors,
	}
	return vres
}
