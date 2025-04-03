package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type 'deck', which is a slice of string
type deck []string

// This is a receiver function attached to the type deck. Any variable of deck type can invoke this function.
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	d := deck{}

	for _, cardValue := range cardValues {
		for _, cardSuit := range cardSuits {
			d = append(d, fmt.Sprintf("%v of %v", cardValue, cardSuit))
		}
	}

	return d
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toByteSlice() []byte {
	cardConcatenated := strings.Join(d, ",")
	return []byte(cardConcatenated)
}

func (d deck) saveToFile(filepath string) error {
	return os.WriteFile(filepath, d.toByteSlice(), 0666)
}

func (d deck) takeOutAndPutOver(fIndex, lIndex int) deck {
	result := append(d[:fIndex], append(d[lIndex+1:], d[fIndex:lIndex+1]...)...)
	return result
}

func (d *deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)
	times := r.Intn(20)

	for times > 0 {
		firstIndex := r.Intn(len(*d))
		lastIndex := r.Intn(len(*d)-firstIndex) + firstIndex
		*d = d.takeOutAndPutOver(firstIndex, lastIndex)
		times--
	}
}

func newDeckFromFile(filepath string) (deck, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return deck{}, err
	}
	cardConcatenated := string(file)
	return strings.Split(cardConcatenated, ","), nil
}
