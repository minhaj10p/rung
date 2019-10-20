package rung_test

import (
	"testing"

	"github.com/minhajuddinkhan/rung"
	"github.com/stretchr/testify/assert"
)

func TestDeck_ShouldHaveFiftyTwoCards(t *testing.T) {

	deck := rung.NewDeck()
	assert.Equal(t, len(deck.CardsInDeck()), 52)
}

func TestDeck_HasFourOfSpades(t *testing.T) {

	deck := rung.NewDeck()
	found := true
	for _, card := range deck.CardsInDeck() {
		if card.House() == rung.Spade && card.Number() == 4 {
			found = true
		}
	}
	assert.True(t, found)
}

func TestDeck_IsCardPresentInDeck(t *testing.T) {

	deck := rung.NewDeck()
	card := rung.NewCard(rung.Spade, rung.Ace)
	assert.True(t, deck.IsCardPresent(card))

}
func TestDeck_IsCardNotPresentInDeck(t *testing.T) {

	deck := rung.NewDeck()
	card, err := deck.DrawCard(0)
	assert.Nil(t, err)
	assert.False(t, deck.IsCardPresent(card))
}

func TestDeck_AfterDrawingCardFromDeck(t *testing.T) {

	deck := rung.NewDeck()
	card, err := deck.DrawCard(0)
	assert.Nil(t, err)
	assert.False(t, deck.IsCardPresent(card))
	assert.Equal(t, len(deck.CardsInDeck()), 51)

}

func TestDeck_DrawCards(t *testing.T) {
	deck := rung.NewDeck()
	cards, err := deck.DrawCards(0, 2)
	assert.Nil(t, err)
	assert.False(t, deck.IsCardPresent(cards[0]))
	assert.False(t, deck.IsCardPresent(cards[1]))
	assert.False(t, deck.IsCardPresent(cards[2]))

}

func TestDeck_PutCard(t *testing.T) {

	deck := rung.NewDeck()
	card, err := deck.DrawCard(0)
	assert.Nil(t, err)
	deck.PutCard(card)
	assert.Equal(t, len(deck.CardsInDeck()), 52)
}

func TestDeck_PutCards(t *testing.T) {

	deck := rung.NewDeck()
	cards, err := deck.DrawCards(0, 1)
	assert.Nil(t, err)
	assert.Equal(t, len(deck.CardsInDeck()), 52-2)
	err = deck.PutCards(cards)
	assert.Nil(t, err)
	assert.Equal(t, len(deck.CardsInDeck()), 52)
}
func TestDeck_AfterShufflingDeck(t *testing.T) {

	deck := rung.NewDeck()
	err := deck.Shuffle(30)
	assert.Nil(t, err)
	assert.Equal(t, len(deck.CardsInDeck()), 52)

}

func TestDeck_DrawCardNotPresentInDeck(t *testing.T) {

	deck := rung.NewDeck()
	_, err := deck.DrawCard(52)
	assert.Error(t, err)
}

func TestDeck_DrawMoreCardsThenInDeck(t *testing.T) {
	deck := rung.NewDeck()
	_, err := deck.DrawCards(0, 53)
	assert.Error(t, err)
}

func TestDeck_DrawCardAndShuffle(t *testing.T) {
	deck := rung.NewDeck()
	deck.DrawCards(0, 51)
	assert.Error(t, deck.Shuffle(5))
}

func TestDeck_PutCardAlreadyPresent(t *testing.T) {
	deck := rung.NewDeck()
	err := deck.PutCard(rung.NewCard(rung.Spade, rung.Ace))
	assert.Error(t, err)
}
func TestDeck_PutMultipleCardAlreadyPresent(t *testing.T) {
	deck := rung.NewDeck()
	cards := []rung.Card{rung.NewCard(rung.Spade, rung.Ace)}
	err := deck.PutCards(cards)
	assert.Error(t, err)
}
func TestDeck_InvalidGetQueryCard(t *testing.T) {
	deck := rung.NewDeck()
	_, err := deck.DrawCards(2, 1)
	assert.Error(t, err)
}
