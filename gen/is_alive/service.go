// Code generated by goa v3.15.2, DO NOT EDIT.
//
// IsAlive service
//
// Command:
// $ goa gen film-lib/design

package isalive

import (
	"context"
)

// API for checking if the server is in fact alive
type Service interface {
	// Check implements check.
	Check(context.Context) (err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "film-lib"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "1.0"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "IsAlive"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"check"}
