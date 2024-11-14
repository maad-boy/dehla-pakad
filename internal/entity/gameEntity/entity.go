package gameEntity

import "github.com/maad-boy/dehla-pakad/internal/dto/card"

type PlayerInfo struct {
	ID   string
	Name string
}

type GameCard struct {
	PlayerCard [][]card.Card
}

type Game struct {
	ID      string
	Players []PlayerInfo
	Cards   GameCard
	TurnIdx int
}
