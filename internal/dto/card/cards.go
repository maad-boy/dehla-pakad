package card

import (
	"math/rand"
	"time"
)

type Cards []Card

func newCards() Cards {
	return make(Cards, 0)
}

func (c Cards) Len() int {
	return len(c)
}

func (c Cards) swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Shuffle return vs inline change
func (c Cards) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i, _ := range c {
		j := rand.Intn(len(c))
		c.swap(i, j)
	}
}
