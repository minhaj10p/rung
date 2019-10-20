package rung

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
	deck          Deck
	handsOnGround []Hand
	handsWon      map[string]int
	ring          Ring
}

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
	r, _ := NewRing(players)

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

func (g *game) PlayHand(turn int, trump *string, lastHead Player) (Hand, error) {

	hand := NewHand(trump)
	cardsDelt := 0

	if isFirstHand(turn) {
		for i, p := range g.players {
			if has, cardAt := p.HasCard(NewCard(Club, Two)); has {
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
		player, err := g.ring.Next()
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

	if isFirstHand(turn) {
		return hand, nil
	}

	head, err := hand.Head()
	if err != nil {
		return nil, err
	}
	g.ring.SetCurrentPlayer(head)

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
