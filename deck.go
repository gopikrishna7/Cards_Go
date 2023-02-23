package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamons", "Club", "Heart"}
	cardValues := []string{"Ace", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := fmt.Sprintf("%v of %v", value, suit)
			cards = append(cards, card)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	// []string(d) it is converting our deck type to slice of strings
	// join function joins all the slice of strings to single string with , separator
}

func (d deck) saveFile(filename string) error {

	err := os.WriteFile(filename, []byte(d.toString()), 066)

	return err

}

func newDeckfromfile(filename string) deck {

	data, err := os.ReadFile(filename)
	if err != nil {

		//fmt.Println("Error", err)
		log.Fatal(err)

	}
	// cards:=strings.Split(string(data),",")
	//from read file we got []byte , we convert that to string using string([]byte)
	// spilt the string as slice of strings and convert that to to a type of deck.
	return deck(strings.Split(string(data), ","))

}

func (d deck) shuffle() {
	n := len(d) - 1
	for i := range d {
		//swapping the cards
		d[i], d[rand.Intn(n)] = d[rand.Intn(n)], d[i]
	}
}
