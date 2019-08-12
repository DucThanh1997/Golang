## Packages
Mỗi chương trình go bắt đầu với 1 package 

```
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```
Đoạn code trên bắt đầu với package main

Đoạn code này bắt đầu với package `main` và import path là `fmt` và `math/rand`

Nếu bạn để math/rand thì nó chỉ nhận rand thôi không nhận math

## Import
Có 2 cách để viết import
- C1: 
```
import (
	"fmt"
	"math/rand"
)
```
- C2:
```
import "fmt"
import "math"
```

T thích cách 2 hơn và nhớ rằng tên package phải trong dấu ngoặc kép

## Exported names
Trong go, các hàm các biến bạn lấy ra từ các library bạn đã import phải được viết hoa chữ đầu tiên

Ví dụ
```
math.Pi
```
còn nếu là:
```
math.pi
```
nó sẽ lỗi ngay. 

## Functions
Đây là cách khai báo 1 hàm
```
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```
nhớ rằng kiểu biến đi sau tên biến và biến nó trả về cái gì phải định nghĩa kiểu cho nó luôn
hoặc bạn có thể khai báo kiểu như này và hàm có thể không cần nhận tham số nào hoặc nhiều tham số, tùy vào bạn muốn
```
func add(x, y int) int {
	return x + y
}
```
Ngoài ra hàm cũng có thể trả về nhiều giá trị
```
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

### naked return
```
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```
Cái return ở đây chống thì nó sẽ hiểu là auto trả về x, y

## Variables
Đây là cách khai báo biến
```
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

bắt đầu bắt buộc phải có var sau đó là tên biến rồi kiểu biến

bạn cũng có thể set giá trị cho chúng luôn
```
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```

Còn 1 cách nhanh để khai báo
```
i := 3
```

Đây là các kiểu dữ liệu trong go
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```
### zero values
1 biến khi khai báo chưa được set giá trị sẽ thành zero value
- 0 cho loại num
- false cho bool
- "" cho string

### ép kiểu 
đây là cách ép kiểu 
```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```
ngắn gọn hơn
```
i := 42
f := float64(i)
u := uint(f)
```
### Gán biến
Khi biến không được gán kiểu dữ liệu thì nó sẽ nhận kiểu dữ liệu của biến nó được gán 
```
var i int
j := i // j is an int
```

nếu không sẽ dựa vào go
```
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

### Constants
biến constant được gán như variable, có thể là bool, num hay string nó không thể được khai báo kiểu `:=`
```
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```





















































