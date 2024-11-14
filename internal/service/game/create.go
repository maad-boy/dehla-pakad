package game

import (
	"fmt"
	"github.com/maad-boy/dehla-pakad/internal/dto/deck"
	"github.com/maad-boy/dehla-pakad/internal/entity/gameEntity"
	"github.com/maad-boy/dehla-pakad/internal/utils"
	"golang.org/x/net/context"
)

func (s Service) Create(ctx context.Context) (*gameEntity.Game, error) {
	players := make([]gameEntity.PlayerInfo, 4)
	for i := 0; i < 4; i++ {
		players[i] = gameEntity.PlayerInfo{
			ID:   fmt.Sprint(i),
			Name: fmt.Sprint(i),
		}
	}

	game := gameEntity.Game{
		ID:      utils.GetStrId(10),
		Players: players,
		Cards: gameEntity.GameCard{
			PlayerCard: deck.NewDeck().Shuffle().Dealing(4),
		},
		TurnIdx: 0,
	}
	err := s.repo.Create(ctx, game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}
