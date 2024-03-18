// Code generated by goa v3.15.2, DO NOT EDIT.
//
// SignIn HTTP client CLI support package
//
// Command:
// $ goa gen film-lib/design

package client

import (
	signin "film-lib/gen/sign_in"
)

// BuildAuthPayload builds the payload for the SignIn auth endpoint from CLI
// flags.
func BuildAuthPayload(signInAuthUsername string, signInAuthPassword string) (*signin.AuthPayload, error) {
	var username string
	{
		username = signInAuthUsername
	}
	var password string
	{
		password = signInAuthPassword
	}
	v := &signin.AuthPayload{}
	v.Username = username
	v.Password = password

	return v, nil
}
