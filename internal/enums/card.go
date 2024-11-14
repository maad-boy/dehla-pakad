package enums

import "slices"

type CardType string

func (ct CardType) IsValid() bool {
	return slices.Contains(AllCardType(), ct)
}

const (
	Hearts   CardType = "HEARTS"
	Diamonds CardType = "DIAMONDS"
	Clubs    CardType = "CLUBS"
	Spades   CardType = "SPADES"
)

func AllCardType() []CardType {
	return []CardType{Hearts, Diamonds, Clubs, Spades}
}

type CardNumber int

func (cn CardNumber) IsValid() bool {
	return slices.Contains(AllCardNumber(), cn)
}

func AllCardNumber() []CardNumber {
	result := make([]CardNumber, 13)
	for i := range result {
		result[i] = CardNumber(i)
	}
	return result
}
