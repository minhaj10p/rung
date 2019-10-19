package rung

import (
	"sync"
)

//Game a game of court piece
type Game interface {
	Players() []Player
	DistributeCards() error
	PlayHand(turn int, trump *string) (Hand, error)
	ShuffleDeck(n int) error
}

type game struct {
	m       *sync.Mutex
	players []Player
	deck    Deck
}

const (
	EastPlayer  = "East Player"
	WestPlayer  = "West Player"
	NorthPlayer = "North Player"
	SouthPlayer = "South Player"
)
const (
	FirstHandForClub = 0
)

//NewGame NewGame
func NewGame() Game {

	playerNames := []string{EastPlayer, WestPlayer, NorthPlayer, SouthPlayer}
	var players []Player
	for i := 0; i < 4; i++ {
		players = append(players, NewPlayer(playerNames[i]))
	}
	deck := NewDeck()
	return &game{players: players, deck: deck, m: &sync.Mutex{}}
}

func (g *game) Players() []Player {
	return g.players
}

func (g *game) ShuffleDeck(n int) error {
	return g.deck.Shuffle(n)
}
func (g *game) DistributeCards() error {

	cards := len(g.deck.CardsInDeck())
	playerTurn := 0
	for i := cards - 1; i >= 0; i-- {
		card, err := g.deck.DrawCard(i)
		if err != nil {
			return err
		}
		g.players[playerTurn].ReceiveCard(card)
		playerTurn++

		if playerTurn == 4 {
			playerTurn = 0
		}

	}

	return nil
}

func (g *game) isFirstHand(turn int) bool {
	return turn == FirstHandForClub
}

func (g *game) isFirstCardTwoOfClubs(c Card) bool {
	return c.House() == Club && c.Number() == Two
}

func (g *game) PlayerToStart() (Player, int) {

	twoClub := NewCard(Club, Two)
	for k, p := range g.players {
		if has, _ := p.HasCard(twoClub); has {
			return p, k
		}
	}
	return nil, -1

}

type Move struct {
	Player Player
	CardAt int
}

func (g *game) PlayHand(turn int, trump *string) (Hand, error) {

	hand := NewHand(nil)
	cardsDelt := 0
	handCh := make(chan Move, 4)
	for _, p := range g.players {
		go func(player Player) {
			cardAt := player.CardOnTable()
			handCh <- Move{Player: player, CardAt: cardAt}
		}(p)
	}
	for cardsDelt < 4 {
		select {
		case move := <-handCh:
			err := hand.AddCard(move.Player, move.CardAt)
			if err != nil {
				return nil, err
			}
			cardsDelt++
		}
	}
	return hand, nil

}
