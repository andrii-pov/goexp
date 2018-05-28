package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of Spades" //automatically finds out the type by :=
	card = newCard()        //reassign - no need in := anymore
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, card)
	fmt.Println(card)
	fmt.Println(cards)
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
