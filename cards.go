package blackJack

type card string

var cardNames []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

var cardsKey map[card][]int = map[card][]int{
	"2":[]int{2},
	"3":[]int{3},
	"4":[]int{4},
	"5":[]int{5},
	"6":[]int{6},
	"7":[]int{7},
	"8":[]int{8},
	"9":[]int{9},
	"10":[]int{10},
	"J":[]int{10},
	"Q":[]int{10},
	"K":[]int{10},
	"A":[]int{1, 11}}