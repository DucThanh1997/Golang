package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}
	// // tạo 1 byte slice
	// bs := make([]byte, 99999)

	// // cái hàm read này đưa mọi kiểu dữ liệu về dạng byte slice để có thể đọc cho dễ, nó rất hay xuất hiện
	// // trong các standard package của go, nó nhận 1 byte slice vào rồi đưa dữ liệu vào trong đó và trả ra số lượng
	// // byte và số lỗi
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWritter{}
	io.Copy(lw, resp.Body)
	
}


func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println(len(bs))
	return 1, nil
}