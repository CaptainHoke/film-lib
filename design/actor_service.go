package design

import . "goa.design/goa/v3/dsl"

var _ = Service("ActorService", func() {
	Description("API for actor-related requests")

	Method("addActor", func() {
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			Attribute("ActorInfo", ActorInfo)

			Required("ActorInfo")
		})

		Security(JWTAuth, func() {
			Scope("api:write")
		})

		Result(UInt64)

		Error("invalid-scopes", String, "Token scopes are invalid")
		Error("already-exists", AlreadyExists, "Actor already exists")

		HTTP(func() {
			POST("/actors")
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
			Attribute("ActorID", UInt64)
			Attribute("ActorInfo", ActorInfo)

			Required("ActorID", "ActorInfo")
		})

		Security(JWTAuth, func() {
			Scope("api:write")
		})

		Error("invalid-scopes", String, "Token scopes are invalid")
		Error("not-found", NotFound, "Actor not found")

		HTTP(func() {
			PUT("/actors/{ActorID}")
			Response("invalid-scopes", StatusForbidden)
			Response(StatusCreated)
		})
	})

	Method("deleteActor", func() {
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			Attribute("ActorID", UInt64, "Actor ID")

			Required("ActorID")
		})

		Security(JWTAuth, func() {
			Scope("api:write")
		})

		Error("invalid-scopes", String, "Token scopes are invalid")
		Error("not-found", NotFound, "Actor not found")

		HTTP(func() {
			DELETE("/actors/{ActorID}")
			Response("invalid-scopes", StatusForbidden)
			Response(StatusNoContent)
		})
	})
})

var ActorResult = ResultType("application/vnd.actor", func() {
	Reference(Actor)
	TypeName("ActorResult")

	Attributes(func() {
		Attribute("ActorID")
		Attribute("ActorName")
		Attribute("ActorSex")
		Attribute("ActorBirthdate")
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
	Description("Actor + ID")

	Attribute("ActorID", UInt64, "Unique ID of an Actor", func() {
		Example(239)
	})
	Attribute("ActorInfo", ActorInfo, "Actor Info")

	Required("ActorID", "ActorInfo")
})

var ActorInfo = Type("ActorInfo", func() {
	Description("Describes an Actor to be added")

	Attribute("ActorName", String, "Name of an Actor", func() {
		Pattern(`^.*\S.*$`)
		MaxLength(32)
		Example("Margo Robbie")
	})
	Attribute("ActorSex", String, "Sex of an Actor", func() {
		Pattern(`^(M|F)$`)
		Example("F")
	})
	Attribute("ActorBirthdate", String, "YYYY-MM-DD", func() {
		Pattern(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`)
		Example("2024-03-18")
	})

	Required("ActorName", "ActorSex", "ActorBirthdate")
})
