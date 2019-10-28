# Rung

A game engine for courtpiece.

## APIs
#### Game
A new game exports the game instance that implements the following interface
```go
type Game interface {
  
        //Players returns the players
        Players() []pattay.Player
	
        //DistributeCards distrubutes card among players of the game
      	DistributeCards() error

	//PlayHand begins play the hand
	PlayHand(turn int, trump *string, lastHead pattay.Player) (Hand, error)

	//ShuffleDeck shuffes the deck n times
	ShuffleDeck(n int) error

	//HandsOnGround returns the hands on ground that are not won yet.
	HandsOnGround() []Hand
}
```

#### Hand

Hand represents the hand that's been played or is playing. 
A new hand exports the hand instance that implements the following interface

```go
//Hand Round of a card
type Hand interface {

        //Cards returns the list of cards on a hand
	Cards() []pattay.Card

	//AddCard adds a card at the current hand
	AddCard(playedBy pattay.Player, cardAtHandIndex int) error

	//HasCard checks if hand has card
	HasCard(c pattay.Card) (hasCard bool, atIndex int)

	//IsComplete returns if a hand is complete or not
	IsComplete() bool

	//Head returns the player who has thrown the biggest card
	Head() (pattay.Player, error)

	//House returns the House/Color of the hand being played
	House() (string, error)

	//Trump returns the trump house of the hand being played
	Trump() (string, error)
}
```

## Test
```bash
$ go test -v 
```
