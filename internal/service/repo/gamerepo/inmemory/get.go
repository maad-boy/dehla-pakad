package inmemory

import (
	"context"
	"errors"
	"github.com/maad-boy/dehla-pakad/internal/entity/gameEntity"
)

func (r repo) GetAllGames(ctx context.Context) ([]gameEntity.Game, error) {
	result := make([]gameEntity.Game, 0)
	for _, game := range r.data {
		result = append(result, game)
	}
	return result, nil
}

func (r repo) GetGameById(ctx context.Context, gameId string) (*gameEntity.Game, error) {
	game, ok := r.data[gameId]
	if !ok {
		return nil, errors.New("game not found")
	}
	return &game, nil
}
