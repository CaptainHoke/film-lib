// Code generated by goa v3.15.2, DO NOT EDIT.
//
// SearchService service
//
// Command:
// $ goa gen film-lib/design

package searchservice

import (
	"context"
	searchserviceviews "film-lib/gen/search_service/views"

	"goa.design/goa/v3/security"
)

// API for querying the library
type Service interface {
	// SearchLibrary implements searchLibrary.
	SearchLibrary(context.Context, *SearchLibraryPayload) (res *Film, err error)
	// GetAllActors implements getAllActors.
	GetAllActors(context.Context, *GetAllActorsPayload) (res ActorWithFilmsResultCollection, err error)
	// GetAllFilms implements getAllFilms.
	GetAllFilms(context.Context, *GetAllFilmsPayload) (res FilmResultCollection, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// APIName is the name of the API as defined in the design.
const APIName = "film-lib"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "1.0"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "SearchService"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"searchLibrary", "getAllActors", "getAllFilms"}

type ActorWithFilmsResult struct {
	// Unique ID of an Actor
	ActorID        uint64
	ActorName      *string
	ActorSex       *string
	ActorBirthdate *string
	FilmIDs        []uint64
}

// ActorWithFilmsResultCollection is the result type of the SearchService
// service getAllActors method.
type ActorWithFilmsResultCollection []*ActorWithFilmsResult

// Film is the result type of the SearchService service searchLibrary method.
type Film struct {
	// Unique ID of a Film
	FilmID uint64
	// Film Info
	FilmInfo *FilmInfo
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
	// Actors' Ids
	Actors []uint64
}

type FilmResult struct {
	// Unique ID of a Film
	FilmID      uint64
	Title       *string
	Description *string
	ReleaseDate *string
	Rating      *string
	Actors      *string
}

// FilmResultCollection is the result type of the SearchService service
// getAllFilms method.
type FilmResultCollection []*FilmResult

// GetAllActorsPayload is the payload type of the SearchService service
// getAllActors method.
type GetAllActorsPayload struct {
	// JWT used for authentication
	Token *string
}

// GetAllFilmsPayload is the payload type of the SearchService service
// getAllFilms method.
type GetAllFilmsPayload struct {
	// JWT used for authentication
	Token  *string
	SortBy *SortBy
}

// SearchLibraryPayload is the payload type of the SearchService service
// searchLibrary method.
type SearchLibraryPayload struct {
	// JWT used for authentication
	Token *string
	// Actor or Film Name
	QueryString *string
}

// Sorting configuration
type SortBy struct {
	// Field to sort by (Rating (default) | Title | Release Date)
	Field string
	// Ascending / Descending
	Ordering string
}

// Token scopes are invalid
type InvalidScopes string

// Credentials are invalid
type Unauthorized string

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

// NewActorWithFilmsResultCollection initializes result type
// ActorWithFilmsResultCollection from viewed result type
// ActorWithFilmsResultCollection.
func NewActorWithFilmsResultCollection(vres searchserviceviews.ActorWithFilmsResultCollection) ActorWithFilmsResultCollection {
	return newActorWithFilmsResultCollection(vres.Projected)
}

// NewViewedActorWithFilmsResultCollection initializes viewed result type
// ActorWithFilmsResultCollection from result type
// ActorWithFilmsResultCollection using the given view.
func NewViewedActorWithFilmsResultCollection(res ActorWithFilmsResultCollection, view string) searchserviceviews.ActorWithFilmsResultCollection {
	p := newActorWithFilmsResultCollectionView(res)
	return searchserviceviews.ActorWithFilmsResultCollection{Projected: p, View: "default"}
}

// NewFilmResultCollection initializes result type FilmResultCollection from
// viewed result type FilmResultCollection.
func NewFilmResultCollection(vres searchserviceviews.FilmResultCollection) FilmResultCollection {
	return newFilmResultCollection(vres.Projected)
}

// NewViewedFilmResultCollection initializes viewed result type
// FilmResultCollection from result type FilmResultCollection using the given
// view.
func NewViewedFilmResultCollection(res FilmResultCollection, view string) searchserviceviews.FilmResultCollection {
	p := newFilmResultCollectionView(res)
	return searchserviceviews.FilmResultCollection{Projected: p, View: "default"}
}

// newActorWithFilmsResultCollection converts projected type
// ActorWithFilmsResultCollection to service type
// ActorWithFilmsResultCollection.
func newActorWithFilmsResultCollection(vres searchserviceviews.ActorWithFilmsResultCollectionView) ActorWithFilmsResultCollection {
	res := make(ActorWithFilmsResultCollection, len(vres))
	for i, n := range vres {
		res[i] = newActorWithFilmsResult(n)
	}
	return res
}

// newActorWithFilmsResultCollectionView projects result type
// ActorWithFilmsResultCollection to projected type
// ActorWithFilmsResultCollectionView using the "default" view.
func newActorWithFilmsResultCollectionView(res ActorWithFilmsResultCollection) searchserviceviews.ActorWithFilmsResultCollectionView {
	vres := make(searchserviceviews.ActorWithFilmsResultCollectionView, len(res))
	for i, n := range res {
		vres[i] = newActorWithFilmsResultView(n)
	}
	return vres
}

// newActorWithFilmsResult converts projected type ActorWithFilmsResult to
// service type ActorWithFilmsResult.
func newActorWithFilmsResult(vres *searchserviceviews.ActorWithFilmsResultView) *ActorWithFilmsResult {
	res := &ActorWithFilmsResult{
		ActorName:      vres.ActorName,
		ActorSex:       vres.ActorSex,
		ActorBirthdate: vres.ActorBirthdate,
	}
	if vres.ActorID != nil {
		res.ActorID = *vres.ActorID
	}
	if vres.FilmIDs != nil {
		res.FilmIDs = make([]uint64, len(vres.FilmIDs))
		for i, val := range vres.FilmIDs {
			res.FilmIDs[i] = val
		}
	}
	return res
}

// newActorWithFilmsResultView projects result type ActorWithFilmsResult to
// projected type ActorWithFilmsResultView using the "default" view.
func newActorWithFilmsResultView(res *ActorWithFilmsResult) *searchserviceviews.ActorWithFilmsResultView {
	vres := &searchserviceviews.ActorWithFilmsResultView{
		ActorID:        &res.ActorID,
		ActorName:      res.ActorName,
		ActorSex:       res.ActorSex,
		ActorBirthdate: res.ActorBirthdate,
	}
	if res.FilmIDs != nil {
		vres.FilmIDs = make([]uint64, len(res.FilmIDs))
		for i, val := range res.FilmIDs {
			vres.FilmIDs[i] = val
		}
	}
	return vres
}

// newFilmResultCollection converts projected type FilmResultCollection to
// service type FilmResultCollection.
func newFilmResultCollection(vres searchserviceviews.FilmResultCollectionView) FilmResultCollection {
	res := make(FilmResultCollection, len(vres))
	for i, n := range vres {
		res[i] = newFilmResult(n)
	}
	return res
}

// newFilmResultCollectionView projects result type FilmResultCollection to
// projected type FilmResultCollectionView using the "default" view.
func newFilmResultCollectionView(res FilmResultCollection) searchserviceviews.FilmResultCollectionView {
	vres := make(searchserviceviews.FilmResultCollectionView, len(res))
	for i, n := range res {
		vres[i] = newFilmResultView(n)
	}
	return vres
}

// newFilmResult converts projected type FilmResult to service type FilmResult.
func newFilmResult(vres *searchserviceviews.FilmResultView) *FilmResult {
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
func newFilmResultView(res *FilmResult) *searchserviceviews.FilmResultView {
	vres := &searchserviceviews.FilmResultView{
		FilmID:      &res.FilmID,
		Title:       res.Title,
		Description: res.Description,
		ReleaseDate: res.ReleaseDate,
		Rating:      res.Rating,
		Actors:      res.Actors,
	}
	return vres
}
