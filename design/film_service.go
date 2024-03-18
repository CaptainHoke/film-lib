package design

import . "goa.design/goa/v3/dsl"

var _ = Service("FilmService", func() {
	Description("API for film-related requests")

	Error("unauthorized", String, "Credentials are invalid")
	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
	})

	Method("addFilm", func() {
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			Attribute("FilmInfo", FilmInfo)

			Required("FilmInfo")
		})

		Security(JWTAuth, func() {
			Scope("api:write")
		})

		Result(FilmResult)

		Error("invalid-scopes", String, "Token scopes are invalid")
		Error("already-exists", AlreadyExists, "Film already exists")

		HTTP(func() {
			POST("/films")
			Response("invalid-scopes", StatusForbidden)
			Response("already-exists", StatusBadRequest)
			Response(StatusCreated)
		})
	})

	Method("updateFilmInfo", func() {
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			Attribute("FilmID", UInt64)
			Attribute("FilmInfo", FilmInfo)

			Required("FilmInfo")
		})

		Security(JWTAuth, func() {
			Scope("api:write")
		})

		Error("invalid-scopes", String, "Token scopes are invalid")
		Error("not-found", NotFound, "Film not found")

		HTTP(func() {
			PUT("/films/{FilmID}")
			Response("invalid-scopes", StatusForbidden)
			Response("not-found", StatusNoContent)
			Response(StatusCreated)
		})
	})

	Method("deleteFilm", func() {
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			Attribute("FilmID", UInt64, "Film ID")

			Required("FilmID")
		})

		Security(JWTAuth, func() {
			Scope("api:write")
		})

		Error("invalid-scopes", String, "Token scopes are invalid")

		HTTP(func() {
			DELETE("/films/{FilmID}")
			Response("invalid-scopes", StatusForbidden)
			Response(StatusNoContent)
		})
	})
})

var SortingOptions = Type("SortBy", func() {
	Description("Sorting configuration")
	Attribute("Field", String, "Field to sort by (Rating (default) | Title | Release Date)", func() {
		Default("Rating")
		Example("Rating")
	})
	Attribute("Ordering", String, "Ascending / Descending", func() {
		Default("Descending")
		Example("Descending")
	})
	Required("Field", "Ordering")
})

var FilmResult = ResultType("application/vnd.film", func() {
	Reference(Film)
	TypeName("FilmResult")

	Attributes(func() {
		Attribute("FilmID")
		Attribute("Title")
		Attribute("Description")
		Attribute("ReleaseDate")
		Attribute("Rating")
		Attribute("Actors")
	})

	View("default", func() {
		Attribute("FilmID")
		Attribute("Title")
		Attribute("Description")
		Attribute("ReleaseDate")
		Attribute("Rating")
		Attribute("Actors")
	})

	Required("FilmID")
})

var Film = Type("Film", func() {
	Description("Film + ID")

	Attribute("FilmID", UInt64, "Unique ID of a Film", func() {
		Example(239)
	})
	Attribute("FilmInfo", FilmInfo, "Film Info")

	Required("FilmID", "FilmInfo")
})

var FilmInfo = Type("FilmInfo", func() {
	Description("Describes a Film to be added")

	Attribute("Title", String, "Title of a film", func() {
		MinLength(1)
		MaxLength(150)
	})
	Attribute("Description", String, "Description of a film", func() {
		MaxLength(1000)
	})
	Attribute("ReleaseDate", String, "YYYY-MM-DD", func() {
		Pattern(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`)
		Example("2024-03-18")
	})
	Attribute("Rating", Float32, "Rating (0.0 - 10.0)", func() {
		Minimum(0.0)
		Maximum(10.0)
	})
	Attribute("Actors", ArrayOf(UInt64), "Actors' Ids")

	Required("Title", "Description", "ReleaseDate", "Rating", "Actors")
})
