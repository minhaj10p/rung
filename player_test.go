package courtpiece_test

import (
	"testing"

	"github.com/minhajuddinkhan/courtpiece"
	"github.com/stretchr/testify/assert"
)

func TestNewPlayerHasZeroCards(t *testing.T) {

	player := courtpiece.NewPlayer(courtpiece.SouthPlayer)
	assert.Equal(t, len(player.CardsAtHand()), 0)
}

func TestReceiveCardFromDeck(t *testing.T) {
	deck := courtpiece.NewDeck()
	card, err := deck.DrawCard(0)
	assert.Nil(t, err)

	player := courtpiece.NewPlayer(courtpiece.SouthPlayer)
	err = player.ReceiveCard(card)
	assert.Nil(t, err)
	assert.Equal(t, len(player.CardsAtHand()), 1)
}

func TestThrowErrorOnDrawingCardNotAtHand(t *testing.T) {
	player := courtpiece.NewPlayer(courtpiece.SouthPlayer)
	_, err := player.DrawCard(15)
	assert.NotNil(t, err)
}

// func TestReceiveCardFromDeck(t *testing.T) {

// 	player := courtpiece.NewPlayer(courtpiece.SouthPlayer)

// }
