package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// var để thông báo go chúng ta sẽ tạo 1 biến mới
	// card là tên biến
	// string là để thông báo cho go rằng chỉ có loại dữ liệu kiểu string được chấp nhận cho biến này

	// 1 vài kiểu biến khác
	// bool: true, false
	// int: 1,2,3
	// string: "lalala"
	// float: 0.1, 2.3

	// array và slice
	// array có là 1 chuỗi có độ dài được fix cứng
	// slice là array nhưng có thể tăng hoặc giảm độ dài

	// slice và array chỉ có thể chứa 1 kiểu dữ liệu duy nhất
	cards := deck{"Ace of Diamonds", newCard()}

	cards = append(cards, "Six of Spades")
	cards.print()

	deck1 := newDeck()
	fmt.Println("deck: ", deck1)

	// deck11, deck12 := deal(deck1, 20)

	// cards1 := newDeck()

	// cards1.saveToFile("a")

	cards2 := newDeckFromFile("a")

	cards2.shuffle()
	fmt.Println("cards2: ", cards2)
}

// function tên newCard không nhận tham số đầu vào nào và trả về string
func newCard() string {
	return "Five of Diamonds"
}
