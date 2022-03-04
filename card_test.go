package deck

import (
	"fmt"
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
