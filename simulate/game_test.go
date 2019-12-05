package simulate

import (
	"testing"
)

func TestPayout(t *testing.T) {
	tests := []struct{
		name string
		dHand cards
		pHand cards
		wins int
		losses int
		ties int
	} {
		{name: "blackjack", dHand: []card{k,a}, pHand: []card{k,a}, wins: 0, losses: 0, ties: 1},
		{name: "blackjack", dHand: []card{k,a}, pHand: []card{k,nine}, wins: 0, losses: 1, ties: 0},
		{name: "blackjack", dHand: []card{k,nine}, pHand: []card{k,a}, wins: 1, losses: 0, ties: 0},
		{name: "blackjack", dHand: []card{k,k,k}, pHand: []card{k,k,k}, wins: 0, losses: 0, ties: 1},
		{name: "blackjack", dHand: []card{k,nine}, pHand: []card{k,k,k}, wins: 0, losses: 1, ties: 0},
		{name: "blackjack", dHand: []card{k,k,k}, pHand: []card{k,nine}, wins: 1, losses: 0, ties: 0},
		{name: "blackjack", dHand: []card{a,a,a,k}, pHand: []card{k, a, nine}, wins: 1, losses: 0, ties: 0},
	}

	for _,tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			d := newDealer()
			d.hand = tc.dHand
			p := newPlayer(5, 10)
			p.hand = tc.pHand
			players := []*player{&p}
			payout(&d, players)
			r := p.results()
			if p.wins != tc.wins || p.losses != tc.losses || p.ties != tc.ties {
				t.Errorf("got %v want %v", r, tc)
			}
		})
	}
}
