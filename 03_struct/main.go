package main

import "fmt"

// đây là cách tạo 1 struct
type contactInfo struct {
	email   string
	zipCode int
}

// trong 1 struct cũng có thể chứa 1 thuộc tính có kiểu là struct khác
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// để như này cũng được
	// contactInfo
}

func main() {
	// cách tạo 1 variable cho 1 struct, đó gọi là embeded struct
	alex := person{
		firstName: "alex",
		lastName: "Anderson",
	}
	fmt.Println(alex)
	// cách khác
	// alex := person{"alex", "anderson"}
	// cách này không hay vì nó dựa vào thứ tự của các thuộc tính trong struct recommend cách đầu tiên


	// đây là cách thứ 3
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

// *person để minh họa và chỉ cho việc hàm này đang làm việc với pointer trỏ đến kiểu person, chỉ với mỗi * nó
// mới thế thôi
func (pointerToPerson *person) updateName (newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	// lấy giá trị của thanh ghi mà chúng ta đang có pointer trỏ đến
}

// Khi bạn gọi hàm khác thì nó sẽ sao chép cái biến hay giá trị đó sang 1 địa chỉ khác rồi sửa ở trong đó


 // Biến địa chỉ thành giá trị bằng dấu *
 // Biến giá trị thành địa chỉ bằng dấu &