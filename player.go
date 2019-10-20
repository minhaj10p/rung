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

	//Throw enqueues throwing of a card from players deck so the game can receive it
	ThrowCard(cardAt int)

	//CardOnTable receives the card from the queue to be added in the game's current hand
	CardOnTable() int

	//AnySpade returns any spade if it has one
	AnySpade() (Card, int, error)

	//AnyClub returns any club if it has one
	AnyClub() (Card, int, error)

	//AnyHeart any heart if it has one
	AnyHeart() (Card, int, error)

	//AnyDiamond returns any diamond if it has one
	AnyDiamond() (Card, int, error)
}

type player struct {
	cardsAtHand []Card
	name        string
	//TODO:: come up with a better name for the queue
	decisionChannel chan int
	handsWon        []Hand
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
func (p *player) ThrowCard(cardAt int) {
	go func() {
		p.decisionChannel <- cardAt
	}()
}

func (p *player) CardOnTable() int {
	return <-p.decisionChannel
}

func (p *player) anyCardOfHouse(house string) (Card, int, error) {
	for at, c := range p.CardsAtHand() {
		if c.House() == house {
			return c, at, nil
		}
	}
	return nil, -1, fmt.Errorf("player doesn't have any %s cards", house)

}

func (p *player) AnyDiamond() (Card, int, error) {
	return p.anyCardOfHouse(Diamond)
}

func (p *player) AnySpade() (Card, int, error) {
	return p.anyCardOfHouse(Spade)
}

func (p *player) AnyClub() (Card, int, error) {
	return p.anyCardOfHouse(Club)
}

func (p *player) AnyHeart() (Card, int, error) {
	return p.anyCardOfHouse(Heart)
}
