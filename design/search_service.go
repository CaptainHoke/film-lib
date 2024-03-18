package design

import . "goa.design/goa/v3/dsl"

// TODO: Kinda makes sense to move getAllActors and getAllFilms here
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
			Required("token", "QueryString")
		})

		Security(JWTAuth, func() {
			Scope("api:read")
		})

		Result(Film)

		Error("invalid-scopes", String, "Token scopes are invalid")

		HTTP(func() {
			GET("/search")
			Header("token:X-Authorization")
			Response("invalid-scopes", StatusForbidden)
			Response(StatusOK)
		})
	})
})
