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

func deal(d deck, handSize int) (deck, deck) {
	// bài 19
	// cái ở trong dấu () sẽ chứa các argument truyền vào trong function
	// d[:handSize] lấy từ 0 đến handsize
	// d[handSize:] lấy từ handsize đến hết
	return d[:handSize], d[handSize:]
	// trả về 2 biến deck
}

func (d deck) toString() string {
	// bài 21
	// vì deck nó là kiểu string sẵn rồi nên chỉ cần điền []string(d) là nó về kiểu string
	// strings.Join biến 1 slice chứa string thành 1 đoạn string duy nhất
	// arg1 là 1 biến slice strings 
	// arg2 là đoạn string để phân tách
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {

	// bài 20 và 23
	// ioutil.WriteFile có tác dụng ghi ra file với 
	// argument đầu tiên là tên file
	// []byte là kiểu dữ liệu dạng slice của byte
	
	// type conversion là chuyển từ kiểu dữ liệu này sang kiểu dữ liệu khác

	// cái argument cuối cùng là để set permission cho cái file đấy
	err := ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	return err
}

func newDeckFromFile(filename string) deck {
	// bài 24
	// ioutil.ReadFile để đọc dữ liệu từ file trả về biển error nếu lỗi và 1 slice dạng byte
	// err là biến để gán cho các lỗi, nếu không có lỗi thì nó sẽ là nil
	bs, err := ioutil.ReadFile(filename)

	// check nếu err khác nil thì điền lỗi đó ra
	if err != nil {
		fmt.Println("err: ", err)
		// os.Exit dể thoát khỏi chương trình
		// argument ở đây nếu truyền vào 0 thì chương trình đã thành công còn bh sẽ exit còn nếu khác 0 thì nó fail lòi
		os.Exit(1)
	}
	// bài 25
	// string(bs) biến slice byte về kiểu string
	// strings.Split(string(bs), ",") đưa kiểu string thành kiểu dữ liệu slice []string bằng cách tách bởi dấu ,
	s := strings.Split(string(bs), ",")

	// từ slice string đưa về kiểu dữ liệu deck thì như này deck(s)
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
