package simulate

const faceup = 0

type dealer struct {
	hand cards
}

func (d *dealer) play(players []*player) {
}

func (d *dealer) is21() bool {
	if d.hand[faceup] == a {
		if cardsKey[d.hand[first]][0] == 10 {
			return true
		}
	}
	return false
}

func newDealer() dealer {
	return dealer{hand: make([]card, 2)}
}