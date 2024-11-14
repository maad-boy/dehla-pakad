package deck

import (
	"github.com/maad-boy/dehla-pakad/internal/dto/card"
	"github.com/maad-boy/dehla-pakad/internal/enums"
	"math/rand"
	"time"
)

type Deck struct {
	cards []card.Card
}

func (d *Deck) Shuffle() *Deck {
	rand.Seed(time.Now().UnixNano())
	for i, _ := range d.cards {
		j := rand.Intn(len(d.cards))
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
	return d
}

func (d *Deck) Dealing(numberOfHand int) [][]card.Card {
	result := make([][]card.Card, numberOfHand)
	for i := 0; i < numberOfHand; i++ {
		result[i] = make([]card.Card, 0)
	}
	cardInEachHand := len(d.cards) / numberOfHand
	for i := 0; i < cardInEachHand; i++ {
		result[i] = d.cards[cardInEachHand*i : cardInEachHand*(i+1)]
	}
	return result
}

func NewDeck() *Deck {
	result := make([]card.Card, 0)
	for _, cardType := range enums.AllCardType() {
		for _, cardNumber := range enums.AllCardNumber() {
			_card := card.NewCard(cardType, cardNumber)
			result = append(result, _card)
		}
	}
	return &Deck{cards: result}
}
