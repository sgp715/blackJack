package simulate

import (
	"testing"
)

func TestScore(t *testing.T) {
	tests := []struct{
		name string
		hand cards
		want int
	} {
		{name: "blackjack", hand: []card{k,a}, want: 21},
		{name: "aces", hand: []card{a,a}, want: 12},
		{name: "kings", hand: []card{k,k}, want: 20},
		{name: "bunch o'cards", hand: []card{a,a,k,six}, want: 18},
		{name: "eighteen", hand: []card{ten, eight}, want: 18},
		{name: "bust", hand: []card{ten, eight, six}, want: 24},
	}
	for _,tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := score(tc.hand)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}
