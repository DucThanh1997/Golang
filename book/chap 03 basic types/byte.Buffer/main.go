package main

import (
	"fmt"
	"bytes"
	"strconv"
)

// intsToString is like fmt.Sprintf(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		// ép kiểu sang string rồi ghi vào đoạn buf
		fmt.Fprintf(&buf, "%d", v)
		// buf.WriteString(strconv.Itoa(v))
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}
	