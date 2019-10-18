package rung_test

import (
	"testing"

	"github.com/minhajuddinkhan/rung"
	"github.com/stretchr/testify/assert"
)

func TestGameHasFourPlayers(t *testing.T) {

	game := rung.NewGame()
	assert.Equal(t, len(game.Players()), 4)
}

func TestEachPlayerHasZeroCardsBeforeDistribution(t *testing.T) {

	game := rung.NewGame()
	players := game.Players()

	for _, player := range players {
		assert.Equal(t, len(player.CardsAtHand()), 0)
	}

}

func TestEachPlayerHasThirteenCardsAfterDistribution(t *testing.T) {
	game := rung.NewGame()
	err := game.DistributeCards()
	assert.Nil(t, err)
	players := game.Players()
	for _, p := range players {
		assert.Equal(t, len(p.CardsAtHand()), 13)
	}

}

func TestNoTwoPlayersHaveSameCard(t *testing.T) {

	game := rung.NewGame()
	err := game.DistributeCards()
	assert.Nil(t, err)
	players := game.Players()

	secondPlayer := players[1]

	cardWithfirstPlayer := players[0].CardsAtHand()[0]
	playerOneHasAceOfSpade := false
	playerTwoHasAceOfSpade := false

	for _, card := range secondPlayer.CardsAtHand() {
		if card.House() == cardWithfirstPlayer.House() && cardWithfirstPlayer.Number() == rung.Ace {
			playerTwoHasAceOfSpade = true
		}
	}

	assert.NotEqual(t, playerOneHasAceOfSpade, playerTwoHasAceOfSpade)

}
