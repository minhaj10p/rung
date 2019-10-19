package rung_test

import (
	"testing"

	"github.com/minhajuddinkhan/rung"
	"github.com/stretchr/testify/assert"
)

func TestNewPlayerHasIdentity(t *testing.T) {
	player := rung.NewPlayer(rung.SouthPlayer)
	assert.Equal(t, player.Name(), rung.SouthPlayer)
}
func TestNewPlayerHasZeroCards(t *testing.T) {

	player := rung.NewPlayer(rung.SouthPlayer)
	assert.Equal(t, len(player.CardsAtHand()), 0)
}

func TestReceiveCardFromDeck(t *testing.T) {
	deck := rung.NewDeck()
	card, err := deck.DrawCard(0)
	assert.Nil(t, err)

	player := rung.NewPlayer(rung.SouthPlayer)
	err = player.ReceiveCard(card)
	assert.Nil(t, err)
	assert.Equal(t, len(player.CardsAtHand()), 1)
}

func TestThrowErrorOnDrawingCardNotAtHand(t *testing.T) {
	player := rung.NewPlayer(rung.SouthPlayer)
	_, err := player.DrawCard(15)
	assert.NotNil(t, err)
}

func TestCannotReceiveCardAlreadyAtHand(t *testing.T) {
	player := rung.NewPlayer(rung.SouthPlayer)
	c1 := rung.NewCard(rung.Spade, rung.Ace)
	err := player.ReceiveCard(c1)
	assert.Nil(t, err)
	err = player.ReceiveCard(c1)
	assert.NotNil(t, err)

}

func TestIfPlayerHasCardOfGivenHouse(t *testing.T) {

	player := rung.NewPlayer(rung.SouthPlayer)
	c1 := rung.NewCard(rung.Spade, rung.Ace)
	c2 := rung.NewCard(rung.Club, rung.Ace)
	c3 := rung.NewCard(rung.Diamond, rung.Ace)

	err := player.ReceiveCard(c1)
	assert.Nil(t, err)
	err = player.ReceiveCard(c2)
	assert.Nil(t, err)
	err = player.ReceiveCard(c3)
	assert.Nil(t, err)

	assert.False(t, player.HasHouse(rung.Heart))
	assert.True(t, player.HasHouse(rung.Spade))
	assert.True(t, player.HasHouse(rung.Club))
	assert.True(t, player.HasHouse(rung.Diamond))

}
