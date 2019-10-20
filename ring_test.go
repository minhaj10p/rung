package rung_test

import (
	"testing"

	"github.com/minhajuddinkhan/rung"
	"github.com/stretchr/testify/assert"
)

var defaultPlayersNo = 4

func beforeEach(numberOfPlayers int) (rung.Ring, error) {
	playerNames := []string{rung.NorthPlayer, rung.EastPlayer, rung.SouthPlayer, rung.WestPlayer}
	var players []rung.RingPlayer
	for i := 0; i < numberOfPlayers; i++ {
		pl := rung.NewPlayer(playerNames[i]).(rung.RingPlayer)
		players = append(players, pl)
	}
	return rung.NewRing(players...)

}

func TestRingHasFourPlayers(t *testing.T) {
	ring, err := beforeEach(defaultPlayersNo)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(ring.Players()))
}

func TestRing_NextAfterSouthPlayerNextPlayerIsWest(t *testing.T) {

	ring, err := beforeEach(defaultPlayersNo)
	assert.Nil(t, err)
	players := ring.Players()
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

func TestRing_IsCurrentPlayerSetInRing(t *testing.T) {
	ring, err := beforeEach(defaultPlayersNo)
	assert.Nil(t, err)
	assert.False(t, false, ring.HasCurrentPlayer())

}

func TestRing_NextWithoutSettingCurrentShouldError(t *testing.T) {
	r, err := beforeEach(defaultPlayersNo)
	assert.Nil(t, err)
	_, err = r.Next()
	assert.Error(t, err)
}

func TestRing_CreateRingWithThreePlayers(t *testing.T) {
	_, err := beforeEach(3)
	assert.Error(t, err)
}

func TestRing_GetAndSetPlayers(t *testing.T) {
	r, err := beforeEach(defaultPlayersNo)
	assert.Nil(t, err)
	players := r.Players()
	p1 := players[0]
	assert.Nil(t, r.GetCurrentPlayer())
	r.SetCurrentPlayer(p1)
	assert.Equal(t, p1.Name(), r.GetCurrentPlayer().Name())
}
