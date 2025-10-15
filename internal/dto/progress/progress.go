package progress

import "github.com/maad-boy/dehla-pakad/internal/enums"

type HandInfo struct {
	TotalWon int
	DehlaWon int
}

func newHandInfo() HandInfo {
	result := HandInfo{
		TotalWon: 0,
		DehlaWon: 0,
	}
	return result
}

type GameHandInfo map[enums.Team]HandInfo

func newGameHandInfo() GameHandInfo {
	result := map[enums.Team]HandInfo{
		enums.TeamZero: newHandInfo(),
		enums.TeamOne:  newHandInfo(),
	}
	return result
}

type Progress struct {
	GameID   string           `json:"game_id"`
	Status   enums.GameStatus `json:"status"`
	HandInfo GameHandInfo     `json:"hand_info"`
	Winner   enums.Team       `json:"winner"`
}

// NewProgress should this be a pointer
func NewProgress(gameId string) Progress {
	result := Progress{
		GameID:   gameId,
		Status:   enums.GameCreated,
		HandInfo: newGameHandInfo(),
		Winner:   enums.TeamNon,
	}
	return result
}

func (p Progress) WithStatus(status enums.GameStatus) Progress {
	p.Status = status
	return p
}
