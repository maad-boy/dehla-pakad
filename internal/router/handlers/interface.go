package handlers

import "golang.org/x/net/context"

type Handler[Req any, Res any] interface {
	Validate(ctx context.Context, req Req) error
	Handle(ctx context.Context, req Req) (*Res, error)
}
