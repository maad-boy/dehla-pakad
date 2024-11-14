package inmemory

import (
	"context"
	"github.com/maad-boy/dehla-pakad/internal/entity/gameEntity"
)

func (r repo) Create(ctx context.Context, game gameEntity.Game) error {
	r.data[game.ID] = game
	return nil
}
