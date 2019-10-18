package courtpiece

import (
	"fmt"
)

//Player Player
type Player interface {

	//CardsAtHand returns card at hand
	CardsAtHand() []Card

	//DrawCard draws a card
	DrawCard(i int) (Card, error)

	//ReceiveCard
	ReceiveCard(c Card) error
}

type player struct {
	cardsAtHand []Card
	name        string
}

//NewPlayer NewPlayer
func NewPlayer(name string) Player {
	return &player{cardsAtHand: []Card{}, name: name}
}

func (p *player) CardsAtHand() []Card {
	return p.cardsAtHand
}

func (p *player) DrawCard(i int) (Card, error) {

	if len(p.cardsAtHand) <= i {
		return nil, fmt.Errorf("not this many cards")
	}
	card := p.cardsAtHand[i]
	p.cardsAtHand = append(p.cardsAtHand[:i], p.cardsAtHand[i+1:]...)
	return card, nil
}

func (p *player) ReceiveCard(c Card) error {

	if len(p.cardsAtHand) == 13 {
		return fmt.Errorf("cannot receive more cards. all thirteen at hand")
	}
	p.cardsAtHand = append(p.cardsAtHand, c)
	return nil

}
