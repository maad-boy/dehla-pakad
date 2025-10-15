package game

import (
	"github.com/maad-boy/dehla-pakad/internal/dto/card"
	"github.com/maad-boy/dehla-pakad/internal/dto/progress"
	"github.com/maad-boy/dehla-pakad/internal/dto/seat"
	"github.com/maad-boy/dehla-pakad/internal/enums"
	"golang.org/x/net/context"
)

type IGame struct {
	GameId   string
	progress progress.Progress
	seats    seat.Seats
	deck     card.Deck
}

func (g *IGame) AssignUser(ctx context.Context, SeatNumber enums.SeatNumber, userId string) error {
	return g.seats.AssignUser(SeatNumber, userId)
}
func (g *IGame) Start(ctx context.Context) error {
	g.progress = g.progress.WithStatus(enums.GameStarted)
	g.deck.Shuffle()
	cards := g.deck.Dealing(4)
	g.seats.
	return nil
}
func (g *IGame) GetProgress(ctx context.Context) progress.Progress {
	return g.progress
}
func (g *IGame) PlayHand(ctx context.Context, SeatNumber enums.SeatNumber, card card.Card) error {

}

func (g *IGame) Cancel(ctx context.Context) bool {

}
