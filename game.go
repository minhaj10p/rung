package rung

import (
	"github.com/minhajuddinkhan/pattay"
)

//Game a game of court piece
type Game interface {
	//Players returns the players
	Players() []Player
	//DistributeCards distrubutes card among players of the game
	DistributeCards() error

	//PlayHand begins play the hand
	PlayHand(turn int, trump *string, lastHead Player) (Hand, error)

	//ShuffleDeck shuffes the deck n times
	ShuffleDeck(n int) error

	//HandsOnGround returns the hands on ground that are not won yet.
	HandsOnGround() []Hand

	//HandsWonBy returns the number of hands won by a player
	HandsWonBy(player Player) int
}

type game struct {
	players       []Player
	deck          pattay.Deck
	handsOnGround []Hand
	handsWon      map[string]int
	ring          pattay.Ring
}

const (
	FirstHandForClub = 0
	SecondLastHand   = 11
)

//NewGame NewGame
func NewGame() Game {

	playerNames := []string{EastPlayer, WestPlayer, NorthPlayer, SouthPlayer}
	var players []Player
	for i := 0; i < 4; i++ {
		players = append(players, NewPlayer(playerNames[i]))
	}
	deck := pattay.NewDeck()

	var rp []pattay.RingPlayer
	for i := 0; i < len(players); i++ {
		rp = append(rp, players[i].(pattay.RingPlayer))
	}

	r, _ := pattay.NewRing(rp...)

	return &game{
		ring:     r,
		players:  players,
		deck:     deck,
		handsWon: make(map[string]int, 4),
	}
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
		card, _ := g.deck.DrawCard(i)
		g.players[playerTurn].ReceiveCard(card)
		playerTurn++
		if playerTurn == 4 {
			playerTurn = 0
		}

	}

	return nil
}

func isFirstHand(turn int) bool {
	return turn == FirstHandForClub
}
func isSecondLastHand(turn int) bool {
	return turn == SecondLastHand
}
func canWinHand(turn int) bool {
	if isFirstHand(turn) {
		return false
	}
	if isSecondLastHand(turn) {
		return false
	}
	return true
}

func (g *game) PlayHand(turn int, trump *string, lastHead Player) (Hand, error) {

	hand := NewHand(trump)
	cardsDelt := 0

	if isFirstHand(turn) {
		for i, p := range g.players {
			clubTwo := pattay.NewCard(pattay.Club, pattay.Two)
			if has, cardAt := p.HasCard(clubTwo); has {
				hand.AddCard(p, cardAt)
				g.players = append(g.players[:i], g.players[i+1:]...)
				cardsDelt++
				g.ring.SetCurrentPlayer(p)
				break
			}
		}
	}

	if cardsDelt == 0 {
		g.ring.SetCurrentPlayer(lastHead)
	}

	for i := 0; i < 4-cardsDelt; i++ {
		rp, err := g.ring.Next()
		player := (rp).(Player)
		if err != nil {
			return nil, err
		}
		cardAt := player.CardOnTable()
		err = hand.AddCard(player, cardAt)
		if err != nil {
			return nil, err
		}

	}

	g.handsOnGround = append(g.handsOnGround, hand)
	head, err := hand.Head()
	if err != nil {
		return nil, err
	}
	g.ring.SetCurrentPlayer(head)

	if !canWinHand(turn) {
		return hand, nil
	}

	if head.Name() == lastHead.Name() {
		g.handsWon[lastHead.Name()] = len(g.handsOnGround)
		g.handsOnGround = nil
	}
	return hand, nil

}

func (g *game) HandsOnGround() []Hand {
	return g.handsOnGround
}

func (g *game) HandsWonBy(player Player) int {
	return g.handsWon[player.Name()]
}
