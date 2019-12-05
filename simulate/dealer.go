package simulate

type dealer struct {
	facedown card
	faceup card
}

func (d *dealer) play(players []*player) {
}

func (d *dealer) is21() bool {
	if d.faceup == a {
		if cardsKey[d.facedown][0] == 10 {
			return true
		}
	}
	return false
}

func newDealer() dealer {
	return dealer{}
}