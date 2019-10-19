package dataset_test

import (
	"testing"

	"github.com/minhajuddinkhan/rung"
	"github.com/minhajuddinkhan/rung/dataset"
	"github.com/stretchr/testify/assert"
)

func TestClubTwo(t *testing.T) {

	g := rung.NewGame()
	g.DistributeCards()
	p, cardAt := dataset.PlayerWithTwoOfClubs(g)
	hasCard, at := p.HasCard(rung.NewCard(rung.Club, rung.Two))
	assert.True(t, hasCard)
	assert.Equal(t, cardAt, at)
}

// func TestClubWithoutTwo(t *testing.T) {

// 	g := rung.NewGame()
// 	g.DistributeCards()
// 	players := dataset.PLayersWithoutTwoOfClubs(g)
// 	for _, x := range players {

// 	}
// 	hasCard, at := p.HasCard(rung.NewCard(rung.Club, rung.Two))
// 	assert.False(t, hasCard)
// 	assert.Equal(t, at, -1)
// }
