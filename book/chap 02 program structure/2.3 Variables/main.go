package main

// import (
// 	"fmt"
// )

// func main() {
// 	v := 1
// 	incr(&v) // side effect: v is now 2
// 	fmt.Println(incr(&v)) // "3" (and v is 3)
// }

// func incr(p *int) int {
// 	*p++ // increments what p points to; does not change p
// 	return *p
// }

import (
	// "flag"
	"fmt"
	// "strings"
)
	

// var n = flag.Bool("n", false, "omit trailing newline")
// var sep = flag.String("s", " ", "separator")


func main() {
	// flag.Parse()
	// fmt.Println(strings.Join(flag.Args(), *sep))
	// if !*n {
	// 	fmt.Println("a")
	// }
	p := new(int) // p, of type *int, points to an unnamed int variable
	fmt.Println(p)
	fmt.Println(*p) // "0"
	*p = 2 // sets the unnamed int to 2
	fmt.Println(*p) // "2"

}
	