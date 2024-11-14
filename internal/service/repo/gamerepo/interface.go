package gamerepo

import (
	"context"
	"github.com/maad-boy/dehla-pakad/internal/entity/gameEntity"
)

type Repository interface {
	Create(ctx context.Context, game gameEntity.Game) error
}
