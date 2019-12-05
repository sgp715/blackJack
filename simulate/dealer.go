package simulate

//import "fmt"

const faceup = 0

type dealer struct {
	hand cards
}

func (d *dealer) play(s *shoe, players []*player) {
	allBust := true
	for _, p := range players {
		pScore := score(p.hand)
		if pScore < 22 {
			allBust = false
			break
		}
	}
	if allBust { return }
	dScore := score(d.hand)
	//fmt.Printf("score before: %v\n", dScore)
	for dScore < 17 {
		topCard := s.next()
		d.hand = append(d.hand, topCard)
		//fmt.Printf("hit card %v score %v\n", topCard, dScore)
		dScore = score(d.hand)
	}
}

func (d *dealer) is21() bool {
	if d.hand[faceup] == a {
		if cardsKey[d.hand[second]][0] == 10 {
			return true
		}
	}
	return false
}

func (d *dealer) reset() {
	d.hand = make([]card, 2)
}


func newDealer() dealer {
	return dealer{hand: make([]card, 2)}
}