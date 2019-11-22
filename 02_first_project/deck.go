package main

import "fmt"

// Tạo 1 kiểu dữ liệu tên là deck
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"nhép", "bích", "cơ", "dô"}

	cardNum := []string{"át", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, Suit := range cardSuits {
		for _, Num := range cardNum {
			cards = append(cards, Num + " " + Suit)
		}
	}
	return cards
}

// d là cái biến của kiểu đó
// func này chứng tỏ nó thuộc về deck nó gần giống this hoặc self
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
	// index của cái phần tử trong card
	// card là giá trị của phần tử đó
	// range cards là slice của cái đó
}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

func toByteSlice(d deck) []byte{

}
