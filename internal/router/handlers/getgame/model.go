package getgame

import "github.com/maad-boy/dehla-pakad/internal/entity/gameEntity"

type Request struct {
	GameId string `uri:"game_id"`
}

type Response struct {
	Game gameEntity.Game `json:"game"`
}
