package simulate

import "fmt"

type player struct {
	hand cards
	chips  int
	start int
	min int
	bet int
	wins int
	losses int
	ties int
}

func newPlayer(min, start int) player {
	return player{hand: make([]card, 2), min: min, chips: start, start: start}
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

func (p *player) play(sh *shoe, d dealer) {
	sc := score(p.hand)
	if sc >= 17 {
		return
	}
	mv := hardTotals[d.hand[upcard]][sc]
	if mv == db {
		p.bet += p.calcBet(p.bet)
		topCard := sh.next()
		p.hand = append(p.hand, topCard)
		return
	}
	for mv != st {
		topCard := sh.next()
		p.hand = append(p.hand, topCard)
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
	net := fmt.Sprintf("net=%v", (p.chips - p.start) * p.min)
	wp := fmt.Sprintf("wins percent=%.2f", (float64(p.wins) / float64(t)) * 100)
	lp := fmt.Sprintf("loss percent=%.2f", (float64(p.losses) / float64(t)) * 100)
	tp := fmt.Sprintf("ties percent=%.2f", (float64(p.ties) / float64(t)) * 100)
	wvl := float64(p.wins) / float64(p.losses)
	ratio := fmt.Sprintf("ratio=%.2f", wvl)
	expected := fmt.Sprintf("expected=%.2f", (wvl * float64(p.start)) - float64(p.start))
	return winnings{ played: total,  wins: wins, losses: losses, ties: ties, net: net, chips: chips, lp: lp, wp: wp, tp: tp, ratio: ratio, expected: expected}
}
