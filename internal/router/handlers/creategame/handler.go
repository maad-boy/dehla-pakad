package creategame

import (
	"github.com/maad-boy/dehla-pakad/internal/command/creategame"
	"golang.org/x/net/context"
)

type handler struct{}

func NewHandler() handler {
	return handler{}
}

func (h handler) Validate(ctx context.Context, req Request) error {
	return nil
}

func (h handler) Handle(ctx context.Context, req Request) (*Response, error) {
	payload := creategame.Payload{}
	_res, err := creategame.Execute(ctx, payload)
	if err != nil {
		return nil, err
	}

	res := Response{
		GameId: _res.GameId,
	}
	return &res, nil
}
