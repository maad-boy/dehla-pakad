package health

import (
	"golang.org/x/net/context"
)

type Req struct {
}

type Res struct {
	Message string `json:"message"`
}

type handler struct {
}

func NewHandler() handler {
	return handler{}
}

func (h handler) Validate(ctx context.Context, req int) error {
	return nil
}

func (h handler) Handle(ctx context.Context, res int) (*Res, error) {
	return &Res{"running"}, nil
}
