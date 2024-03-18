package design

import . "goa.design/goa/v3/dsl"

// Docker shenanigans drive me insane
var _ = Service("IsAlive", func() {
	Description("API for checking if the server is in fact alive")

	Method("check", func() {
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})
})
