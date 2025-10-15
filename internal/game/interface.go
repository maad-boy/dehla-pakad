package game

import (
	"github.com/maad-boy/dehla-pakad/internal/dto/card"
	"github.com/maad-boy/dehla-pakad/internal/dto/progress"
	"github.com/maad-boy/dehla-pakad/internal/enums"
	"golang.org/x/net/context"
)

type Game interface {
	AssignUser(ctx context.Context, SeatNumber enums.SeatNumber, userId string) error
	Start(ctx context.Context) error
	GetProgress(ctx context.Context) progress.Progress
	PlayHand(ctx context.Context, SeatNumber enums.SeatNumber, card card.Card) error
	Cancel(ctx context.Context) bool
}
