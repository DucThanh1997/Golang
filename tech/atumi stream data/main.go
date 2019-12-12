package main

import (
	"log"
	"os"

	"github.com/vladimirvivien/automi/collectors"
	"github.com/vladimirvivien/automi/stream"
)


func main() {
	// tạo 1 stream chứa sẵn emiter data
	strm := stream.New(
	   []rune("B世!ぽ@opqDQRS#$%^...O6PTnVWXѬYZbcef7ghijCklrAstvw"),
	)
	// lọc data
	strm.Filter(func(item rune) bool {
	   return item >= 65 && item < (65+26)
	})
	// xử lí data
	strm.Map(func(item rune) string {
	   return string(item)
	})
	// đưa vào đầu ra
	strm.Into(collectors.Writer(os.Stdout))
   // mở cái stream
   if err := <-strm.Open(); err != nil {
	 log.Fatal(err)
   }
 }

 // cái này để xử lí stream data. Đã có sẵn template