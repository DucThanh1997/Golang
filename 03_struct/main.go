package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// để như này cũng được
	// contactInfo
}

func main() {
	alex := person{firstName: "alex",
		lastName: "Anderson"}
	fmt.Println(alex)

	// john giờ là kiểu zero value
	var john person

	// gán vào
	john.firstName = "john"
	john.lastName = "soap"
	john.contact = contactInfo{
		email:   "aaa",
		zipCode: 999,
	}
	// lấy địa chỉ trong thanh nhớ rồi gán vào nó
	johnPointer := &john
	johnPointer.updateName("jonny")
	john.print()

	// chúng ta cũng có thể thu gọn như này, go cho phép sử dụng shortcut như này
	john.updateName("Thành")
	john.print()
}


func (p person) print() {
	fmt.Printf("%+v", p)
}

// *person để minh họa và chỉ cho việc hàm này đang làm việc với pointer trỏ đến kiểu person
func (p *person) updateName (newFirstName string) {
	(*p).firstName = newFirstName
	// lấy giá trị của thanh ghi mà chúng ta đang có pointer trỏ đến
}

// Khi bạn gọi hàm khác thì nó sẽ sao chép cái biến hay giá trị đó sang 1 địa chỉ khác rồi sửa ở trong đó