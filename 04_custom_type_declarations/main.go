package main

import (
	"fmt"
)

func main() {

	myDeck := newDeck()

	// myDeck.print()

	hand, _ := myDeck.deal(23)
	// fmt.Println(hand)
	// fmt.Println(remainingCards)

	// filepath := "./04_custom_type_declarations/MyCards.txt"

	// err := hand.saveToFile(filepath)
	// if err != nil {
	// 	log.Println("Unable to save file: ", err)
	// }

	// deck, err := newDeckFromFile(filepath)
	// if err != nil {
	// 	log.Println("Unable to read file: ", err)
	// }

	fmt.Println(hand)
	hand.shuffle()
	fmt.Println(hand)

}
