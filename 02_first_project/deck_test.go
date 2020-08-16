package main

import (
	"fmt"
	"testing"
	"os"
)


// đây là cách tạo 1 test kiểm tra xem có đủ bài khi tạo 1 bộ bài hay không
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		fmt.Println(len(d))
		t.Errorf("Thiếu bài, có %v lá", len(d))
	}

	if d[0] != "át nhép" {
		fmt.Println(d[0])
		t.Errorf("Lá đầu tiên không phải át nhép")
	}
}

func TestSaveToDeckAndNewDeckTestFromFile( t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck() 
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")


	if len(loadedDeck) != 52 {
		t.Errorf("Thiếu bài, có %v lá", len(deck))
	}

	os.Remove("_decktesting")
}