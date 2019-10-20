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

func TestAfterSouthPlayerNextPlayerIsWest(t *testing.T) {
	playOrder := []string{rung.SouthPlayer, rung.WestPlayer, rung.NorthPlayer, rung.EastPlayer}
	var players []rung.Player
	for i := 0; i < 4; i++ {
		players = append(players, rung.NewPlayer(playOrder[i]))
	}
	ring, err := rung.NewRing(players)
	assert.Nil(t, err)
	player := ring.Next()
	assert.Equal(t, rung.WestPlayer, player.Name())

}

//TODO:: Add Get current player tests
//TODO:: Add Set current player tests
