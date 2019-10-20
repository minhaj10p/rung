package rung_test

import (
	"testing"

	"github.com/minhajuddinkhan/pattay"
	"github.com/minhajuddinkhan/rung"
	"github.com/stretchr/testify/assert"
)

func TestPlayer_NewPlayerHasIdentity(t *testing.T) {
	player := rung.NewPlayer(rung.SouthPlayer)
	assert.Equal(t, player.Name(), rung.SouthPlayer)
}
func TestPlayer_NewPlayerHasZeroCards(t *testing.T) {

	player := rung.NewPlayer(rung.SouthPlayer)
	assert.Equal(t, len(player.CardsAtHand()), 0)
}

func TestPlayer_ReceiveCardFromDeck(t *testing.T) {
	deck := pattay.NewDeck()
	card, err := deck.DrawCard(0)
	assert.Nil(t, err)

	player := rung.NewPlayer(rung.SouthPlayer)
	err = player.ReceiveCard(card)
	assert.Nil(t, err)
	assert.Equal(t, len(player.CardsAtHand()), 1)
}

func TestPlayer_ThrowErrorOnDrawingCardNotAtHand(t *testing.T) {
	player := rung.NewPlayer(rung.SouthPlayer)
	_, err := player.DrawCard(15)
	assert.NotNil(t, err)
}

func TestPlayer_CannotReceiveCardAlreadyAtHand(t *testing.T) {
	player := rung.NewPlayer(rung.SouthPlayer)
	c1 := pattay.NewCard(pattay.Spade, pattay.Ace)
	err := player.ReceiveCard(c1)
	assert.Nil(t, err)
	err = player.ReceiveCard(c1)
	assert.NotNil(t, err)

}

func TestPlayer_IfPlayerHasCardOfGivenHouse(t *testing.T) {

	player := rung.NewPlayer(rung.SouthPlayer)
	c1 := pattay.NewCard(pattay.Spade, pattay.Ace)
	c2 := pattay.NewCard(pattay.Club, pattay.Ace)
	c3 := pattay.NewCard(pattay.Diamond, pattay.Ace)

	err := player.ReceiveCard(c1)
	assert.Nil(t, err)
	err = player.ReceiveCard(c2)
	assert.Nil(t, err)
	err = player.ReceiveCard(c3)
	assert.Nil(t, err)

	assert.False(t, player.HasHouse(pattay.Heart))
	assert.True(t, player.HasHouse(pattay.Spade))
	assert.True(t, player.HasHouse(pattay.Club))
	assert.True(t, player.HasHouse(pattay.Diamond))

}

func TestPlayer_ThrowCardAndReceiveCardOnTable(t *testing.T) {

	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.WestPlayer)

	for i := 0; i < 10; i++ {
		go func(p1, p2 rung.Player, i int) {
			p1.ThrowCard(i)
			p2.ThrowCard(i)
		}(p1, p2, i)
	}

	count := 0
	doneCh := make(chan interface{})
	for i := 0; i < 20; i++ {
		go func() {
			p1.CardOnTable()
			doneCh <- true
		}()
		go func() {
			p2.CardOnTable()
			doneCh <- true
		}()

	}
	for count < 20 {
		<-doneCh
		count++
	}
	assert.Equal(t, count, 20)
}

func TestPlayer_HasAnySpade(t *testing.T) {

	p := rung.NewPlayer(rung.WestPlayer)
	p.ReceiveCard(pattay.NewCard(pattay.Spade, pattay.Ace))

	card, at, err := p.AnySpade()
	assert.Nil(t, err)
	assert.Equal(t, at, 0)
	assert.Equal(t, card.House(), pattay.Spade)
}
func TestPlayer_HasAnyHeart(t *testing.T) {

	p := rung.NewPlayer(rung.WestPlayer)
	p.ReceiveCard(pattay.NewCard(pattay.Heart, pattay.Ace))

	card, at, err := p.AnyHeart()
	assert.Nil(t, err)
	assert.Equal(t, at, 0)
	assert.Equal(t, card.House(), pattay.Heart)
}
func TestPlayer_HasAnyClub(t *testing.T) {

	p := rung.NewPlayer(rung.WestPlayer)
	p.ReceiveCard(pattay.NewCard(pattay.Club, pattay.Ace))

	card, at, err := p.AnyClub()
	assert.Nil(t, err)
	assert.Equal(t, at, 0)
	assert.Equal(t, card.House(), pattay.Club)
}
func TestPlayer_HasAnyDiamond(t *testing.T) {

	p := rung.NewPlayer(rung.WestPlayer)
	p.ReceiveCard(pattay.NewCard(pattay.Diamond, pattay.Ace))

	card, at, err := p.AnyDiamond()
	assert.Nil(t, err)
	assert.Equal(t, at, 0)
	assert.Equal(t, card.House(), pattay.Diamond)
}

func TestPlayer_ErrorOnPlayerReceivesFourteenCards(t *testing.T) {
	player := rung.NewPlayer(rung.WestPlayer)
	deck := pattay.NewDeck()
	cards, _ := deck.DrawCards(0, 13)
	assert.Len(t, cards, 14)
	for i := 0; i < 13; i++ {
		player.ReceiveCard(cards[i])
	}
	assert.NotNil(t, player.ReceiveCard(cards[13]))
}

func TestPlayer_TakeOutSpadeWhenPlayerDoesntHaveSpade(t *testing.T) {

	player := rung.NewPlayer(rung.SouthPlayer)
	_, _, err := player.AnySpade()
	assert.NotNil(t, err)
}
