// this runs go generate and creates the files to get them as string
//go:generate stringer -type=Suit,Rank
package deck

import "fmt"

type Suit uint8

//stopped at stringer 8:50
//iota make increments the constants starts with 0 and then 1 for diamonds etc.
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker //special case
)

//[...] ist ein array und kein slice
var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

// with the blank Rank we can set the value of Two three etc. to the matching cards
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New() []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{suit, rank})
		}
	}
	return cards
}
