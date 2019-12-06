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
	p.bet = 0
}

func (p *player) lose() {
	if p.done() {
		return
	}
	p.losses++
	p.bet = 0
}

func (p *player) tie() {
	if p.done() {
		return
	}
	p.chips += (p.bet)
	p.ties++
	p.bet = 0
}

func (p *player) placeBet(sh *shoe) {
	if p.chips <= 0 { return }
	if 1 + (sh.count) < p.chips {
		p.bet = sh.count
	} else {
		p.bet = p.chips
	}
	p.chips -= p.bet
}

func (p *player) reset() {
	p.hand = make([]card, 2)
}


type move string
const (
	h move = "H"
	s move = "S"
)

var strategy map[card]map[int]move = map[card]map[int]move{
	two: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: s, 14: s, 15: s, 16: s, 17: s},
	three: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: s, 14: s, 15: s, 16: s, 17: s},
	four: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: s, 13: s, 14: s, 15: s, 16: s, 17: s},
	five: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: s, 13: s, 14: s, 15: s, 16: s, 17: s},
	six: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: s, 13: s, 14: s, 15: s, 16: s, 17: s},
	seven: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
	eight: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
	nine: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
	ten: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
	j: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
	q: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
	k: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
	a: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
}

func (p *player) play(sh *shoe, d dealer) {
	sc := score(p.hand)
	for sc <= 17 && strategy[d.hand[faceup]][sc] != s {
		topCard := sh.next()
		p.hand = append(p.hand, topCard)
		sc = score(p.hand)
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
	return winnings{ played: total,  wins: wins, losses: losses, ties: ties, net: net, chips: chips, lp: lp, wp: wp, tp: tp}
}
