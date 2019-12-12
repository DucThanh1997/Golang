package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// khi khởi tạo sẽ đọc hết dữ liệu trong folder lo
func init() {
	tpl = template.Must(template.ParseGlob("*"))
}
func main() {

	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Thành"}
	err := tpl.ExecuteTemplate(os.Stdout, "index2.html", sages)
	if err !=nil {
		log.Println(err)
	}

	
}