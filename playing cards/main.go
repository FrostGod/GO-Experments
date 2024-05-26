package main

// import "fmt"

func main() {
	cards := newDeck()
	// cards.print()
	// cards_received, rem_cards := deal(cards, 3)

	// cards_received.print()
	// rem_cards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// temp_deck := newDeckFromFile("my_card")
	// temp_deck.print()
	cards.shuffle()
	cards.print()
}
