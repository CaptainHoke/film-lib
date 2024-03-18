package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("film-lib", func() {
	Title("Film Library")
	Description("BiS Service for querying and managing films and actors info")
	Server("film-lib", func() {
		Host("localhost", func() {
			URI("http://localhost:3239/api/v1")
		})
	})
})

var _ = Service("ActorService", func() {
	Description("API for actor-related requests")

	Error("unauthorized", String, "Credentials are invalid")
	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
	})

	Method("getAllActors", func() {
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			Required("token")
		})

		Security(JWTAuth, func() {
			Scope("api:read")
		})

		Result(CollectionOf(ActorResult))

		Error("invalid-scopes", String, "Token scopes are invalid")

		HTTP(func() {
			GET("/api/v1/actors")
			Header("token:X-Authorization")
			Response("invalid-scopes", StatusForbidden)
			Response(StatusOK)
		})
	})

	Method("addActor", func() {
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			//Field(2, "ActorID", String, "Actor ID")
			Field(2, "ActorName", String, "Actor Name")
			Field(3, "ActorSex", String, "Actor Sex M/F")
			Field(4, "ActorBirthdate", String, "Actor Birthdate YYYY-MM-DD")

			Required("token", "ActorName", "ActorSex", "ActorBirthdate")
		})

		Security(JWTAuth, func() {
			Scope("api:write")
		})

		// TODO: Return ID only ?
		Result(ActorResult)

		Error("invalid-scopes", String, "Token scopes are invalid")
		Error("already-exists", AlreadyExists, "Actor already exists")

		HTTP(func() {
			POST("/api/v1/actors")
			Header("token:X-Authorization")
			Response("invalid-scopes", StatusForbidden)
			Response("already-exists", StatusBadRequest)
			Response(StatusCreated)
		})
	})

	Method("updateActorInfo", func() {
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			Field(2, "ActorID", String, "Actor ID")
			Field(3, "ActorName", String, "Actor Name")
			Field(4, "ActorSex", String, "Actor Sex M/F")
			Field(5, "ActorBirthdate", String, "Actor Birthdate YYYY-MM-DD")

			Required("token", "ActorID", "ActorName", "ActorSex", "ActorBirthdate")
		})

		Security(JWTAuth, func() {
			Scope("api:write")
		})

		Error("invalid-scopes", String, "Token scopes are invalid")
		Error("not-found", NotFound, "Actor not found")

		HTTP(func() {
			PUT("/api/v1/actors/{ActorID}")
			Header("token:X-Authorization")
			Response("invalid-scopes", StatusForbidden)
			Response(StatusCreated)
		})
	})

	Method("deleteActor", func() {
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			Field(2, "ActorID", String, "Actor ID")

			Required("token", "ActorID")
		})

		Security(JWTAuth, func() {
			Scope("api:write")
		})

		Error("invalid-scopes", String, "Token scopes are invalid")
		Error("not-found", NotFound, "Actor not found")

		HTTP(func() {
			DELETE("/api/v1/actors/{ActorID}")
			Header("token:X-Authorization")
			Response("invalid-scopes", StatusForbidden)
			Response(StatusNoContent)
		})
	})
})

var ActorResult = ResultType("application/vnd.actor", func() {
	Description("I don't really understand what's the difference between a regular Actor and this one is just yet")

	Reference(Actor)
	TypeName("ActorResult")

	Attributes(func() {
		Attribute("ActorID", String, "Unique ID of an Actor", func() {
			Example("239")
		})
		Field(2, "ActorName")
		Field(3, "ActorSex")
		Attribute("ActorBirthdate", String, "YYYY-MM-DD", func() {
			Example("2024-03-18")
		})
	})

	View("default", func() {
		Attribute("ActorID")
		Attribute("ActorName")
		Attribute("ActorSex")
		Attribute("ActorBirthdate")
	})

	Required("ActorID")
})

var Actor = Type("Actor", func() {
	Description("Should be self-explanatory duh")

	Attribute("ActorID", String, "Unique ID of an Actor", func() {
		Example("239")
	})
	Attribute("ActorName", String, "Name of an Actor", func() {
		Example("Margo Robbie")
	})
	Attribute("ActorSex", String, "Sex of an Actor", func() {
		Example("F")
	})
	Attribute("ActorBirthdate", String, "YYYY-MM-DD", func() {
		Example("2024-03-18")
	})
	Required("ActorID", "ActorName", "ActorSex", "ActorBirthdate")
})

// AlreadyExists is a custom type returned when trying to add an entity that is already present in the db
var AlreadyExists = Type("AlreadyExists", func() {
	Description("AlreadyExists is a custom type returned when trying to add an entity that is already present in the db")
	Attribute("message", String, "Error message", func() {
		Example("Entity already exists")
	})
	Field(2, "id", String, "ID of existing data")
	Required("message", "id")
})

// NotFound is a custom type where we add the queried field in the response
var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when " +
		"the requested data that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Example("Client ABCDEF12356890 not found")
	})
	Field(2, "id", String, "ID of missing data")
	Required("message", "id")
})

////////// Sign-In

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint.`)
	Scope("api:read", "Read-only access")
	Scope("api:write", "Read and write access")
})

// BasicAuth is DSL function for adding basic auth support in the API
var BasicAuth = BasicAuthSecurity("basic", func() {
	Description("Basic authentication used to authenticate security principal during sign-in")
	Scope("api:read", "Read-only access")
})

// Sign-In Service is the service used to authenticate users and assign JWT tokens for their sessions
var _ = Service("sign-in", func() {
	Description("The Sign-In service authenticates users and validates tokens")

	Error("unauthorized", String, "Credentials are invalid")
	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
	})

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
			POST("/sign-in/auth")
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
