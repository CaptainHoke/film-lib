package design

import . "goa.design/goa/v3/dsl"

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
