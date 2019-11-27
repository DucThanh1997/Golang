package main

import "fmt"

func main() {
	// tạo 1 map tên color có key là kiểu string và value là kiểu string
	// key và value chỉ có thể được gán bởi 1 kiểu giá trị
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// dưới đây là 2 cách tạo khác
	// var colors2 map[string]string
	// hoặc
	// colors3 := make(map[string]string)

	// add thêm vào như này
	// colors3["lo"] = "la"

	// xóa
	// delete(colors3, "lo")

	printMap(colors)


	
}

// lướt qua map
func printMap(c map[string] string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}