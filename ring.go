package rung

import (
	"fmt"
)

const (
	PlayerOne   = iota
	PlayerTwo   = iota
	PlayerThree = iota
	PlayerFour  = iota
)

var TurnPlaySequence = map[string]string{
	SouthPlayer: WestPlayer,
	WestPlayer:  NorthPlayer,
	NorthPlayer: EastPlayer,
	EastPlayer:  SouthPlayer,
}

type ring struct {
	players       []Player
	currentPlayer Player
}

//Ring ring of players in a game of cards
type Ring interface {
	// Players returns
	Players() []Player
	//Next returns player to play next
	Next() Player

	//SetCurrentPlayer sets current player in the ring
	SetCurrentPlayer(player Player)

	//GetCurrentPlayer gets current player in the ring
	GetCurrentPlayer() (player Player)
}

//NewRing creates a new ring of four card players
func NewRing(players []Player) (Ring, error) {
	if len(players) != 4 {
		return nil, fmt.Errorf("Ring must contain four players")
	}
	r := &ring{
		players:       players,
		currentPlayer: players[0],
	}
	return r, nil
}

func (r *ring) Players() []Player {
	return r.players
}

func (r *ring) Next() Player {

	name := r.currentPlayer.Name()
	for _, p := range r.Players() {
		if p.Name() == TurnPlaySequence[name] {
			return p
		}
	}
	return nil
}

func (r *ring) SetCurrentPlayer(p Player) {
	r.currentPlayer = p
}
func (r *ring) GetCurrentPlayer() Player {
	return r.currentPlayer
}
