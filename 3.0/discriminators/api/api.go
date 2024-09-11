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
func (cs *CustomServer) OneOfWithTopLevelDiscriminatorAndMapping(ctx context.Context, request OneOfWithTopLevelDiscriminatorAndMappingRequestObject) (OneOfWithTopLevelDiscriminatorAndMappingResponseObject, error) {
	cs.Lock.Lock()
	defer cs.Lock.Unlock()

	return OneOfWithTopLevelDiscriminatorAndMapping200Response{}, nil
}
