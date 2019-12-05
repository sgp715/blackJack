package simulate

import (
	"math/rand"
	"time"
)

const shoeSize = 6

type shoe struct {
	cards cards
	selected int
	split int
	count int
}

func reseed() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
}

func newShoe() shoe {
	s := shoe{}
	s.shuffle()
	return s
}

func pop(a []card) (card, []card) {
	x, a := a[len(a)-1], a[:len(a)-1]
	return x, a
}

func findCount(c card) int {
	val := cardsKey[c][0]
	if val >= 2 || val <= 6 {
		return 1
	} else if val >= 7 || val <= 9 {
		return 0
	}
	return -1
}

func (s *shoe) next() card {
	if s.selected > s.split {
		s.shuffle()
	}
	var top card
	top, s.cards = pop(s.cards)
	c := findCount(top)
	s.count += c
	s.selected += 1
	return top
}

func (s *shoe) shuffle() {
	var cards []card
	for i := 0; i < shoeSize; i++ {
		cards = append(cards, newDeck()...)
	}
	s.cards = cards
	s.count = 0
	s.selected = 0
	reseed()
	halfShoe := int(shoeSize * 0.5) * deckSize
	s.split = rand.Intn(halfShoe) + halfShoe
	rand.Shuffle(shoeSize * deckSize, func(i, j int) { s.cards[i], s.cards[j] = s.cards[j], s.cards[i] })
}
