package main

import (
	"reflect"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("a")
	db, err := sql.Open("sqlite3", "./foo.db")
	fmt.Println("a: ", reflect.TypeOf(db))
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}