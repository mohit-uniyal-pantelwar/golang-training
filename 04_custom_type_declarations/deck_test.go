package main

import (
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected length of cards: 52, Actual length: %v", len(d))
	}

	if !strings.Contains(d[0], "Ace") {
		t.Errorf("Expected first card to be of Ace")
	}
}
