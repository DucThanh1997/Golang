package main

import (
	"bytes"
	"fmt"
)

func initToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
		fmt.Println("a: ", buf.String())
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(initToString([]int{1, 2, 3})) // "[1, 2, 3]"
}