package design

import . "goa.design/goa/v3/dsl"

var _ = Service("SearchService", func() {
	Description("API for querying the library")

	Error("unauthorized", String, "Credentials are invalid")
	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
	})

	Method("searchLibrary", func() {
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			Attribute("QueryString", String, "Actor or Film Name")
		})

		Security(JWTAuth, func() {
			Scope("api:read")
		})

		Result(Film)

		Error("invalid-scopes", String, "Token scopes are invalid")

		HTTP(func() {
			GET("/search")
			Param("QueryString")
			Response("invalid-scopes", StatusForbidden)
			Response(StatusOK)
		})
	})

	Method("getAllActors", func() {
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
		})

		Security(JWTAuth, func() {
			Scope("api:read")
		})

		Result(CollectionOf(ActorResult))

		Error("invalid-scopes", String, "Token scopes are invalid")

		HTTP(func() {
			GET("/actors")
			Response(StatusOK)
			Response("invalid-scopes", StatusForbidden)
		})
	})

	Method("getAllFilms", func() {
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			Attribute("SortBy", SortingOptions)
			Required("SortBy")
		})

		Security(JWTAuth, func() {
			Scope("api:read")
		})

		Result(CollectionOf(FilmResult))

		Error("invalid-scopes", String, "Token scopes are invalid")

		HTTP(func() {
			GET("/films")
			Response("invalid-scopes", StatusForbidden)
			Response(StatusOK)
		})
	})
})
