package dataset

import (
	"github.com/minhajuddinkhan/pattay"
	"github.com/minhajuddinkhan/rung"
)

func PlayerWithTwoOfClubs(g rung.Game) (rung.Player, int) {

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
func PLayersWithoutTwoOfClubs(g rung.Game) []rung.Player {

	twoClub := pattay.NewCard(pattay.Club, pattay.Two)

	var without2Clubs []rung.Player
	for _, p := range g.Players() {
		if hasCard, _ := p.HasCard(twoClub); !hasCard {
			without2Clubs = append(without2Clubs, p)
		}
	}
	return without2Clubs

}
