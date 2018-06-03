package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of Spades" //automatically finds out the type by :=
	card = newCard()        //reassign - no need in := anymore
	cards := newDeck()      // deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, card)
	fmt.Println(card)
	fmt.Println(cards)
	cards.print()
	hand, remainingValues := deal(cards, 5)
	hand.print()
	remainingValues.print()
}

func newCard() string {
	return "Five of Diamonds"
}
