package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// khi khởi tạo sẽ đọc hết dữ liệu trong folder lo
func init() {
	tpl = template.Must(template.ParseGlob("lo/*"))
}
func main() {
	// đọc dữ liệu từ 1 file
	// tpl, err := template.ParseFiles("tpl.gohtml")
	// if err != nil {
	// 	log.Println(err)
	// }
	// // in ra
	// err = tpl.Execute(os.Stdout, nil)
	// if err != nil {
	// 	log.Println(err)
	// }

	// index, err := os.Create("index.html")  
	// if err != nil {
	// 	log.Println(err)
	// }
	
	// // ghi dữ liệu parse được vào index
	// err = tpl.Execute(index, nil)
	// if err != nil {
	// 	log.Println(err)
	// }
	
	// in cái đầu ra
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Println(err)
	}

	// in theo bất kì tên bạn đọc
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}