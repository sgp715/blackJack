package simulate

import "fmt"

type player struct {
	hand cards
	chips  int
	start int
	min int
	bet int
	wins int
	loses int
	ties int
}

func newPlayer(min, start int) player {
	return player{hand: make([]card, 2), min: min, chips: start, start: start}
}

type winnings struct {
	played string
	wins string
	loses string
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
	p.loses++
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


func (p *player) play(d dealer) {
}

//func (p *player) isBust() {
//	f := cardsKey[p.hand[0]]
//	s := cardsKey[p.hand[1]]
//}

func (p *player) results() winnings {
	t := p.wins + p.loses + p.ties
	total := fmt.Sprintf("total=%v", t)
	wins := fmt.Sprintf("wins=%v", p.wins)
	loses := fmt.Sprintf("loses=%v", p.loses)
	ties := fmt.Sprintf("ties=%v", p.ties)
	chips := fmt.Sprintf("chips=%v", p.chips)
	net := fmt.Sprintf("net=%v", (p.chips - p.start) * p.min)
	wp := fmt.Sprintf("wins percent=%.2f", (float64(p.wins) / float64(t)) * 100)
	lp := fmt.Sprintf("lose percent=%.2f", (float64(p.loses) / float64(t)) * 100)
	tp := fmt.Sprintf("ties percent=%.2f", (float64(p.ties) / float64(t)) * 100)
	return winnings{ played: total,  wins: wins, loses: loses, ties: ties, net: net, chips: chips, lp: lp, wp: wp, tp: tp}
}