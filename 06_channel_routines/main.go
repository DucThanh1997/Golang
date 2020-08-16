package main

import (
	"fmt"
	"time"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)
	// tạo 1 channel giao tiếp bằng kiểu string
	for _, link := range links {
		// khởi tạo go routine
		go checkLink(link, c)
		// chanel là cái để liên lạc giữa các routine với nhau
		// chanel chỉ cho phép 1 kiểu dữ liệu để giao tiếp thôi

		// muốn cái go routine sử dụng channel thì phải thêm nó vào go routine như 1 param
	}
	// để ý cái channel nó bắn ra cái gì cái đó là l
	for l := range(c) {
		// function literal giống cái annonymous function trong javascript
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("err: ", err)
		// dưới đây là cách hứng dữ liệu của channel từ routine khác
		c <- link
		return
	}
	fmt.Println(link)
	c <- link
	return
}
