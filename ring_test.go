package rung_test

import (
	"testing"

	"github.com/minhajuddinkhan/rung"
	"github.com/stretchr/testify/assert"
)

var defaultPlayersNo = 4

func beforeEach(numberOfPlayers int) (rung.Ring, error) {
	playerNames := []string{rung.NorthPlayer, rung.EastPlayer, rung.SouthPlayer, rung.WestPlayer}
	var players []rung.Player
	for i := 0; i < numberOfPlayers; i++ {
		players = append(players, rung.NewPlayer(playerNames[i]))
	}
	return rung.NewRing(players)

}

func TestRingHasFourPlayers(t *testing.T) {
	ring, err := beforeEach(defaultPlayersNo)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(ring.Players()))
}

func TestNext_AfterSouthPlayerNextPlayerIsWest(t *testing.T) {

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

func TestIsCurrentPlayerSetInRing(t *testing.T) {
	ring, err := beforeEach(defaultPlayersNo)
	assert.Nil(t, err)
	assert.False(t, false, ring.HasCurrentPlayer())

}

func TestNextWithoutSettingCurrentShouldError(t *testing.T) {
	r, err := beforeEach(defaultPlayersNo)
	assert.Nil(t, err)
	_, err = r.Next()
	assert.Error(t, err)
}

func TestCreateRingWithThreePlayers(t *testing.T) {
	_, err := beforeEach(3)
	assert.Error(t, err)
}

func TestGetAndSetPlayers(t *testing.T) {
	r, err := beforeEach(defaultPlayersNo)
	assert.Nil(t, err)
	players := r.Players()
	p1 := players[0]
	assert.Nil(t, r.GetCurrentPlayer())
	r.SetCurrentPlayer(p1)
	assert.Equal(t, p1.Name(), r.GetCurrentPlayer().Name())
}

func TestWithInvalidPlayer(t *testing.T) {
	playerNames := []string{"some", "guy", "i", "dk"}
	var players []rung.Player
	for i := 0; i < len(playerNames); i++ {
		players = append(players, rung.NewPlayer(playerNames[i]))
	}
	r, err := rung.NewRing(players)
	assert.Nil(t, err)
	r.SetCurrentPlayer(players[0])
	_, err = r.Next()
	assert.Error(t, err)

}

//TODO:: test with player name not in ring map. expect error
