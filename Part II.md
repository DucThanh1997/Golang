## For
Go có đúng 1 vòng lặp duy nhất đó là vòng for

1 vòng for cơ bản có 3 thành phần chính
- init statement: chạy trước vòng lặp đầu tiên
- condittion expression: nơi để điều kiện lặp
- post statement: chạy khi kết thúc vòng lặp

Và các thành phần này được ngăn cách bởi dấu `;`, và vòng lặp kết thúc khi condittion expression trả về false

```
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

+ Thực ra init statement và post statement không bắt buộc phải có
```
for ; sum < 1000; {
		sum += sum
	}
```

+ Khi bạn bỏ 2 dấu ; đi thì for trở thành while
```
for sum < 1000 {
		sum += sum
	}
```

+ Bạn cũng có thể tạo 1 vòng for vĩnh viễn
```
for {
	}
```

## if 
lệnh `if` cũng giống lệnh for, không cần phải trong dấu ngoặc nhưng phải trong dấu ngoặc `{}`

```
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```

Giống for, if cũng có thể có 1 mệnh đề để chạy trước khi vào mệnh đề điều kiện 

Biến trong mệnh đề chỉ tồn tại cho đến hết vòng if
```
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
 ```

## Switch 

Switch là 1 cách nhanh gọn hơn để viết if else
```
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
```

Bạn cũng có thể viết 1 switch không condittion gọi là switch true
```
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

## Defer 
defer statement sẽ chạy khi các lệnh xung quanh nó chạy xong rồi
```
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
	fmt.Println("hello")
}
```

kể cả nó ở trong vòng lặp


```
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```










































































