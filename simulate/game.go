package simulate

const (
	first = 0
	second = 1
)

func Play(rounds, nPlayers, minBet, multiplier int) []winnings {
	shoe := newShoe()
	var players []*player
	for i := 0; i < nPlayers; i++ {
		p :=  newPlayer(minBet, multiplier)
		players = append(players, &p)
	}
	dealer := newDealer()
	for i := 0; i < rounds; i++ {
		bet(&shoe, players)
		deal(&shoe, &dealer, players)
		if dealer.is21() {
			payout(&shoe, &dealer, players)
			continue
		}
		play(&shoe, &dealer, players)
		dealer.play(players)
		payout(&shoe, &dealer, players)
	}
	var stats []winnings
	for _, p := range players {
		stats = append(stats, p.results())
	}
	return stats
}


func deal(s *shoe, d *dealer, players []*player) {
	for _, p := range players {
		topCard := s.next()
		p.hand[first] = topCard
	}
	topCard := s.next()
	d.facedown = topCard
	for _, p := range players {
		topCard := s.next()
		p.hand[second] = topCard
	}
	topCard = s.next()
	d.faceup = topCard
}

func bet(s *shoe, players []*player) {
	for _, p := range players {
		p.placeBet(s)
	}
}

func play(s *shoe, d *dealer, players []*player) {
	for _, p := range players {
		p.play(*d)
	}
}

func payout(s *shoe, d *dealer, players []*player) {
	dScore := score(d.facedown, d.faceup)
	for _, p := range players {
		if p.done() { continue }
		pScore := score(p.hand[0], p.hand[1])
		if pScore < dScore {
			p.lose()
		} else if pScore > dScore {
			p.win()
		} else {
			p.tie()
		}
	}
}

func score(fc card, sc card) int {
	if fc == a && sc == a {
		return cardsKey[fc][0] + cardsKey[sc][1] // 1 + 11
	} else if fc == a {
		if cardsKey[fc][1] + cardsKey[sc][0] < 21 {
			return cardsKey[fc][1] + cardsKey[sc][0]
		}
		return cardsKey[fc][0] + cardsKey[sc][0]
	} else if sc == a {
		if cardsKey[fc][0] + cardsKey[sc][1] < 21 {
			return cardsKey[fc][0] + cardsKey[sc][1]
		}
		return cardsKey[fc][0] + cardsKey[sc][0]
	}
	return cardsKey[fc][0] + cardsKey[sc][0]
}