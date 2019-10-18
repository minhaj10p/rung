package courtpiece

//Game a game of court piece
type Game interface {
	Players() []Player
	DistributeCards() error
}

type game struct {
	players []Player
	deck    Deck
}

const (
	EastPlayer  = "East Player"
	WestPlayer  = "West Player"
	NorthPlayer = "North Player"
	SouthPlayer = "South Player"
)

//NewGame NewGame
func NewGame() Game {

	playerNames := []string{EastPlayer, WestPlayer, NorthPlayer, SouthPlayer}
	var players []Player
	for i := 0; i < 4; i++ {
		players = append(players, NewPlayer(playerNames[i]))
	}
	deck := NewDeck()
	return &game{players: players, deck: deck}
}

func (g *game) Players() []Player {
	return g.players
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
