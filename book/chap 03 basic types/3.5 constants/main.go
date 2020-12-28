package main

import "fmt"

type Weekday int

// sử dụng iota để gen const 
// Weekday ở đây = 0 các monday đến tues rồi wed sẽ tăng dần lên
const (
	Sunday Weekday = iota 
	Monday
	Tuesday
	Wednesday
)

func main() {
	fmt.Println("Wed: ", Wednesday)
}