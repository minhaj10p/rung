package rung

import (
	"fmt"
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
	Next() (Player, error)

	//SetCurrentPlayer sets current player in the ring
	SetCurrentPlayer(player Player)

	//GetCurrentPlayer gets current player in the ring
	GetCurrentPlayer() (player Player)

	//HasCurrentPlayer return whether current player is set for the ring
	HasCurrentPlayer() bool
}

//NewRing creates a new ring of four card players
func NewRing(players []Player) (Ring, error) {
	if len(players) != 4 {
		return nil, fmt.Errorf("Ring must contain four players")
	}
	r := &ring{
		players: players,
	}
	return r, nil
}

func (r *ring) Players() []Player {
	return r.players
}

func (r *ring) Next() (Player, error) {

	if r.currentPlayer == nil {
		return nil, fmt.Errorf("configuration error, please call SetCurrentPlayer first")
	}
	name := r.currentPlayer.Name()
	for _, p := range r.Players() {
		if p.Name() == TurnPlaySequence[name] {
			r.SetCurrentPlayer(p)
			return p, nil
		}
	}
	return nil, fmt.Errorf("player not found")
}

func (r *ring) SetCurrentPlayer(p Player) {
	r.currentPlayer = p
}
func (r *ring) GetCurrentPlayer() Player {
	return r.currentPlayer
}

func (r *ring) HasCurrentPlayer() bool {
	return r.currentPlayer != nil
}
