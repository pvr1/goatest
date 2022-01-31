package design

import (
	. "goa.design/goa/v3/dsl"
)

// API describes the global properties of the API server.
var _ = API("hello", func() {
	Title("Hello Service")
	Description("HTTP service for saying hello")

	// Server describes a single process listening for client requests. The DSL
	// defines the set of services that the server hosts as well as hosts details.
	Server("helloserver", func() {
		Description("The hello service")

		// List the services hosted by this server.
		Services("hellosvc")

		// List the Hosts and their transport URLs.
		Host("hellodevhost", func() {
			Description("Development hosts.")
			// Transport specific URLs, supported schemes are:
			// 'http', 'https', 'grpc' and 'grpcs' with the respective default
			// ports: 80, 443, 8080, 8443.
			URI("http://localhost:9000/calc")
			URI("grpc://localhost:9080")
		})
	})
})

// Service describes a service
var _ = Service("hellosvc", func() {
	Description("The hello service")

	// Method describes a service method (endpoint)
	Method("greet", func() {
		// Payload describes the method payload.
		// Here the payload is an object that consists of two fields.
		Payload(func() {
			// Attribute describes an object field
			Attribute("name", String, "Name", func() {
				Meta("rpc:tag", "1")
			})
			Required("name")
		})

		// Result describes the method result.
		// Here the result is a simple integer value.
		Result(String)

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests.
			// The payload fields are encoded as path parameters.
			GET("/greet/{name}")
			// Responses use a "200 OK" HTTP status.
			// The result is encoded in the response body.
			Response(StatusOK)
		})

		// GRPC describes the gRPC transport mapping.
		GRPC(func() {
			// Responses use a "OK" gRPC code.
			// The result is encoded in the response message.
			Response(CodeOK)
		})
	})

	// Serve the file gen/http/openapi3.json for requests sent to
	// /openapi.json.
	Files("/openapi.json", "openapi3.json")
	Files("/openapi", "openapi3.yaml")
})
