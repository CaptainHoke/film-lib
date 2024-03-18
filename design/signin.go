package design

import . "goa.design/goa/v3/dsl"

// BasicAuth is DSL function for adding basic auth support in the API
var BasicAuth = BasicAuthSecurity("basic", func() {
	Description("Basic authentication used to authenticate security principal during sign-in")
	Scope("api:read", "Read-only access")
})

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint.`)
	Scope("api:read", "Read-only access")
	Scope("api:write", "Read and write access")
})

// Sign-In Service is the service used to authenticate users and assign JWT tokens for their sessions
var _ = Service("SignIn", func() {
	Description("The Sign-In service authenticates users and validates tokens")

	Method("auth", func() {
		Description("Creates a valid JWT")

		Security(BasicAuth)

		Payload(func() {
			Description("Credentials used to authenticate to retrieve JWT token")
			UsernameField(1, "username",
				String, "Username used to perform sign-in", func() {
					Example("user")
				})
			PasswordField(2, "password",
				String, "Password used to perform sign-in", func() {
					Example("password")
				})
			Required("username", "password")
		})

		Result(Creds)

		HTTP(func() {
			POST("/sign-in")
			Response(StatusOK)
		})
	})
})

// Creds is a custom type for replying Tokens
var Creds = Type("Creds", func() {
	Field(1, "jwt", String, "JWT token", func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
			"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9" +
			"lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHD" +
			"cEfxjoYZgeFONFh7HgQ")
	})

	Required("jwt")
})
