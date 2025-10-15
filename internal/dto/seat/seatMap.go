package seat

import (
	"fmt"
	"github.com/maad-boy/dehla-pakad/internal/enums"
)

type Seats map[enums.SeatNumber]Seat

func NewSeats() Seats {
	seatsNumber := []enums.SeatNumber{
		enums.Seat0,
		enums.Seat1,
		enums.Seat2,
	}
	result := make(map[enums.SeatNumber]Seat)
	for _, seatNumber := range seatsNumber {
		result[seatNumber] = NewSeat(seatNumber)
	}
	return result
}

func (s Seats) AssignUser(seatNumber enums.SeatNumber, userId string) error {
	if !seatNumber.IsValid() {
		return fmt.Errorf("invalid seat number: %d", seatNumber)
	}
	s[seatNumber] = s[seatNumber].WithUser(userId)
	return nil
}

func (s Seats) (seatNumber enums.SeatNumber) (Seat, bool)  {

}
