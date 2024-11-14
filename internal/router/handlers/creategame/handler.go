package creategame

import "golang.org/x/net/context"

type handler struct{}

func NewHandler() handler {
	return handler{}
}

func (h handler) Validate(ctx context.Context, req Request) error {
	return nil
}

func (h handler) Handle(ctx context.Context, req Request) (*Response, error) {
	res := Response{
		GameId: "123",
	}
	return &res, nil
}
