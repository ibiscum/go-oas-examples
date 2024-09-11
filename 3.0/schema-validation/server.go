package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ibiscum/go-oas-examples/3.0/schema-validation/api"
	middleware "github.com/oapi-codegen/nethttp-middleware"
)

func main() {
	port := flag.String("port", "8000", "Port for test HTTP server")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatalf("Error loading swagger spec\n: %s", err)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	custom_server := api.NewCustomServer()
	custom_server_strict_handler := api.NewStrictHandler(custom_server, nil)

	// This is how you set up a basic chi router
	router := mux.NewRouter()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	router.Use(middleware.OapiRequestValidator(swagger))

	// We now register our custom_server above as the handler for the interface
	api.HandlerFromMux(custom_server_strict_handler, router)

	s := &http.Server{
		Handler: router,
		Addr:    net.JoinHostPort("localhost", *port),
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
