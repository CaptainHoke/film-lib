package design

import . "goa.design/goa/v3/dsl"

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
