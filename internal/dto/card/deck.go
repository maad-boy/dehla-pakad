package card

import (
	"github.com/maad-boy/dehla-pakad/internal/enums"
)

type Deck struct {
	cards Cards
}

func (d *Deck) Shuffle() {
	d.cards.Shuffle()
}

func (d *Deck) Dealing(numberOfHand int) []Cards {
	result := make([]Cards, numberOfHand)
	cardInEachHand := len(d.cards) / numberOfHand
	for i := 0; i < numberOfHand; i++ {
		result[i] = d.cards[cardInEachHand*i : cardInEachHand*(i+1)]
	}
	return result
}

func NewDeck() Deck {
	result := newCards()
	for _, cardType := range enums.AllCardType() {
		for _, cardNumber := range enums.AllCardNumber() {
			_card := NewCard(cardType, cardNumber)
			result = append(result, _card)
		}
	}
	return Deck{cards: result}
}
