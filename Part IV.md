# Con trỏ
## Vấn đề
```
package main
 
import "fmt"
 
func zero(x int) {
    x = 0
}
 
func main() {
    x := 5
    zero(x)
    fmt.Println(x)
}
```
x in ra ở đây vẫn sẽ là 5 chứ không phải là 0

## Giải quyết
Để làm được như vậy chúng ta phải dùng con trỏ
```
package main
 
import "fmt"
 
func zero(xPtr *int) {
    *xPtr = 0
}
 
func main() {
    x := 5
    zero(&x)
    fmt.Println(x)  
}
```
Để rõ ràng hơn thì
```
package main
 
import "fmt"
 
func main() {
    var x *int 
    var y int
    y = 0
    x = &y
  
    fmt.Println(x)
    fmt.Println(&y)
    fmt.Println(*x)
    fmt.Println(y)
  
    *x = 1
  
    fmt.Println(*x)
    fmt.Println(y)
}
```
Đầu tiên khởi tạo 1 con trỏ x và 1 biến y
Rồi gán cho y = 0 và x = giá trị địa chỉ của biến y

rồi gán cho giá trị 1 vào địa chỉ của biến y thì sẽ thay đổi cả biến x

# Struct  
- Định nghĩa: 
Struct là kiểu object trong python ấy.
```
type Circle struct {
    x float64
    y float64
    z float64
}
 
type Rectangle struct {
    x1 float64
    y1 float64
    x2 float64
    y2 float64
}
```
trên đây là kiểu khởi tạo hình tròn và hình chữ nhật

Viết gọn hơn thì nó sẽ là như này
```
type Circle struct {
    x, y, r float64 
}
 
type Rectangle struct {
    x1, y1, x2, y2 float64
```

- Khai báo biến
```
var c Circle
```
- Khai báo biến và gán giá trị luôn
```
c := Circle{x: 0, y : 0, r : 5}
r := Rectangle{x1 : 0, y1 : 10, x2 : 0, y2 : 10}
```

- Gọi đến các trường
```
fmt.Println(c.x, c.y, c.z)
c.x = 10
c.y = 5
```

- Func liên quan đến struct
```
func circleArea(c Circle) float64 {
    return math.Pi * c.r*c.r
}
 
func rectangleArea(r Rectangle) float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}
```

# Concurency

## Goroutine
là 1 hàm có thể chạy đồng thời với các hàm khác
```
package main
 
import "fmt"
 
func f(n int) {
    for i := 0 ; i < 10 ; i++ {
        fmt.Println(n, ":", i)
    }
}
 
func main() {
    go f(0)
    var input string 
    fmt.Scanln(&input)
}
```
Trong đoạn code trên có 2 hàm goroutine, hàm đầu tiên là hàm main(), hàm thứ hai là hàm f() khi được gọi trong câu lệnh go f(0). Nếu như chúng ta gọi hàm f() một cách bình thường thì khi gọi, hàm main() sẽ phải dừng tất cả mọi thứ lại, đợi cho hàm f() thực hiện công việc của nó xong rồi trả lại quyền điều khiển cho hàm main() thì hàm main() mới tiếp tục công việc của nó.

Trong đoạn code trên chúng ta không gọi hàm f() như bình thường mà chúng ta chuyển lời gọi đó thành một goroutine, như thế sau khi gọi, hàm main() vẫn tiếp tục công việc của nó, hàm f() cũng thực hiện công việc của nó một cách song song với hàm main().

## Chanel
```
package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
```
Đoạn code trên sẽ in dòng chữ “ping” vô số lần cho đến khi có người bấm nút Enter. Trong đó chúng ta dùng 2 hàm goroutine là pinger() và printer() và 1 channel là c. Về cơ bản thì goroutine là các luồng chương trình chạy xuyên suốt, channel có thể coi như là các “ống” truyền dữ liệu qua lại giữa các luồng chương trình đó.

![image](https://user-images.githubusercontent.com/45547213/58221055-568af300-7d3b-11e9-831b-d1bb47c8435e.png)



- Select giống switch nhưng select dùng cho biến chanel


# Package
## Ưu điểm 
 + Giảm thiểu rủi rõ trùng lắp tên hàm. Chẳng hạn như trong gói fmt có hàm Println(), chúng ta có thể định nghĩa một gói khác cũng có hàm Println() nhưng 2 hàm này khác nhau vì chúng nằm ở 2 gói khác nhau.
 + Dễ dàng tổ chức code hơn, do đó giúp chúng ta tìm các hàm cần dùng dễ dàng hơn.
 + Tốc độ biên dịch nhanh, bởi vì trình biên dịch không biên dịch lại code trong các package.

## Tạo 1 package
Bây giờ t sẽ cứu nhân độ thế, hướng dẫn các bạn cách tạo 1 package trong go lang, ngầu lòi luôn

B1: bạn vào gopath mà bạn cài đặt lúc đầu, của t là `C:\Go\src`

B2: tạo 1 thư mục có tên mymath rồi tạo 1 file math.go

B3 trong file math.go gõ các lệnh này vào 
```
package math
 
func Average(xs []float64) float64 {
    total := float64(0)
    for _, x :=  range xs {
        total += x
    }
    return total / float64(len(xs))
}
```

B4: vào file main.go ở ngoài, thử import vào chạy kiểu gì cũng được ha
```
package main
 
import "fmt"
import "myMath/math"
 
func main() {
    xs := []float64{1, 2, 3, 4}
    avg := math.Average(xs)
    fmt.Println(avg)
}
```





















