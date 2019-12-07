package simulate

const (
	first = 0
	second = 1
)

func Play(rounds, nPlayers, chips int) []winnings {
	shoe := newShoe()
	var players []*player
	for i := 0; i < nPlayers; i++ {
		p :=  newPlayer(chips)
		players = append(players, &p)
	}
	dealer := newDealer()
	for i := 0; i < rounds; i++ {
		bet(&shoe, players)
		deal(&shoe, &dealer, players)
		if dealer.is21() {
			payout(&dealer, players)
			continue
		}
		play(&shoe, &dealer, players)
		dealer.play(&shoe, players)
		payout(&dealer, players)
		reset(&dealer, players)
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
	d.hand[upcard] = topCard
	for _, p := range players {
		topCard := s.next()
		p.hand[second] = topCard
	}
	topCard = s.next()
	d.hand[second] = topCard
}

func bet(s *shoe, players []*player) {
	for _, p := range players {
		p.initialBet(s)
	}
}

func play(s *shoe, d *dealer, players []*player) {
	for _, p := range players {
		p.play(s, *d)
	}
}

func payout(d *dealer, players []*player) {
	dScore := score(d.hand)
	for _, p := range players {
		if p.done() { continue }
		pScore := score(p.hand)
		if pScore > 21 && dScore > 21 {
			p.tie()
		} else if pScore > 21 {
			p.lose()
		} else if dScore > 21 {
			p.win()
		} else if pScore < dScore {
			p.lose()
		} else if pScore > dScore {
			p.win()
		} else {
			p.tie()
		}
	}
	d.reset()
}

func reset(d *dealer, players []*player) {
	for _, p := range players {
		p.reset()
	}
	d.reset()
}
