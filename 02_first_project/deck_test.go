package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Tạo 1 kiểu dữ liệu tên là deck
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"nhép", "bích", "cơ", "dô"}

	cardNum := []string{"át", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, Suit := range cardSuits {
		for _, Num := range cardNum {
			cards = append(cards, Num+" "+Suit)
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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	err := ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	return err
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// mỗi lần random cần 1 seed để tạo ra 1 seed mới, mỗi seed sẽ có 1 kiểu random khác nhau
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index, _ := range d {
		newPos := r.Intn(len(d) - 1)
		var temp string
		temp = d[index]
		d[index] = d[newPos]
		d[newPos] = temp
	}
}
