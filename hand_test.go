package rung_test

import (
	"testing"

	"github.com/minhajuddinkhan/rung"
	"github.com/stretchr/testify/assert"
)

func TestEmptyHandShouldHaveZeroCards(t *testing.T) {

	hand := rung.NewHand(nil)
	assert.Equal(t, len(hand.Cards()), 0)
}

func TestEmptyHandShouldHaveNoHouse(t *testing.T) {
	hand := rung.NewHand(nil)
	_, err := hand.House()
	assert.NotNil(t, err)
}
func TestAddCardToHand(t *testing.T) {

	hand := rung.NewHand(nil)
	player := rung.NewPlayer(rung.SouthPlayer)
	card := rung.NewCard(rung.Spade, rung.Ace)
	err := player.ReceiveCard(card)
	assert.Nil(t, err)
	hand.AddCard(player, rung.FirstCardAtHand)
	assert.Equal(t, len(hand.Cards()), 1)
}

func TestNoSamePlayerCanAddToHand(t *testing.T) {

	hand := rung.NewHand(nil)
	p1 := rung.NewPlayer(rung.EastPlayer)
	c1 := rung.NewCard(rung.Spade, rung.Ace)
	c2 := rung.NewCard(rung.Club, rung.Ace)
	err := p1.ReceiveCard(c1)
	assert.Nil(t, err)
	err = p1.ReceiveCard(c2)
	assert.Nil(t, err)

	err = hand.AddCard(p1, rung.FirstCardAtHand)
	assert.Nil(t, err)
	err = hand.AddCard(p1, rung.FirstCardAtHand)
	assert.NotNil(t, err)
}

func TestNoMoreThanFourCardsAtOneHand(t *testing.T) {

	hand := rung.NewHand(nil)
	player1 := rung.NewPlayer(rung.SouthPlayer)
	player2 := rung.NewPlayer(rung.WestPlayer)
	player3 := rung.NewPlayer(rung.EastPlayer)
	player4 := rung.NewPlayer(rung.NorthPlayer)
	player5 := rung.NewPlayer("Wrong player")

	c1 := rung.NewCard(rung.Spade, rung.Ace)
	err := player1.ReceiveCard(c1)
	assert.Nil(t, err)

	c2 := rung.NewCard(rung.Spade, rung.Two)
	err = player2.ReceiveCard(c2)
	assert.Nil(t, err)

	c3 := rung.NewCard(rung.Spade, rung.Three)
	err = player3.ReceiveCard(c3)
	assert.Nil(t, err)

	c4 := rung.NewCard(rung.Spade, rung.Four)
	err = player4.ReceiveCard(c4)
	assert.Nil(t, err)

	c5 := rung.NewCard(rung.Spade, rung.Five)
	err = player5.ReceiveCard(c5)
	assert.Nil(t, err)

	err = hand.AddCard(player1, rung.FirstCardAtHand)
	assert.Nil(t, err)
	err = hand.AddCard(player2, rung.FirstCardAtHand)
	assert.Nil(t, err)
	err = hand.AddCard(player3, rung.FirstCardAtHand)
	assert.Nil(t, err)
	err = hand.AddCard(player4, rung.FirstCardAtHand)
	assert.Nil(t, err)
	err = hand.AddCard(player5, rung.FirstCardAtHand)
	assert.NotNil(t, err)

}

func TestEmptyHandShouldHaveNoHead(t *testing.T) {

	hand := rung.NewHand(nil)
	headPlayer, err := hand.Head()
	assert.Nil(t, headPlayer)
	assert.NotNil(t, err)
}

func TestColorOfHandShouldBeOfTheFirstCardPlacedOnHand(t *testing.T) {

	hand := rung.NewHand(nil)
	p1 := rung.NewPlayer(rung.WestPlayer)
	c1 := rung.NewCard(rung.Diamond, rung.Ace)
	err := p1.ReceiveCard(c1)
	assert.Nil(t, err)
	err = hand.AddCard(p1, rung.FirstCardAtHand)
	assert.Nil(t, err)
	houseOfHand, err := hand.House()
	assert.Nil(t, err)
	assert.Equal(t, houseOfHand, c1.House())
}

func TestHeadShouldBeOfTheMostPowerfullCardAtHand(t *testing.T) {

	hand := rung.NewHand(nil)
	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.EastPlayer)
	p3 := rung.NewPlayer(rung.WestPlayer)
	p4 := rung.NewPlayer(rung.NorthPlayer)

	c1 := rung.NewCard(rung.Spade, rung.King)
	c2 := rung.NewCard(rung.Spade, rung.Ace)
	c3 := rung.NewCard(rung.Spade, rung.Queen)
	c4 := rung.NewCard(rung.Spade, rung.Jack)

	p1.ReceiveCard(c1)
	p2.ReceiveCard(c2)
	p3.ReceiveCard(c3)
	p4.ReceiveCard(c4)

	hand.AddCard(p1, rung.FirstCardAtHand)
	hand.AddCard(p2, rung.FirstCardAtHand)
	hand.AddCard(p3, rung.FirstCardAtHand)
	hand.AddCard(p4, rung.FirstCardAtHand)

	head, err := hand.Head()
	assert.Nil(t, err)
	assert.Equal(t, head.Name(), p2.Name())

}

func TestPlayerCannotPlayCard_OfDifferentHouseThanHouseOfHand(t *testing.T) {

	hand := rung.NewHand(nil)
	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.EastPlayer)

	spadeAce := rung.NewCard(rung.Spade, rung.Ace)
	spadeKing := rung.NewCard(rung.Spade, rung.King)
	heartAce := rung.NewCard(rung.Heart, rung.Ace)

	//p1 has spadeAce and heartAce
	p1.ReceiveCard(spadeAce)
	p1.ReceiveCard(heartAce)

	//p2 has spadeKing
	p2.ReceiveCard(spadeKing)

	//p2 plays spadeKing as first card
	hand.AddCard(p2, rung.FirstCardAtHand)

	//p1 plays heartAce on a spade hand eve
	err := hand.AddCard(p1, rung.SecondCardAtHand)
	assert.NotNil(t, err)

}

func TestPlayerCanMakeTrump(t *testing.T) {

	hand := rung.NewHand(nil)
	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.NorthPlayer)
	p3 := rung.NewPlayer(rung.EastPlayer)
	p4 := rung.NewPlayer(rung.WestPlayer)

	c1 := rung.NewCard(rung.Spade, rung.Ace)
	c2 := rung.NewCard(rung.Spade, rung.King)
	c3 := rung.NewCard(rung.Spade, rung.Queen)
	c4 := rung.NewCard(rung.Heart, rung.Three)

	p1.ReceiveCard(c1)
	p2.ReceiveCard(c2)
	p3.ReceiveCard(c3)
	p4.ReceiveCard(c4)

	hand.AddCard(p1, rung.FirstCardAtHand)
	hand.AddCard(p2, rung.FirstCardAtHand)
	hand.AddCard(p4, rung.FirstCardAtHand)
	hand.AddCard(p3, rung.FirstCardAtHand)

	h, _ := hand.House()
	trump, _ := hand.Trump()
	assert.Equal(t, h, rung.Spade)
	assert.Equal(t, trump, rung.Heart)

}

func TestCannotMakeTrumpAgainIfTrumpAlreadyDeclared(t *testing.T) {

	trump := rung.Spade
	hand := rung.NewHand(&trump)
	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.NorthPlayer)
	p3 := rung.NewPlayer(rung.EastPlayer)
	p4 := rung.NewPlayer(rung.WestPlayer)

	c1 := rung.NewCard(rung.Spade, rung.Three)
	c2 := rung.NewCard(rung.Spade, rung.King)
	c3 := rung.NewCard(rung.Spade, rung.Queen)
	c4 := rung.NewCard(rung.Heart, rung.Ace)

	p1.ReceiveCard(c1)
	p2.ReceiveCard(c2)
	p3.ReceiveCard(c3)
	p4.ReceiveCard(c4)

	hand.AddCard(p1, rung.FirstCardAtHand)
	hand.AddCard(p2, rung.FirstCardAtHand)
	hand.AddCard(p4, rung.FirstCardAtHand)
	hand.AddCard(p3, rung.FirstCardAtHand)

	h, _ := hand.House()
	trump, _ = hand.Trump()
	player, _ := hand.Head()
	assert.Equal(t, h, rung.Spade)
	assert.Equal(t, trump, rung.Spade)
	assert.Equal(t, player.Name(), p2.Name())

}

func TestCardNotWithPlayerAddedInHand(t *testing.T) {
	p := rung.NewPlayer(rung.SouthPlayer)
	hand := rung.NewHand(nil)
	err := hand.AddCard(p, 2)
	assert.NotNil(t, err)
}

func TestCutByTrumpCard(t *testing.T) {

	trump := rung.Diamond
	hand := rung.NewHand(&trump)
	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.NorthPlayer)
	p3 := rung.NewPlayer(rung.EastPlayer)
	p4 := rung.NewPlayer(rung.WestPlayer)

	c1 := rung.NewCard(rung.Spade, rung.Three)
	c2 := rung.NewCard(rung.Spade, rung.King)
	c3 := rung.NewCard(rung.Spade, rung.Queen)
	c4 := rung.NewCard(rung.Diamond, rung.Two)

	p1.ReceiveCard(c1)
	p2.ReceiveCard(c2)
	p3.ReceiveCard(c3)
	p4.ReceiveCard(c4)

	hand.AddCard(p1, rung.FirstCardAtHand)
	hand.AddCard(p2, rung.FirstCardAtHand)
	hand.AddCard(p3, rung.FirstCardAtHand)
	hand.AddCard(p4, rung.FirstCardAtHand)

	x, err := hand.Head()
	assert.Nil(t, err)
	assert.Equal(t, p4.Name(), x.Name())

}

func TestCutByBiggerTrumpCard(t *testing.T) {

	trump := rung.Diamond
	hand := rung.NewHand(&trump)
	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.NorthPlayer)
	p3 := rung.NewPlayer(rung.EastPlayer)
	p4 := rung.NewPlayer(rung.WestPlayer)

	c1 := rung.NewCard(rung.Spade, rung.Three)
	c2 := rung.NewCard(rung.Diamond, rung.Ace)
	c3 := rung.NewCard(rung.Spade, rung.Queen)
	c4 := rung.NewCard(rung.Diamond, rung.King)

	p1.ReceiveCard(c1)
	p2.ReceiveCard(c2)
	p3.ReceiveCard(c3)
	p4.ReceiveCard(c4)

	hand.AddCard(p1, rung.FirstCardAtHand)
	hand.AddCard(p2, rung.FirstCardAtHand)
	hand.AddCard(p3, rung.FirstCardAtHand)
	hand.AddCard(p4, rung.FirstCardAtHand)

	x, err := hand.Head()
	assert.Nil(t, err)
	assert.Equal(t, p2.Name(), x.Name())

}

func TestSameCardsInOneHand(t *testing.T) {
	hand := rung.NewHand(nil)
	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.WestPlayer)
	c := rung.NewCard(rung.Spade, rung.Ace)
	p1.ReceiveCard(c)
	p2.ReceiveCard(c)

	hand.AddCard(p1, rung.FirstCardAtHand)
	err := hand.AddCard(p2, rung.FirstCardAtHand)
	assert.Error(t, err)
}

func TestHandHasCard(t *testing.T) {
	hand := rung.NewHand(nil)
	has, index := hand.HasCard(rung.NewCard(rung.Spade, rung.Ace))
	assert.False(t, has)
	assert.Equal(t, -1, index)
}

func TestEmptyTrump(t *testing.T) {
	hand := rung.NewHand(nil)
	tr, err := hand.Trump()
	assert.Empty(t, tr)
	assert.Error(t, err)
}
