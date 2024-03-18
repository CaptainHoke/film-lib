package design

import . "goa.design/goa/v3/dsl"

var _ = API("film-lib", func() {
	Title("Film Library")
	Description("BiS Service for querying and managing films and actors info")
	Server("film-lib", func() {
		Host("localhost", func() {
			URI("http://localhost:3239/api/v1")
		})
	})
})

var ActorResult = ResultType("application/vnd.actor", func() {
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
	Description("Actor + ID")

	Attribute("ActorID", UInt32, "Unique ID of an Actor", func() {
		Example(239)
	})
	Field(2, "ActorInfo", ActorInfo, "Actor Info")

	Required("ActorID", "ActorInfo")
})

var ActorInfo = Type("ActorInfo", func() {
	Description("Describes an Actor to be added")

	Attribute("ActorName", String, "Name of an Actor", func() {
		Example("Margo Robbie")
	})
	Attribute("ActorSex", String, "Sex of an Actor", func() {
		Example("F")
	})
	Attribute("ActorBirthdate", String, "YYYY-MM-DD", func() {
		Pattern(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`)
		Example("2024-03-18")
	})

	Required("ActorName", "ActorSex", "ActorBirthdate")
})
