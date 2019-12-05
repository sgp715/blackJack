package simulate

const deckSize = 52

func newDeck() []card {
	var cards []card
	for _, c := range cardNames {
		for j := 0; j < 4; j++ {
			cards = append(cards, c)
		}
	}
	return cards
}
