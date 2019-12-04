package blackJack

const deckSize = 52

func newDeck() []card {
	var cards []card
	for i, c := range cardNames {
		for j := 0; j < 4; j++ {
			cards[i] = card(c)
		}
	}
	return cards
}
