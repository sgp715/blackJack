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
	p.chips += 2
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
	p.chips++
	p.ties++
	p.bet = 0
}

func (p *player) placeBet(s *shoe) {
	if p.chips <= 0 { return }
	p.chips--
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
	two: map[int]move{ 8: , 9:, 10:, 11:, 12:, 13:, 14:, 15:, 16:, 17: },
	three: map[int]move{ 8: , 9:, 10:, 11:, 12:, 13:, 14:, 15:, 16:, 17: },
	four: map[int]move{ 8: , 9:, 10:, 11:, 12:, 13:, 14:, 15:, 16:, 17: },
	five: map[int]move{ 8: , 9:, 10:, 11:, 12:, 13:, 14:, 15:, 16:, 17: },
	six: map[int]move{ 8: , 9:, 10:, 11:, 12:, 13:, 14:, 15:, 16:, 17: },
	a: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
	a: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
	a: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
	a: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
	a: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
	a: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
	a: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
	a: map[int]move{ 8: h, 9: h, 10: h, 11: h, 12: h, 13: h, 14: h, 15: h, 16: h, 17: s},
}

func (p *player) play(d dealer) {
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