package dataset

import (
	"github.com/minhajuddinkhan/pattay"
	"github.com/minhajuddinkhan/rung"
)

func PlayerWithTwoOfClubs(g rung.Game) (pattay.Player, int) {

	players := g.Players()

	for _, p := range players {
		for k, c := range p.CardsAtHand() {
			if c.House() == pattay.Club && c.Number() == pattay.Two {
				return p, k
			}
		}
	}
	return nil, -1

}

func PLayersWithoutTwoOfClubs(g rung.Game) []pattay.Player {

	twoClub := pattay.NewCard(pattay.Club, pattay.Two)

	var without2Clubs []pattay.Player
	for _, p := range g.Players() {
		if hasCard, _ := p.HasCard(twoClub); !hasCard {
			without2Clubs = append(without2Clubs, p)
		}
	}
	return without2Clubs

}

//PlayerWithAceOfSpade PlayerWithAceOfSpade
func PlayerWithAceOfSpade(g rung.Game) (pattay.Player, int) {
	aceOfSpade := pattay.NewCard(pattay.Spade, pattay.Ace)
	return PlayerWithCard(g, aceOfSpade)
}

//PlayersWithoutAceOfSpade PlayersWithoutAceOfSpade
func PlayersWithoutAceOfSpade(g rung.Game) []pattay.Player {
	twoClub := pattay.NewCard(pattay.Spade, pattay.Ace)
	return PlayersWithoutCard(g, twoClub)
}

func PlayerWithCard(g rung.Game, card pattay.Card) (pattay.Player, int) {
	for _, p := range g.Players() {
		if has, at := p.HasCard(card); has {
			return p, at
		}
	}
	return nil, -1
}

func PlayersWithoutCard(g rung.Game, card pattay.Card) []pattay.Player {
	var without []pattay.Player
	for _, p := range g.Players() {
		if hasCard, _ := p.HasCard(card); !hasCard {
			without = append(without, p)
		}
	}
	return without

}
