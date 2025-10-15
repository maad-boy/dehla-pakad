package enums

import "slices"

type GameStatus string

const (
	GameCreated  = "CREATED"
	GameStarted  = "STARTED"
	GameFinished = "FINISHED"
)

type Team string

const (
	TeamNon  Team = "NON"
	TeamZero Team = "0"
	TeamOne  Team = "1"
)

type SeatNumber string

const (
	Seat0 SeatNumber = "0"
	Seat1 SeatNumber = "1"
	Seat2 SeatNumber = "2"
	Seat3 SeatNumber = "3"
)

func AllSeatNumbers() []SeatNumber {
	result := []SeatNumber{
		Seat0,
		Seat1,
		Seat2,
		Seat3,
	}
	return result
}

func (sn SeatNumber) IsValid() bool {
	return slices.Contains(AllSeatNumbers(), sn)
}
