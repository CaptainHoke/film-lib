package design

import . "goa.design/goa/v3/dsl"

var _ = API("film-lib", func() {
	Title("Film Library")
	Description("BiS Service for querying and managing films and actors info")
	Version("1.0")
	Server("film-lib", func() {
		Host("localhost", func() {
			URI("http://localhost:3239")
		})
	})
})
