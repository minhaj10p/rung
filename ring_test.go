package rung_test

import (
	"testing"

	"github.com/minhajuddinkhan/rung"
	"github.com/stretchr/testify/assert"
)

func TestRingHasFourPlayers(t *testing.T) {
	playOrder := []string{rung.SouthPlayer, rung.WestPlayer, rung.NorthPlayer, rung.EastPlayer}
	var players []rung.Player
	for i := 0; i < 4; i++ {
		players = append(players, rung.NewPlayer(playOrder[i]))
	}
	ring, err := rung.NewRing(players)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(ring.Players()))
}

func TestNext_AfterSouthPlayerNextPlayerIsWest(t *testing.T) {
	playerNames := []string{rung.NorthPlayer, rung.EastPlayer, rung.SouthPlayer, rung.WestPlayer}
	var players []rung.Player
	for i := 0; i < 4; i++ {
		players = append(players, rung.NewPlayer(playerNames[i]))
	}
	ring, err := rung.NewRing(players)
	assert.Nil(t, err)

	ring.SetCurrentPlayer(players[0])
	p, err := ring.Next()
	assert.Nil(t, err)
	assert.Equal(t, p.Name(), rung.EastPlayer)

	p, err = ring.Next()
	assert.Nil(t, err)
	assert.Equal(t, p.Name(), rung.SouthPlayer)

	p, err = ring.Next()
	assert.Nil(t, err)
	assert.Equal(t, p.Name(), rung.WestPlayer)

	p, err = ring.Next()
	assert.Nil(t, err)
	assert.Equal(t, p.Name(), rung.NorthPlayer)

}

func TestIsCurrentPlayerSetInRing(t *testing.T) {
	playerNames := []string{rung.NorthPlayer, rung.EastPlayer, rung.SouthPlayer, rung.WestPlayer}
	var players []rung.Player
	for i := 0; i < 4; i++ {
		players = append(players, rung.NewPlayer(playerNames[i]))
	}
	ring, err := rung.NewRing(players)
	assert.Nil(t, err)
	assert.False(t, false, ring.HasCurrentPlayer())

}

//TODO:: call next without setting current player. expect error
//TODO:: create ring with three players. expect error
//TODO:: Add Get current player tests
//TODO:: Add Set current player tests
//TODO:: test with player name not in ring map. expect error
