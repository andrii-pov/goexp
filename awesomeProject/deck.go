package main

import "fmt"

type Deck []string
type Mobiles []string

func newDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (m Mobiles) print() {
	for _, mob := range m {
		fmt.Printf("%s \n", mob)
	}

	fmt.Printf("ALL CARDS: \n")
	for _, card := range newDeck() {
		fmt.Printf("%s \n", card)
	}
}
