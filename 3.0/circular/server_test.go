package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ibiscum/go-oas-examples/3.0/circular/api"
	middleware "github.com/oapi-codegen/nethttp-middleware"
	"github.com/oapi-codegen/testutil"
	"github.com/stretchr/testify/require"
)

func doGet(t *testing.T, mux *mux.Router, url string) *httptest.ResponseRecorder {
	response := testutil.NewRequest().Get(url).WithAccept("text/plain; charset=utf-8").GoWithHTTPHandler(t, mux)
	return response.Recorder
}

func TestCustomServer(t *testing.T) {
	var err error

	// Get the swagger description of our API
	swagger, err := api.GetSwagger()
	require.NoError(t, err)

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// This is how you set up a basic router
	router := mux.NewRouter()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	router.Use(middleware.OapiRequestValidator(swagger))

	// Create an instance of our handler which satisfies the generated interface
	custom_server := api.NewCustomServer()
	custom_server_strict_handler := api.NewStrictHandler(custom_server, nil)

	// We now register our custom_server above as the handler for the interface
	api.HandlerFromMux(custom_server_strict_handler, router)

	t.Run("get-anything", func(t *testing.T) {
		rr := doGet(t, router, "/anything")
		_, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("body: %#v", err)
		}
		//require.Equal(t, true, strings.Contains(string(body), "parameter \"id-required\" in query has an error: value is required but missing\n"))
	})
}
