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

// GET /client-and-maybe-identity
func (cs *CustomServer) GetClientAndMaybeIdentity(ctx context.Context, request GetClientAndMaybeIdentityRequestObject) (GetClientAndMaybeIdentityResponseObject, error) {
	cs.Lock.Lock()
	defer cs.Lock.Unlock()

	return GetClientAndMaybeIdentity200JSONResponse{}, nil
}

// GET /client-with-id
func (cs *CustomServer) GetClientWithId(ctx context.Context, request GetClientWithIdRequestObject) (GetClientWithIdResponseObject, error) {
	cs.Lock.Lock()
	defer cs.Lock.Unlock()

	result := ClientWithId{12345, "Client_with_ID"}

	return GetClientWithId200JSONResponse(result), nil
}

// GET /identity-with-duplicate-field
func (cs *CustomServer) GetIdentityWithDuplicateField(ctx context.Context, request GetIdentityWithDuplicateFieldRequestObject) (GetIdentityWithDuplicateFieldResponseObject, error) {
	cs.Lock.Lock()
	defer cs.Lock.Unlock()

	result := IdentityWithDuplicateField{}

	return GetIdentityWithDuplicateField200JSONResponse(result), nil
}
