package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedDeckLength := 16
	if len(d) != expectedDeckLength {
		t.Errorf(fmt.Sprint("Expected deck length of ", expectedDeckLength, " but got ", len(d)))
	}

	expectedFirstElement := "Ace of Spades"
	if d[0] != expectedFirstElement {
		t.Errorf(fmt.Sprint("Expected first element in deck is ", expectedFirstElement, " but got ", d[0]))
	}

	expectedLastElement := "Four of Clubs"
	if d[len(d)-1] != expectedLastElement {
		t.Errorf(fmt.Sprint("Expected last element in deck is ", expectedLastElement, " but got ", d[len(d)-1]))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFilename := "_decktesting"
	os.Remove(testFilename)

	deck := newDeck()
	deck.saveToFile(testFilename)

	loadedDeck := newDeckFromFile(testFilename)

	expectedDeckLength := 16
	if len(loadedDeck) != expectedDeckLength {
		t.Errorf(fmt.Sprint("Expected deck length of ", expectedDeckLength, " but got ", len(loadedDeck)))
	}

	os.Remove(testFilename)
}
