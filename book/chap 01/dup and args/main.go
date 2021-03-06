package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var s, sep string
	// // os.Args sẽ lấy các tham số truyền vào khi chạy file
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)


	// s, sep := "", ""
	// for index, arg := range os.Args[:] {
	// 	s += sep + arg
	// 	sep = " "
	// 	fmt.Println(s)
	// 	fmt.Println(index)
	// 	fmt.Println(arg)
	// }
	

	// counts := make(map[string]int)
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	counts[input.Text()]++
	// }
	// // NOTE: ignoring potential errors from input.Err()
	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d\t%s\n", n, line)
	// 	}
	// }


	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines("", os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			fmt.Println("f: ", f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(arg, f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(arg string, f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[arg + "-" + input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	}
	