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

// GET /anything
func (cs *CustomServer) GetAnything(ctx context.Context, request GetAnythingRequestObject) (GetAnythingResponseObject, error) {
	cs.Lock.Lock()
	defer cs.Lock.Unlock()
	return nil, nil
}
