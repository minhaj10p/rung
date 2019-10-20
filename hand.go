package rung

import (
	"fmt"

	"github.com/minhajuddinkhan/pattay"
)

//Hand Round of a card
type Hand interface {
	//Cards returns the list of cards on a hand
	Cards() []pattay.Card

	//AddCard adds a card at the current hand
	AddCard(playedBy Player, cardAtHandIndex int) error

	//HasCard checks if hand has card
	HasCard(c pattay.Card) (hasCard bool, atIndex int)

	//IsComplete returns if a hand is complete or not
	IsComplete() bool

	//Head returns the player who has thrown the biggest card
	Head() (Player, error)

	//House returns the House/Color of the hand being played
	House() (string, error)

	//Trump returns the trump house of the hand being played
	Trump() (string, error)
}
type hand struct {
	cards     []pattay.Card
	hasPlayed []Player
	house     string
	head      Player
	trump     string
}

func (h *hand) Cards() []pattay.Card {
	return h.cards
}

//NewHand creates a new hand
func NewHand(trump *string) Hand {
	if trump != nil {
		return &hand{trump: *trump}
	}
	return &hand{}
}

func (h *hand) Trump() (string, error) {
	if h.trump == "" {
		return "", fmt.Errorf("trump not declared yet")
	}
	return h.trump, nil
}

func (h *hand) IsComplete() bool {
	return len(h.cards) == 4
}

func (h *hand) HasCard(c pattay.Card) (bool, int) {
	_, at, err := pattay.FindCardInCards(c, h.cards)
	if err != nil {
		return false, -1
	}
	return true, at

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

func (h *hand) isSameHouse(c pattay.Card) bool {
	return c.House() == h.house
}

func (h *hand) validateCard(pl Player, c pattay.Card) error {

	if h.IsComplete() {
		return fmt.Errorf("hand is complete")
	}

	if h.HasAlreadyPlayed(pl) {
		return fmt.Errorf("player %s has already played", pl.Name())
	}

	for _, card := range h.cards {
		if pattay.IsSameCard(card, c) {
			return fmt.Errorf("one hand cannot have two same cards")
		}
	}
	return nil

}

func (h *hand) trumpCardsAtHand() []pattay.Card {

	var cards []pattay.Card
	for _, c := range h.cards {
		if c.House() == h.trump {
			cards = append(cards, c)
		}
	}
	return cards
}

func (h *hand) isTrumpDeclared() bool {
	return h.trump != ""
}

func (h *hand) setHeadForBiggestCard(cards []pattay.Card, c pattay.Card, house string, pl Player) {
	biggestCardAtHand := pattay.GetBiggestCard(cards, house)
	if c.Number() > biggestCardAtHand.Number() {
		h.head = pl
	}

}
func (h *hand) AddCard(pl Player, cardAtHandIndex int) error {

	c, err := pl.DrawCard(cardAtHandIndex)
	if err != nil {
		return err
	}

	if err = h.validateCard(pl, c); err != nil {
		return err
	}
	defer func() {
		h.hasPlayed = append(h.hasPlayed, pl)
		h.cards = append(h.cards, c)
	}()

	if h.IsEmpty() {
		h.house = c.House()
		h.head = pl
		return nil
	}

	if !h.isTrumpDeclared() {
		if !h.isSameHouse(c) {

			if pl.HasHouse(h.house) {
				return fmt.Errorf("player has cards of the same house")
			}

			h.trump = c.House()
			h.head = pl
			return nil

		}
		h.setHeadForBiggestCard(h.cards, c, h.house, pl)
		return nil
	}

	//if trump is running at hand
	if h.trump == h.house {
		if !h.isSameHouse(c) {
			return nil
		}
		h.setHeadForBiggestCard(h.cards, c, h.trump, pl)
		return nil
	}
	if c.House() == h.trump {
		trumpCards := h.trumpCardsAtHand()
		if len(trumpCards) == 0 {
			h.head = pl
			return nil
		}
		h.setHeadForBiggestCard(h.trumpCardsAtHand(), c, h.trump, pl)
	}

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
