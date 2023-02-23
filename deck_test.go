package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != 40 {
		t.Errorf("Expected the length of 40 but it has %v", len(d))
	}

	s := "Ace of Spades"
	if d[0] != s {
		t.Errorf("Expected %v but got %v", s, d[0])
	}

}
