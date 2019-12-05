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
	//decks = flag.Int("decks", 6, "number of decks")
)

func main() {
	flag.Parse()
	fmt.Println(simulate.Play(*rounds, *players, 10, 100))
}

