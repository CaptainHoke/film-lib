package services

import (
	"context"
	signin "film-lib/gen/sign_in"
	"film-lib/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"

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

	// TODO: I have no time
	if user != "admin" || pass != "admin" {
		return ctx, signin.Unauthorized("invalid username and password combination")
	}

	return ctx, nil
}

// Auth Creates a valid JWT
func (s *signInsrvc) Auth(ctx context.Context, p *signin.AuthPayload) (res *signin.Creds, err error) {
	s.logger.Print("signIn.auth")

	scopes := []string{"api:read"}
	if *p.Username == "admin" && *p.Password == "admin" {
		scopes = append(scopes, "api:write")
	}

	// create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Duration(9) * time.Minute).Unix(),
		"scopes": scopes,
	})

	s.logger.Printf("user '%s' logged in", *p.Username)

	// note that if "SignedString" returns an error then it is returned as
	// an internal error to the client
	t, err := token.SignedString(utils.JWTKey)
	if err != nil {
		return nil, err
	}

	res = &signin.Creds{
		JWT: t,
	}

	return
}
