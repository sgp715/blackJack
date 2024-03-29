package simulate

import "fmt"

type player struct {
	hand cards
	chips  int
	start int
	bet int
	wins int
	losses int
	ties int
}

func newPlayer(start int) player {
	return player{hand: make([]card, 2), chips: start, start: start}
}

type winnings struct {
	played string
	wins string
	losses string
	ties string
	lp string
	wp string
	tp string
	net string
	chips string
	ratio string
	expected string
}

func (p *player) done() bool {
	if p.chips <= 0 { return true }
	return false
}

func (p *player) win() {
	if p.done() {
		return
	}
	p.chips += (p.bet * 2)
	p.wins++
}

func (p *player) lose() {
	if p.done() {
		return
	}
	p.losses++
}

func (p *player) tie() {
	if p.done() {
		return
	}
	p.chips += (p.bet)
	p.ties++
}

func (p *player) calcBet(want int) int {
	if want > p.chips {
		return  p.chips
	}
	return want
}

func (p *player) initialBet(sh *shoe) {
	if p.chips <= 0 { return }
	//if 1 + (sh.count) < p.chips {
	///	p.bet = sh.count
	//} else {
	//	p.bet = p.chips
	//}
	p.placeBet(1)
}

func (p *player) placeBet(amount int) {
	p.chips -= amount
	p.bet += amount
}

func (p *player) reset() {
	p.hand = make([]card, 2)
	p.bet = 0
}


type move string
const (
	h  move = "H"
	st move = "S"
	db move = "D"
)

var hardTotals = map[card]map[int]move{
	two: { 8: h, 9: h, 10: db, 11: db, 12: h, 13: st, 14: st, 15: st, 16: st },
	three: { 8: h, 9: db, 10: db, 11: db, 12: h, 13: st, 14: st, 15: st, 16: st },
	four: { 8: h, 9: db, 10: db, 11: db, 12: st, 13: st, 14: st, 15: st, 16: st },
	five: { 8: h, 9: db, 10: db, 11: db, 12: st, 13: st, 14: st, 15: st, 16: st },
	six: { 8: h, 9: db, 10: db, 11: db, 12: st, 13: st, 14: st, 15: st, 16: st },
	seven: { 8: h, 9: h, 10: db, 11: db, 12: h, 13: h, 14: h, 15: h, 16: h },
	eight: { 8: h, 9: h, 10: db, 11: db, 12: h, 13: h, 14: h, 15: h, 16: h },
	nine: { 8: h, 9: h, 10: db, 11: db, 12: h, 13: h, 14: h, 15: h, 16: h },
	ten: { 8: h, 9: h, 10: h, 11: db, 12: h, 13: h, 14: h, 15: h, 16: h },
	j: { 8: h, 9: h, 10: h, 11: db, 12: h, 13: h, 14: h, 15: h, 16: h },
	q: { 8: h, 9: h, 10: h, 11: db, 12: h, 13: h, 14: h, 15: h, 16: h },
	k: { 8: h, 9: h, 10: h, 11: db, 12: h, 13: h, 14: h, 15: h, 16: h },
	a: { 8: h, 9: h, 10: h, 11: db, 12: h, 13: h, 14: h, 15: h, 16: h },
}

var softTotals = map[card]map[card]move{
	two: { two: h, three: h, four: h, five: h, six: h, seven: st},
	three: { two: h, three: h, four: h, five: h, six: db, seven: st},
	four: { two: h, three: h, four: db, five: db, six: db, seven: st},
	five: { two: db, three: db, four: db, five: db, six: db, seven: st},
	six: { two: db, three: db, four: db, five: db, six: db, seven: st},
	seven: { two: h, three: h, four: h, five: h, six: h, seven: st},
	eight: { two: h, three: h, four: h, five: h, six: h, seven: st},
	nine: { two: h, three: h, four: h, five: h, six: h, seven: h},
	ten: { two: h, three: h, four: h, five: h, six: h, seven: h},
	j: { two: h, three: h, four: h, five: h, six: h, seven: h},
	q: { two: h, three: h, four: h, five: h, six: h, seven: h},
	k: { two: h, three: h, four: h, five: h, six: h, seven: h},
	a: { two: h, three: h, four: h, five: h, six: h, seven: h},
}

func soft(hand cards) card {
	if hand[first] == a {
		if hand[second] != a {
			return hand[second]
		}
	}
	if hand[second] == a {
		if hand[first] != a {
			return hand[first]
		}
	}
	return ""
}

func (p *player) play(sh *shoe, d dealer) {
	if hardCard := soft(p.hand); hardCard != "" {
		if score(p.hand) > 18 {
			return
		}
		mv := softTotals[d.hand[upcard]][hardCard]
		if mv == st {
			return
		} else if mv == h {
			p.hand = append(p.hand, sh.next())
		}
	}
	sc := score(p.hand)
	if sc >= 17 {
		return
	}
	mv := hardTotals[d.hand[upcard]][sc]
	if mv == db {
		p.placeBet(p.calcBet(p.bet))
		p.hand = append(p.hand, sh.next())
		return
	}
	for mv == h {
		p.hand = append(p.hand, sh.next())
		sc = score(p.hand)
		if sc >= 17 {
			return
		}
		mv = hardTotals[d.hand[upcard]][sc]
	}
}

func (p *player) results() winnings {
	t := p.wins + p.losses + p.ties
	total := fmt.Sprintf("total=%v", t)
	wins := fmt.Sprintf("wins=%v", p.wins)
	losses := fmt.Sprintf("losses=%v", p.losses)
	ties := fmt.Sprintf("ties=%v", p.ties)
	chips := fmt.Sprintf("chips=%v", p.chips)
	net := fmt.Sprintf("net=%v", p.chips - p.start)
	wp := fmt.Sprintf("wins percent=%.2f", (float64(p.wins) / float64(t)) * 100)
	lp := fmt.Sprintf("loss percent=%.2f", (float64(p.losses) / float64(t)) * 100)
	tp := fmt.Sprintf("ties percent=%.2f", (float64(p.ties) / float64(t)) * 100)
	wvl := float64(p.wins) / float64(p.losses)
	ratio := fmt.Sprintf("ratio=%.2f", wvl)
	expected := fmt.Sprintf("expected=%.2f", (wvl * float64(p.start)) - float64(p.start))
	return winnings{ played: total,  wins: wins, losses: losses, ties: ties, net: net, chips: chips, lp: lp, wp: wp, tp: tp, ratio: ratio, expected: expected}
}
