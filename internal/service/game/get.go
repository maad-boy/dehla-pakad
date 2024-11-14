package game

import (
	"context"
	"github.com/maad-boy/dehla-pakad/internal/entity/gameEntity"
)

func (s Service) GetAllGame(ctx context.Context) ([]string, error) {
	games, err := s.repo.GetAllGames(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]string, len(games))
	for i, game := range games {
		result[i] = game.ID
	}
	return result, nil
}

func (s Service) GetGameById(ctx context.Context, gameId string) (*gameEntity.Game, error) {
	return s.repo.GetGameById(ctx, gameId)
}
