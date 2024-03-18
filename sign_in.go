package filmlib

import (
	"context"
	signin "film-lib/gen/sign_in"
	"fmt"
	"log"

	"goa.design/goa/v3/security"
)

// SignIn service example implementation.
// The example methods log the requests and return zero values.
type signInsrvc struct {
	logger *log.Logger
}

// NewSignIn returns the SignIn service implementation.
func NewSignIn(logger *log.Logger) signin.Service {
	return &signInsrvc{logger}
}

// BasicAuth implements the authorization logic for service "SignIn" for the
// "basic" security scheme.
func (s *signInsrvc) BasicAuth(ctx context.Context, user, pass string, scheme *security.BasicScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// Creates a valid JWT
func (s *signInsrvc) Auth(ctx context.Context, p *signin.AuthPayload) (res *signin.Creds, err error) {
	res = &signin.Creds{}
	s.logger.Print("signIn.auth")
	return
}
