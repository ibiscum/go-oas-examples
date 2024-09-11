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
