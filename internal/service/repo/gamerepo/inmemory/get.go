package inmemory

import (
	"context"
	"github.com/maad-boy/dehla-pakad/internal/entity/gameEntity"
)

func (r repo) GetAllGames(ctx context.Context) ([]gameEntity.Game, error) {
	result := make([]gameEntity.Game, 0)
	for _, game := range r.data {
		result = append(result, game)
	}
	return result, nil
}
