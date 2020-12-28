package main

import "fmt"

func main() {
	// cái new ấy nó tạo ra 1 biến địa chỉ 
	p := new(int) // p, of type *int, points to an unnamed int variable
	fmt.Println(p)
	fmt.Println(*p) // "0"
	*p = 2 // sets the unnamed int to 2
	fmt.Println(*p)
}

