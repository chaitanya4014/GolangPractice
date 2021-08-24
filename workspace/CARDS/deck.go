package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error Obtained is: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(byteSlice), ","))
}

func (d deck) shuffle() {
	// for i := range d {
	// 	newPos := rand.Intn(len(d) - 1)
	// 	d[i], d[newPos] = d[newPos], d[i]
	// }

	for i := range d {
		src := rand.NewSource(time.Now().UnixNano())
		r := rand.New(src)
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
