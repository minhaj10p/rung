package rung

import (
	"fmt"
	"math/rand"
)

//Deck Deck of cards
type Deck interface {
	//Cards returns all cards
	CardsInDeck() []Card

	//IsCardPresent  Checks if a given card is present in deck
	IsCardPresent(c Card) bool

	//DrawCard draws a card from the deck
	DrawCard(i int) (Card, error)

	//DrawCards draws multiple cards from the deck
	DrawCards(m, n int) ([]Card, error)

	//PutCard puts card in th deck
	PutCard(c Card) error

	//PutCard puts multiple card in th deck
	PutCards(cards []Card) error

	//Shuffle shuffles the deck n times
	Shuffle(n int) error
}

type deck struct {
	cardsInDeck []Card
}

//NewDeck creates a new deck
func NewDeck() Deck {

	var cards []Card
	houses := []string{Spade, Club, Heart, Diamond}

	for i := 0; i < len(houses); i++ {
		house := houses[i]
		for j := 0; j < 13; j++ {
			cards = append(cards, NewCard(house, j))
		}
	}
	return &deck{cardsInDeck: cards}
}

func (d *deck) CardsInDeck() []Card {
	return d.cardsInDeck
}

func (d *deck) DrawCard(i int) (Card, error) {

	if len(d.cardsInDeck) < i {
		return nil, fmt.Errorf("not this many cards")
	}
	card := d.cardsInDeck[i]
	d.cardsInDeck = append(d.cardsInDeck[:i], d.cardsInDeck[i+1:]...)
	return card, nil

}

func (d *deck) DrawCards(m, n int) ([]Card, error) {

	cardCount := len(d.CardsInDeck())
	if cardCount < n || cardCount < m {
		return nil, fmt.Errorf("card index not present in deck")
	}

	x := d.cardsInDeck[m : n+1]
	cards := make([]Card, len(x))
	copy(cards, x)
	d.cardsInDeck = append(d.cardsInDeck[:m], d.cardsInDeck[n+1:]...)
	return cards, nil
}
func (d *deck) Shuffle(n int) error {

	for i := 0; i < n; i++ {

		limit := rand.Intn(51)
		offset := rand.Intn(limit)
		cards, err := d.DrawCards(offset, limit)
		if err != nil {
			return fmt.Errorf("error shuffling cards: %v", err)
		}
		err = d.PutCards(cards)
		if err != nil {
			return err
		}

	}
	return nil
}

func (d *deck) PutCard(c Card) error {

	if d.IsCardPresent(c) {
		return fmt.Errorf("card already in deck")
	}

	d.cardsInDeck = append(d.cardsInDeck, c)
	return nil
}
func (d *deck) PutCards(cards []Card) error {
	var cardsToPut []Card
	for _, c := range cards {
		if d.IsCardPresent(c) {
			return fmt.Errorf("card already in deck")
		}
		cardsToPut = append(cardsToPut, c)
	}
	d.cardsInDeck = append(d.cardsInDeck, cardsToPut...)
	return nil
}

func (d *deck) IsCardPresent(c Card) bool {

	cards := d.CardsInDeck()

	found := make(chan bool)
	done := make(chan bool)

	for _, card := range cards {
		go func(cardToCheck Card, cardInDeck Card) {
			if isSameCard(cardToCheck, cardInDeck) {
				found <- true
				return
			}
			done <- true
		}(c, card)
	}

	cardsChecked := 0
	for {
		select {
		case <-found:
			return true
		case <-done:
			cardsChecked++
			if cardsChecked >= len(cards) {
				return false
			}

		}
	}
}

func isSameCard(c1 Card, c2 Card) bool {
	return c1.House() == c2.House() && c1.Number() == c2.Number()
}
