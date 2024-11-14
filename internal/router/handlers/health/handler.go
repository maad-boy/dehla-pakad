package health

import (
	"golang.org/x/net/context"
)

type handler struct {
}

func NewHandler() handler {
	return handler{}
}

func (h handler) Validate(ctx context.Context, req int) error {
	return nil
}

func (h handler) Handle(ctx context.Context, req int) (*int, error) {
	return nil, nil
}
