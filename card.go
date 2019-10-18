package courtpiece

const (
	Two   = iota
	Three = iota
	Four  = iota
	Five  = iota
	Six   = iota
	Seven = iota
	Eight = iota
	Nine  = iota
	Ten   = iota
	Jack  = iota
	Queen = iota
	King  = iota
	Ace   = iota
)

const (
	Spade   = "Spades"
	Diamond = "Diamonds"
	Club    = "Clubs"
	Heart   = "Hearts"
)

//Card Card
type Card interface {
	//House returns house of the card
	House() string
	//Number returns number of the card
	Number() int
}

//NewCard NewCard
func NewCard(house string, cardNumber int) Card {
	return &card{house: house, cardNumber: cardNumber}
}

type card struct {
	house      string
	cardNumber int
}

func (c *card) House() string {
	return c.house
}

func (c *card) Number() int {
	return c.cardNumber
}
