package gamerepo

import "github.com/maad-boy/dehla-pakad/internal/service/repo/gamerepo/inmemory"

func GetRepo() Repository {
	return inmemory.New()
}
