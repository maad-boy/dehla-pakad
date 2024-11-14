package inmemory

import "github.com/maad-boy/dehla-pakad/internal/entity/gameEntity"

var data = make(map[string]gameEntity.Game)

type repo struct {
	data map[string]gameEntity.Game
}

func New() *repo {
	return &repo{
		data: data,
	}
}
