package simulate

type card string

const(
	two card = "2"
	three card = "3"
	four card = "4"
	five card = "5"
	six card = "6"
	seven card = "7"
	eight card = "8"
	nine card = "9"
	ten card = "10"
	j card = "J"
	q card = "Q"
	k card = "K"
	a card = "A"
)

var cardNames []card = []card{two, three, four, five, six, seven, eight, nine, ten, j, q, k, a}

var cardsKey map[card][]int = map[card][]int{
	two:[]int{2},
	three:[]int{3},
	four:[]int{4},
	five:[]int{5},
	six:[]int{6},
	seven:[]int{7},
	eight:[]int{8},
	nine:[]int{9},
	ten:[]int{10},
	j:[]int{10},
	q:[]int{10},
	k:[]int{10},
	a:[]int{1, 11}}

type cards []card

func score(hand []card) int {
	var aces int
	var score int
	for _, c := range hand {
		if c == a {
			aces++
			continue
		}
		score += cardsKey[c][0]
	}
	for i := 0; i < aces; i++ {
		if 11 + score > 21 {
			score += 1
		} else {
			score += 11
		}
	}
	return score
}