package main

import "fmt"

func main() {
	var card string = "Ace of Spades"
	anotherCard := "Five of Diamonds"
	fmt.Printf("Hello World, and here is the keyword: \n")
	fmt.Printf(card + "\n")
	fmt.Printf("And another hidden card is: %s \n", anotherCard)
	card = newCard()
	fmt.Printf("updated card name:\n%s\n", card)

	mobiles := Mobiles{"Iphone", newMobile()}
	mobiles = append(mobiles, "Pixel")
	mobiles.print()
}

func newCard() string {
	return "just another unnamed card"
}

func newMobile() string {
	return "Android"
}
