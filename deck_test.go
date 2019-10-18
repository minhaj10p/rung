package rung_test

import (
	"testing"

	"github.com/minhajuddinkhan/rung"
	"github.com/stretchr/testify/assert"
)

func TestDeckShouldHaveFiftyTwoCards(t *testing.T) {

	deck := rung.NewDeck()
	assert.Equal(t, len(deck.CardsInDeck()), 52)
}

func TestNewDeckHasFourOfSpades(t *testing.T) {

	deck := rung.NewDeck()
	found := true
	for _, card := range deck.CardsInDeck() {
		if card.House() == rung.Spade && card.Number() == 4 {
			found = true
		}
	}
	assert.True(t, found)
}

func TestIsCardPresentInDeck(t *testing.T) {

	deck := rung.NewDeck()
	card := rung.NewCard(rung.Spade, rung.Ace)
	deck.IsCardPresent(card)
}
func TestIsCardNotPresentInDeck(t *testing.T) {

	deck := rung.NewDeck()
	card, err := deck.DrawCard(0)
	assert.Nil(t, err)
	assert.False(t, deck.IsCardPresent(card))
}

func TestAfterDrawingCardFromDeck(t *testing.T) {

	deck := rung.NewDeck()
	card, err := deck.DrawCard(0)
	assert.Nil(t, err)
	assert.False(t, deck.IsCardPresent(card))
	assert.Equal(t, len(deck.CardsInDeck()), 51)

}

func TestDrawCards(t *testing.T) {
	deck := rung.NewDeck()
	cards, err := deck.DrawCards(0, 2)
	assert.Nil(t, err)
	assert.False(t, deck.IsCardPresent(cards[0]))
	assert.False(t, deck.IsCardPresent(cards[1]))
	assert.False(t, deck.IsCardPresent(cards[2]))

}

func TestPutCard(t *testing.T) {

	deck := rung.NewDeck()
	card, err := deck.DrawCard(0)
	assert.Nil(t, err)
	deck.PutCard(card)
	assert.Equal(t, len(deck.CardsInDeck()), 52)
}

func TestPutCards(t *testing.T) {

	deck := rung.NewDeck()
	cards, err := deck.DrawCards(0, 1)
	assert.Nil(t, err)
	assert.Equal(t, len(deck.CardsInDeck()), 52-2)
	err = deck.PutCards(cards)
	assert.Nil(t, err)
	assert.Equal(t, len(deck.CardsInDeck()), 52)
}
func TestAfterShufflingDeck(t *testing.T) {

	deck := rung.NewDeck()
	err := deck.Shuffle(30)
	assert.Nil(t, err)
	assert.Equal(t, len(deck.CardsInDeck()), 52)

}
