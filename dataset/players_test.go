package dataset_test

import (
	"testing"

	"github.com/minhajuddinkhan/pattay"
	"github.com/minhajuddinkhan/rung"
	"github.com/minhajuddinkhan/rung/dataset"
	"github.com/stretchr/testify/assert"
)

func TestClubTwo(t *testing.T) {

	g := rung.NewGame()
	g.DistributeCards()
	p, cardAt := dataset.PlayerWithTwoOfClubs(g)
	hasCard, at := p.HasCard(pattay.NewCard(pattay.Club, pattay.Two))
	assert.True(t, hasCard)
	assert.Equal(t, cardAt, at)
}
