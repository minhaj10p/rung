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

func TestPlayerInput(t *testing.T) {

	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.WestPlayer)

	for i := 0; i < 10; i++ {
		go func(p1, p2 rung.Player, i int) {
			p1.Input() <- i
			p2.Input() <- i
		}(p1, p2, i)
	}

	count := 0
	for i := 0; i < 20; i++ {
		select {
		case <-p1.Input():
			count++
		case <-p2.Input():
			count++
		}
	}
	assert.Equal(t, count, 20)
}
