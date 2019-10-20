package rung

import (
	"fmt"
)

type ring struct {
	players       []RingPlayer
	currentPlayer RingPlayer
	sequence      map[string]string
}

//RingPlayer player of a ring in a game
type RingPlayer interface {
	Name() string
}

//Ring ring of players in a game of cards
type Ring interface {
	// Players returns
	Players() []RingPlayer
	//Next returns player to play next
	Next() (RingPlayer, error)

	//SetCurrentPlayer sets current player in the ring
	SetCurrentPlayer(player RingPlayer)

	//GetCurrentPlayer gets current player in the ring
	GetCurrentPlayer() (player RingPlayer)

	//HasCurrentPlayer return whether current player is set for the ring
	HasCurrentPlayer() bool
}

//NewRing creates a new ring of four card players
func NewRing(players ...RingPlayer) (Ring, error) {
	if len(players) != 4 {
		return nil, fmt.Errorf("Ring must contain four players")
	}
	sequence := make(map[string]string)
	for i := 0; i < len(players)-1; i++ {
		p := players[i]
		sequence[p.Name()] = players[i+1].Name()
	}
	lastPlayer := players[len(players)-1]
	sequence[lastPlayer.Name()] = players[0].Name()

	r := &ring{
		players:  players,
		sequence: sequence,
	}
	return r, nil
}

func (r *ring) Players() []RingPlayer {
	return r.players
}

func (r *ring) Next() (RingPlayer, error) {

	if r.currentPlayer == nil {
		return nil, fmt.Errorf("configuration error, please call SetCurrentPlayer first")
	}
	name := r.currentPlayer.Name()
	for _, p := range r.Players() {
		if p.Name() == r.sequence[name] {
			r.SetCurrentPlayer(p)
			return p, nil
		}
	}
	return nil, fmt.Errorf("player not found")
}

func (r *ring) SetCurrentPlayer(p RingPlayer) {
	r.currentPlayer = p
}
func (r *ring) GetCurrentPlayer() RingPlayer {
	return r.currentPlayer
}

func (r *ring) HasCurrentPlayer() bool {
	return r.currentPlayer != nil
}
