package rung

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

//Hand Round of a card
type Hand interface {
	//Cards returns the list of cards on a hand
	Cards() []Card

	//AddCard adds a card at the current hand
	AddCard(playedBy Player, cardAtHandIndex int) error

	//IsComplete returns if a hand is complete or not
	IsComplete() bool

	//Head returns the player who has thrown the biggest card
	Head() (Player, error)

	//House returns the House/Color of the hand being played
	House() (string, error)
}
type hand struct {
	cards     []Card
	hasPlayed []Player
	house     string
	head      Player
}

//NewHand creates a new hand
func NewHand() Hand {
	return &hand{}
}

func (h *hand) Cards() []Card {
	return h.cards
}

func (h *hand) IsComplete() bool {
	return len(h.cards) == 4
}

func (h *hand) HasAlreadyPlayed(pl Player) bool {

	for _, player := range h.hasPlayed {
		if player.Name() == pl.Name() {
			return true
		}
	}
	return false

}

func (h *hand) IsEmpty() bool {
	return len(h.cards) == 0
}

func (h *hand) isSameHouse(c Card) bool {
	return c.House() == h.house
}
func (h *hand) AddCard(pl Player, cardAtHandIndex int) error {

	if h.IsComplete() {
		return fmt.Errorf("hand is complete")
	}
	if h.HasAlreadyPlayed(pl) {
		return fmt.Errorf("player %s has already played", pl.Name())
	}
	c, err := pl.DrawCard(cardAtHandIndex)
	if err != nil {
		return err
	}

	for _, card := range h.cards {
		if isSameCard(card, c) {
			return fmt.Errorf("one hand cannot have two same cards")
		}
	}
	if h.IsEmpty() {
		h.house = c.House()
		h.head = pl
		h.hasPlayed = append(h.hasPlayed, pl)
		h.cards = append(h.cards, c)
		return nil
	}

	if !h.isSameHouse(c) {
		houseOfHand := h.house
		if pl.HasHouse(houseOfHand) {
			return fmt.Errorf("not the same house. please play the card with same house")
		}

		spew.Dump("TRUMP!")
	}
	for _, cardAtHand := range h.cards {
		if c.Number() > cardAtHand.Number() {
			h.head = pl
			break
		}
	}

	h.hasPlayed = append(h.hasPlayed, pl)
	h.cards = append(h.cards, c)
	return nil
}

func (h *hand) Head() (Player, error) {
	if h.IsEmpty() {
		return nil, fmt.Errorf("no head because no card has been played yet")
	}

	return h.head, nil
}

func (h *hand) House() (string, error) {
	if h.IsEmpty() {
		return "", fmt.Errorf("no house because no card has been played yet")
	}
	return h.house, nil
}
