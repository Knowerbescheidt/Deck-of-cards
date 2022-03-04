// this runs go generate and creates the files to get them as string
//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"sort"
)

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

//Minute 9:40
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

//functional Options
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{suit, rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

//New(DefaultSort)

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}
