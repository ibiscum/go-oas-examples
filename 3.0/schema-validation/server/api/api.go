package api

import (
	"context"
)

type CustomServer struct{}

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ StrictServerInterface = (*CustomServer)(nil)

func NewCustomServer() *CustomServer {
	return &CustomServer{}
}

// GET /anything/booleans
func (cs *CustomServer) GetAnythingBooleans(ctx context.Context, request GetAnythingBooleansRequestObject) (GetAnythingBooleansResponseObject, error) {

	return nil, nil
}

// GET /anything/jsonschema-formats
func (cs *CustomServer) GetAnythingJsonschemaFormats(ctx context.Context, request GetAnythingJsonschemaFormatsRequestObject) (GetAnythingJsonschemaFormatsResponseObject, error) {

	return nil, nil
}

// GET /anything/numbers
func (cs *CustomServer) GetAnythingNumbers(ctx context.Context, request GetAnythingNumbersRequestObject) (GetAnythingNumbersResponseObject, error) {

	return nil, nil
}

// GET /anything/oas-formats
func (cs *CustomServer) GetAnythingOasFormats(ctx context.Context, request GetAnythingOasFormatsRequestObject) (GetAnythingOasFormatsResponseObject, error) {

	return nil, nil
}

// GET /anything/strings
func (cs *CustomServer) GetAnythingStrings(ctx context.Context, request GetAnythingStringsRequestObject) (GetAnythingStringsResponseObject, error) {

	return nil, nil
}
