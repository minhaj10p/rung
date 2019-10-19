package rung

import (
	"fmt"
)

const (
	FirstCardAtHand      = iota
	SecondCardAtHand     = iota
	ThirdCardAtHand      = iota
	FourthCardAtHand     = iota
	FifthCardAtHand      = iota
	SixthCardAtHand      = iota
	SeventhCardAtHand    = iota
	EidthCardAtHand      = iota
	NinthCardAtHand      = iota
	TenthCardAtHand      = iota
	EleventhCardAtHand   = iota
	TwelvthCardAtHand    = iota
	ThirteenthCardAtHand = iota
)

//Player Player
type Player interface {

	//Name returns name of the player
	Name() string

	//CardsAtHand returns card at hand
	CardsAtHand() []Card

	//DrawCard draws a card
	DrawCard(i int) (Card, error)

	//ReceiveCard receives a card
	ReceiveCard(c Card) error

	//HasHouse returns if a player has any card of the given house
	HasHouse(house string) bool

	//HasCard(c Card) bool
	HasCard(c Card) (hasCard bool, cardAtIndex int)
	//Throw throw returns a channel identifying the decision of what card to throw
	Input() chan int
}

type player struct {
	cardsAtHand     []Card
	name            string
	decisionChannel chan int
}

//NewPlayer NewPlayer
func NewPlayer(name string) Player {
	return &player{cardsAtHand: []Card{}, name: name, decisionChannel: make(chan int)}
}

func (p *player) Name() string {
	return p.name
}

func (p *player) HasCard(c Card) (bool, int) {
	for cardAtIndex, card := range p.cardsAtHand {
		if isSameCard(c, card) {
			return true, cardAtIndex
		}
	}
	return false, -1
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

func (p *player) alreadyAtHand(c Card) bool {
	for _, atHand := range p.cardsAtHand {
		if isSameCard(atHand, c) {
			return true
		}
	}
	return false
}
func (p *player) ReceiveCard(c Card) error {

	if len(p.cardsAtHand) == 13 {
		return fmt.Errorf("cannot receive more cards. all thirteen at hand")
	}
	if p.alreadyAtHand(c) {
		return fmt.Errorf("cannot receive a card it already has")
	}

	p.cardsAtHand = append(p.cardsAtHand, c)
	return nil

}

func (p *player) HasHouse(house string) bool {
	for _, c := range p.CardsAtHand() {
		if c.House() == house {
			return true
		}
	}
	return false
}
func (p *player) Input() chan int {
	return p.decisionChannel
}
