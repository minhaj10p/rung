package courtpiece_test

import (
	"testing"

	"github.com/minhajuddinkhan/courtpiece"
	"github.com/stretchr/testify/assert"
)

func TestCardShouldHaveValidNumberAndHouse(t *testing.T) {

	house := "Spades"
	cardNumber := 1
	card := courtpiece.NewCard(house, cardNumber)
	assert.Equal(t, card.House(), house)
	assert.Equal(t, card.Number(), cardNumber)
}
