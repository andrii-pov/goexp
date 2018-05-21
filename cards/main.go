package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of Spades" //automatically finds out the type by :=
	card = newCard()        //reassign - no need in := anymore
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, card)
	fmt.Println(card)
	fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
