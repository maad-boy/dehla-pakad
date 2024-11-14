package service

import (
	"github.com/maad-boy/dehla-pakad/internal/service/game"
	"github.com/maad-boy/dehla-pakad/internal/service/repo/gamerepo"
)

func GetGameService() game.Service {
	return game.GetService(gamerepo.GetRepo())
}
