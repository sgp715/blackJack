package main

import (
	"./simulate"
	"flag"
	"fmt"
)

var (
	//decks = flag.Int("decks", 6, "number of decks")
	players = flag.Int("players", 1, "number of players at table")
	rounds = flag.Int("rounds", 100, "rounds to simulate")
	min = flag.Int("min", 10, "min bet")
	multi = flag.Int("multi", 40, "multiplier is the min times a number to give starting stack")
)

func main() {
	flag.Parse()
	fmt.Println(simulate.Play(*rounds, *players, *min, *multi))
}

