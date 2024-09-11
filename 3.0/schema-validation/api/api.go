package api

import (
	"context"
	"sync"
)

type CustomServer struct {
	Lock sync.Mutex
}

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ StrictServerInterface = (*CustomServer)(nil)

func NewCustomServer() *CustomServer {
	return &CustomServer{}
}

// GET /anything/booleans
func (cs *CustomServer) GetAnythingBooleans(ctx context.Context, request GetAnythingBooleansRequestObject) (GetAnythingBooleansResponseObject, error) {
	cs.Lock.Lock()
	defer cs.Lock.Unlock()
	return nil, nil
}

// GET /anything/jsonschema-formats
func (cs *CustomServer) GetAnythingJsonschemaFormats(ctx context.Context, request GetAnythingJsonschemaFormatsRequestObject) (GetAnythingJsonschemaFormatsResponseObject, error) {
	cs.Lock.Lock()
	defer cs.Lock.Unlock()
	return nil, nil
}

// GET /anything/numbers
func (cs *CustomServer) GetAnythingNumbers(ctx context.Context, request GetAnythingNumbersRequestObject) (GetAnythingNumbersResponseObject, error) {
	cs.Lock.Lock()
	defer cs.Lock.Unlock()
	return GetAnythingNumbers200Response{}, nil
}

// GET /anything/oas-formats
func (cs *CustomServer) GetAnythingOasFormats(ctx context.Context, request GetAnythingOasFormatsRequestObject) (GetAnythingOasFormatsResponseObject, error) {
	cs.Lock.Lock()
	defer cs.Lock.Unlock()
	return nil, nil
}

// GET /anything/strings
func (cs *CustomServer) GetAnythingStrings(ctx context.Context, request GetAnythingStringsRequestObject) (GetAnythingStringsResponseObject, error) {
	cs.Lock.Lock()
	defer cs.Lock.Unlock()
	return GetAnythingStrings200Response{}, nil
}
