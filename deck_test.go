package courtpiece_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/minhajuddinkhan/courtpiece"
	"github.com/stretchr/testify/assert"
)

func TestDeckShouldHaveFiftyTwoCards(t *testing.T) {

	deck := courtpiece.NewDeck()
	assert.Equal(t, len(deck.CardsInDeck()), 52)
}

func TestIsCardNotPresentInDeck(t *testing.T) {

	deck := courtpiece.NewDeck()
	card, err := deck.DrawCard(0)
	assert.Nil(t, err)
	assert.False(t, deck.IsCardPresent(card))
}

func TestNewDeckHasFourOfSpades(t *testing.T) {

	deck := courtpiece.NewDeck()
	found := true
	for _, card := range deck.CardsInDeck() {
		if card.House() == courtpiece.Spade && card.Number() == 4 {
			found = true
		}
	}
	assert.True(t, found)
}

func TestAfterDrawingCardFromDeck(t *testing.T) {

	deck := courtpiece.NewDeck()
	card, err := deck.DrawCard(0)
	assert.Nil(t, err)
	assert.False(t, deck.IsCardPresent(card))
	assert.Equal(t, len(deck.CardsInDeck()), 51)

}

func TestIsCardPresentInDeck(t *testing.T) {

	deck := courtpiece.NewDeck()
	card := courtpiece.NewCard(courtpiece.Spade, courtpiece.Ace)
	deck.IsCardPresent(card)
}
func TestDrawCards(t *testing.T) {
	deck := courtpiece.NewDeck()
	cards, err := deck.DrawCards(0, 2)
	assert.Nil(t, err)
	assert.False(t, deck.IsCardPresent(cards[0]))
	assert.False(t, deck.IsCardPresent(cards[1]))
	assert.False(t, deck.IsCardPresent(cards[2]))

}

func TestPutCard(t *testing.T) {

	deck := courtpiece.NewDeck()
	card, err := deck.DrawCard(0)
	assert.Nil(t, err)
	deck.PutCard(card)
	assert.Equal(t, len(deck.CardsInDeck()), 52)
}

func TestPutCards(t *testing.T) {

	deck := courtpiece.NewDeck()
	cards, err := deck.DrawCards(0, 1)
	assert.Nil(t, err)
	assert.Equal(t, len(deck.CardsInDeck()), 52-2)
	err = deck.PutCards(cards)
	assert.Nil(t, err)
	assert.Equal(t, len(deck.CardsInDeck()), 52)
}
func TestAfterShufflingDeck(t *testing.T) {

	deck := courtpiece.NewDeck()
	err := deck.Shuffle(30)
	assert.Nil(t, err)
	assert.Equal(t, len(deck.CardsInDeck()), 52)
	spew.Dump(deck.CardsInDeck())

}
