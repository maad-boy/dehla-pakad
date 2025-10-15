package seat

import (
	"github.com/maad-boy/dehla-pakad/internal/dto/card"
	"github.com/maad-boy/dehla-pakad/internal/enums"
)

type Seat struct {
	Number enums.SeatNumber
	UserId string
	Cards  []card.Card
}

func NewSeat(number enums.SeatNumber) Seat {
	return Seat{
		Number: number,
		Cards:  make([]card.Card, 0),
	}
}

func (s Seat) WithUser(userId string) Seat {
	s.UserId = userId
	return s
}
