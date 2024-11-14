package game

import "github.com/maad-boy/dehla-pakad/internal/service/repo/gamerepo"

type Service struct {
	repo gamerepo.Repository
}

func newService() Service {
	return Service{}
}
