package service

import "github.com/maad-boy/dehla-pakad/internal/service/game"

func GetGameService() game.Service {
	return game.GetService()
}
