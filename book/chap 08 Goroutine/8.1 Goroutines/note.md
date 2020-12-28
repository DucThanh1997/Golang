## Goroutine
Trong Go, mỗi hành động thực hiện đồng thời được gọi là `goroutine`. Ban đầu có thể hiểu `goroutine` giống như thread.

Khi 1 chương trình chạy, chỉ có 1 go routine được chạy đó là hàm main. Nên chúng ta gọi nó là main goroutine. `goroutine` mới được tạo bởi câu lệnh `go`. Câu lệnh `go` là 1 hàm thông thường. Một câu lệnh gọi go sẽ tạo 1 `goroutine` mới vào chạy cái hàm nó gọi ở trong đó

```
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
```
Khi chương trình kết thúc tất cả các goroutines sẽ bị terminated và chương trình exit. Ngoài cách này ra không có cách nào để 1 goroutine stop 1 goroutine khác, nhưng các goroutine có thể giao tiếp với nhau