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
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	//13 ranks *4 suits
	expectedCard := Card{Rank: Four, Suit: Spade}
	if cards[3] != expectedCard {
		t.Error("Expected three of Spades, received:", cards[3])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	//13 ranks *4 suits
	expectedCard := Card{Rank: Four, Suit: Spade}
	if cards[3] != expectedCard {
		t.Error("Expected three of Spades, received:", cards[3])
	}
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
	filter := func(card Card) bool {
		return card.Rank == Two
	}
	cards := New(Filter(filter))
	for i, card := range cards {
		if card.Rank == Two {
			t.Errorf("Expected filtering for Two's but found two at index %v", i)
		}
	}
}

func TestDeck(t *testing.T) {
	numb_decks := rand.Intn(5)
	cards := New(Deck(numb_decks))
	if len(cards) != 13*4*numb_decks {
		t.Errorf("Expected Deck to have length %v; but instead the length is: %v", 13*4*numb_decks, len(cards))
	}
}
