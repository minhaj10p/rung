package rung_test

import (
	"testing"

	"github.com/minhajuddinkhan/rung"
	"github.com/stretchr/testify/assert"
)

func TestCardShouldHaveValidNumberAndHouse(t *testing.T) {

	house := "Spades"
	cardNumber := 1
	card := rung.NewCard(house, cardNumber)
	assert.Equal(t, card.House(), house)
	assert.Equal(t, card.Number(), cardNumber)
}
