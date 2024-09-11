package main

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ibiscum/go-oas-examples/3.0/schema-validation/api"
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

	// This is how you set up a basic chi router
	router := mux.NewRouter()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	router.Use(middleware.OapiRequestValidator(swagger))

	// Create an instance of our handler which satisfies the generated interface
	custom_server := api.NewCustomServer()
	custom_server_strict_handler := api.NewStrictHandler(custom_server, nil)

	// We now register our custom_server above as the handler for the interface
	api.HandlerFromMux(custom_server_strict_handler, router)

	t.Run("get-anything-numbers-01", func(t *testing.T) {
		rr := doGet(t, router, "/anything/numbers")
		body, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("body: %#v", err)
		}
		require.Equal(t, true, strings.Contains(string(body), "parameter \"id-required\" in query has an error: value is required but missing\n"))
	})

	t.Run("get-anything-numbers-02", func(t *testing.T) {
		rr := doGet(t, router, "/anything/numbers?id-required=9")
		body, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("body: %#v", err)
		}
		require.Equal(t, true, strings.Contains(string(body), "parameter \"id-required\" in query has an error: number must be at least 10\n"))
	})

	// 	t.Run("List all pets", func(t *testing.T) {
	// 		store.Pets = map[int64]api.Pet{
	// 			1: {},
	// 			2: {},
	// 		}

	// 		// Now, list all pets, we should have two
	// 		rr := doGet(t, r, "/pets")
	// 		assert.Equal(t, http.StatusOK, rr.Code)

	// 		var petList []api.Pet
	// 		err = json.NewDecoder(rr.Body).Decode(&petList)
	// 		assert.NoError(t, err, "error getting response", err)
	// 		assert.Equal(t, 2, len(petList))
	// 	})

	// 	t.Run("Filter pets by tag", func(t *testing.T) {
	// 		tag := "TagOfFido"

	// 		store.Pets = map[int64]api.Pet{
	// 			1: {
	// 				Tag: &tag,
	// 			},
	// 			2: {},
	// 		}

	// 		// Filter pets by tag, we should have 1
	// 		rr := doGet(t, r, "/pets?tags=TagOfFido")
	// 		assert.Equal(t, http.StatusOK, rr.Code)

	// 		var petList []api.Pet
	// 		err = json.NewDecoder(rr.Body).Decode(&petList)
	// 		assert.NoError(t, err, "error getting response", err)
	// 		assert.Equal(t, 1, len(petList))
	// 	})

	// 	t.Run("Filter pets by tag", func(t *testing.T) {
	// 		store.Pets = map[int64]api.Pet{
	// 			1: {},
	// 			2: {},
	// 		}

	// 		// Filter pets by non existent tag, we should have 0
	// 		rr := doGet(t, r, "/pets?tags=NotExists")
	// 		assert.Equal(t, http.StatusOK, rr.Code)

	// 		var petList []api.Pet
	// 		err = json.NewDecoder(rr.Body).Decode(&petList)
	// 		assert.NoError(t, err, "error getting response", err)
	// 		assert.Equal(t, 0, len(petList))
	// 	})

	// 	t.Run("Delete pets", func(t *testing.T) {
	// 		store.Pets = map[int64]api.Pet{
	// 			1: {},
	// 			2: {},
	// 		}

	// 		// Let's delete non-existent pet
	// 		rr := testutil.NewRequest().Delete("/pets/7").GoWithHTTPHandler(t, r).Recorder
	// 		assert.Equal(t, http.StatusNotFound, rr.Code)

	// 		var petError api.Error
	// 		err = json.NewDecoder(rr.Body).Decode(&petError)
	// 		assert.NoError(t, err, "error unmarshaling PetError")
	// 		assert.Equal(t, int32(http.StatusNotFound), petError.Code)

	// 		// Now, delete both real pets
	// 		rr = testutil.NewRequest().Delete("/pets/1").GoWithHTTPHandler(t, r).Recorder
	// 		assert.Equal(t, http.StatusNoContent, rr.Code)

	// 		rr = testutil.NewRequest().Delete("/pets/2").GoWithHTTPHandler(t, r).Recorder
	// 		assert.Equal(t, http.StatusNoContent, rr.Code)

	//		// Should have no pets left.
	//		var petList []api.Pet
	//		rr = doGet(t, r, "/pets")
	//		assert.Equal(t, http.StatusOK, rr.Code)
	//		err = json.NewDecoder(rr.Body).Decode(&petList)
	//		assert.NoError(t, err, "error getting response", err)
	//		assert.Equal(t, 0, len(petList))
	//	})
}
