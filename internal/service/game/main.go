package game

import "github.com/maad-boy/dehla-pakad/internal/service/repo/gamerepo"

func GetService(repo gamerepo.Repository) Service {
	return Service{
		repo: repo,
	}
}
