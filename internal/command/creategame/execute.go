package creategame

import (
	"github.com/maad-boy/dehla-pakad/internal/service"
	"golang.org/x/net/context"
)

func Execute(ctx context.Context, payload Payload) (*Response, error) {
	game, err := service.GetGameService().Create(ctx)
	if err != nil {
		return nil, err
	}
	return &Response{
		GameId: game.ID,
	}, nil
}
