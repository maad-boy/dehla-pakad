package game

import "context"

func (s Service) GetAllGame(ctx context.Context) ([]string, error) {
	games, err := s.repo.GetAllGames(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]string, len(games))
	for i, game := range games {
		result[i] = game.ID
	}
	return result, nil
}
