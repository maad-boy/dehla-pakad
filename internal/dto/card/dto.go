package card

import "github.com/maad-boy/dehla-pakad/internal/enums"

type Card struct {
	Type   enums.CardType
	Number enums.CardNumber
}

func NewCard(_type enums.CardType, number enums.CardNumber) Card {
	if !_type.IsValid() || !number.IsValid() {
		panic("invalid card args")
	}
	return Card{
		Type:   _type,
		Number: number,
	}
}
