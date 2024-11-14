package getallgame

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
	games, err := service.GetGameService().GetAllGame(ctx)
	if err != nil {
		return nil, err
	}
	res := Response{
		GameIds: games,
	}
	return &res, nil
}
