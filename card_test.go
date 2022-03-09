package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Jack, Suit: Club})
	fmt.Println(Card{Rank: King, Suit: Heart})
	fmt.Println(Card{Suit: Joker})

	//Output:
	//Ace of Hearts
	//Two of Spades
	//Nine of Diamonds
	//Jack of Clubs
	//King of Hearts
	//Joker
}

func TestNew(t *testing.T) {
	cards := New()
	//13 ranks *4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards")
	}
	fmt.Println(cards)
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	//13 ranks *4 suits
	expectedCard := Card{Rank: Four, Suit: Spade}
	if cards[3] != expectedCard {
		t.Error("Expected three of Spades, received:", cards[3])
	}
	fmt.Println(cards)
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	//13 ranks *4 suits
	expectedCard := Card{Rank: Four, Suit: Spade}
	if cards[3] != expectedCard {
		t.Error("Expected three of Spades, received:", cards[3])
	}
	fmt.Println(cards)
}

func TestJokers(t *testing.T) {
	numb_jokers := rand.Intn(10)
	cards := New(Jokers(numb_jokers))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != numb_jokers {
		t.Errorf("Expected %v Jokers but got: %v", numb_jokers, count)
	}
}

func TestFilter(t *testing.T) {

}
