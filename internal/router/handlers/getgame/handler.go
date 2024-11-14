package getgame

import (
	"github.com/maad-boy/dehla-pakad/internal/service"
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
	game, err := service.GetGameService().GetGameById(ctx, req.GameId)
	if err != nil {
		return nil, err
	}
	res := Response{
		Game: *game,
	}
	return &res, nil
}
