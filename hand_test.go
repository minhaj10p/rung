package rung_test

import (
	"testing"

	"github.com/minhajuddinkhan/pattay"
	"github.com/minhajuddinkhan/rung"
	"github.com/stretchr/testify/assert"
)

func TestHand_EmptyHandShouldHaveZeroCards(t *testing.T) {

	hand := rung.NewHand(nil)
	assert.Equal(t, len(hand.Cards()), 0)
}

func TestHand_EmptyHandShouldHaveNoHouse(t *testing.T) {
	hand := rung.NewHand(nil)
	_, err := hand.House()
	assert.NotNil(t, err)
}
func TestHand_AddCardToHand(t *testing.T) {

	hand := rung.NewHand(nil)
	player := rung.NewPlayer(rung.SouthPlayer)
	card := pattay.NewCard(pattay.Spade, pattay.Ace)
	err := player.ReceiveCard(card)
	assert.Nil(t, err)
	hand.AddCard(player, rung.FirstCardAtHand)
	assert.Equal(t, len(hand.Cards()), 1)
}

func TestHand_NoSamePlayerCanAddToHand(t *testing.T) {

	hand := rung.NewHand(nil)
	p1 := rung.NewPlayer(rung.EastPlayer)
	c1 := pattay.NewCard(pattay.Spade, pattay.Ace)
	c2 := pattay.NewCard(pattay.Club, pattay.Ace)
	err := p1.ReceiveCard(c1)
	assert.Nil(t, err)
	err = p1.ReceiveCard(c2)
	assert.Nil(t, err)

	err = hand.AddCard(p1, rung.FirstCardAtHand)
	assert.Nil(t, err)
	err = hand.AddCard(p1, rung.FirstCardAtHand)
	assert.NotNil(t, err)
}

func TestHand_NoMoreThanFourCardsAtOneHand(t *testing.T) {

	hand := rung.NewHand(nil)
	player1 := rung.NewPlayer(rung.SouthPlayer)
	player2 := rung.NewPlayer(rung.WestPlayer)
	player3 := rung.NewPlayer(rung.EastPlayer)
	player4 := rung.NewPlayer(rung.NorthPlayer)
	player5 := rung.NewPlayer("Wrong player")

	c1 := pattay.NewCard(pattay.Spade, pattay.Ace)
	err := player1.ReceiveCard(c1)
	assert.Nil(t, err)

	c2 := pattay.NewCard(pattay.Spade, pattay.Two)
	err = player2.ReceiveCard(c2)
	assert.Nil(t, err)

	c3 := pattay.NewCard(pattay.Spade, pattay.Three)
	err = player3.ReceiveCard(c3)
	assert.Nil(t, err)

	c4 := pattay.NewCard(pattay.Spade, pattay.Four)
	err = player4.ReceiveCard(c4)
	assert.Nil(t, err)

	c5 := pattay.NewCard(pattay.Spade, pattay.Five)
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

func TestHand_EmptyHandShouldHaveNoHead(t *testing.T) {

	hand := rung.NewHand(nil)
	headPlayer, err := hand.Head()
	assert.Nil(t, headPlayer)
	assert.NotNil(t, err)
}

func TestHand_ColorOfHandShouldBeOfTheFirstCardPlacedOnHand(t *testing.T) {

	hand := rung.NewHand(nil)
	p1 := rung.NewPlayer(rung.WestPlayer)
	c1 := pattay.NewCard(pattay.Diamond, pattay.Ace)
	err := p1.ReceiveCard(c1)
	assert.Nil(t, err)
	err = hand.AddCard(p1, rung.FirstCardAtHand)
	assert.Nil(t, err)
	houseOfHand, err := hand.House()
	assert.Nil(t, err)
	assert.Equal(t, houseOfHand, c1.House())
}

func TestHand_HeadShouldBeOfTheMostPowerfullCardAtHand(t *testing.T) {

	hand := rung.NewHand(nil)
	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.EastPlayer)
	p3 := rung.NewPlayer(rung.WestPlayer)
	p4 := rung.NewPlayer(rung.NorthPlayer)

	c1 := pattay.NewCard(pattay.Spade, pattay.King)
	c2 := pattay.NewCard(pattay.Spade, pattay.Ace)
	c3 := pattay.NewCard(pattay.Spade, pattay.Queen)
	c4 := pattay.NewCard(pattay.Spade, pattay.Jack)

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

func TestHand_PlayerCannotPlayCard_OfDifferentHouseThanHouseOfHand(t *testing.T) {

	hand := rung.NewHand(nil)
	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.EastPlayer)

	spadeAce := pattay.NewCard(pattay.Spade, pattay.Ace)
	spadeKing := pattay.NewCard(pattay.Spade, pattay.King)
	heartAce := pattay.NewCard(pattay.Heart, pattay.Ace)

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

func TestHand_PlayerCanMakeTrump(t *testing.T) {

	hand := rung.NewHand(nil)
	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.NorthPlayer)
	p3 := rung.NewPlayer(rung.EastPlayer)
	p4 := rung.NewPlayer(rung.WestPlayer)

	c1 := pattay.NewCard(pattay.Spade, pattay.Ace)
	c2 := pattay.NewCard(pattay.Spade, pattay.King)
	c3 := pattay.NewCard(pattay.Spade, pattay.Queen)
	c4 := pattay.NewCard(pattay.Heart, pattay.Three)

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
	assert.Equal(t, h, pattay.Spade)
	assert.Equal(t, trump, pattay.Heart)

}

func TestHand_CannotMakeTrumpAgainIfTrumpAlreadyDeclared(t *testing.T) {

	trump := pattay.Spade
	hand := rung.NewHand(&trump)
	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.NorthPlayer)
	p3 := rung.NewPlayer(rung.EastPlayer)
	p4 := rung.NewPlayer(rung.WestPlayer)

	c1 := pattay.NewCard(pattay.Spade, pattay.Three)
	c2 := pattay.NewCard(pattay.Spade, pattay.King)
	c3 := pattay.NewCard(pattay.Spade, pattay.Queen)
	c4 := pattay.NewCard(pattay.Heart, pattay.Ace)

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
	assert.Equal(t, h, pattay.Spade)
	assert.Equal(t, trump, pattay.Spade)
	assert.Equal(t, player.Name(), p2.Name())

}

func TestHand_CardNotWithPlayerAddedInHand(t *testing.T) {
	p := rung.NewPlayer(rung.SouthPlayer)
	hand := rung.NewHand(nil)
	err := hand.AddCard(p, 2)
	assert.NotNil(t, err)
}

func TestHand_CutByTrumpCard(t *testing.T) {

	trump := pattay.Diamond
	hand := rung.NewHand(&trump)
	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.NorthPlayer)
	p3 := rung.NewPlayer(rung.EastPlayer)
	p4 := rung.NewPlayer(rung.WestPlayer)

	c1 := pattay.NewCard(pattay.Spade, pattay.Three)
	c2 := pattay.NewCard(pattay.Spade, pattay.King)
	c3 := pattay.NewCard(pattay.Spade, pattay.Queen)
	c4 := pattay.NewCard(pattay.Diamond, pattay.Two)

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

func TestHand_CutByBiggerTrumpCard(t *testing.T) {

	trump := pattay.Diamond
	hand := rung.NewHand(&trump)
	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.NorthPlayer)
	p3 := rung.NewPlayer(rung.EastPlayer)
	p4 := rung.NewPlayer(rung.WestPlayer)

	c1 := pattay.NewCard(pattay.Spade, pattay.Three)
	c2 := pattay.NewCard(pattay.Diamond, pattay.Ace)
	c3 := pattay.NewCard(pattay.Spade, pattay.Queen)
	c4 := pattay.NewCard(pattay.Diamond, pattay.King)

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

func TestHand_SameCardsInOneHand(t *testing.T) {
	hand := rung.NewHand(nil)
	p1 := rung.NewPlayer(rung.SouthPlayer)
	p2 := rung.NewPlayer(rung.WestPlayer)
	c := pattay.NewCard(pattay.Spade, pattay.Ace)
	p1.ReceiveCard(c)
	p2.ReceiveCard(c)

	hand.AddCard(p1, rung.FirstCardAtHand)
	err := hand.AddCard(p2, rung.FirstCardAtHand)
	assert.Error(t, err)
}

func TestHand_HandHasCard(t *testing.T) {
	hand := rung.NewHand(nil)
	has, index := hand.HasCard(pattay.NewCard(pattay.Spade, pattay.Ace))
	assert.False(t, has)
	assert.Equal(t, -1, index)
}

func TestHand_EmptyTrump(t *testing.T) {
	hand := rung.NewHand(nil)
	tr, err := hand.Trump()
	assert.Empty(t, tr)
	assert.Error(t, err)
}
