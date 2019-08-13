## Pointers
Go có con trỏ, con trỏ giữ địa chỉ memory của 1 giá trị

pointer zero value là nil

`&` để gán địa chỉ của i vào 1 biến
```
i := 42
p = &i
```

`*` để chỉ giá trị của memory address
```
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

## Structs
Là 1 collection các trường

```
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

### Struct Field
truy cập các phần tử trong

## Pointer to struct
Struct field có thể được gọi qua struct pointer
```
func main() {
	v := Vertex{1, 2}
	p := &v
	fmt.Println(p)
	p.X = 1e9
	fmt.Println(v)
}
```
## Arrays
Cách khai báo 1 arrays

`var a [10]int`: mảng a 10 phần tử có kiểu int

arrays không thể bị resize

### Slices
Cú pháp: a[low:high]
```
func main() {
	 a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	 fmt.Println(a[1:5])
}
```
Slice không chứa các data, nó chỉ mô tả lại những gì có trong array thôi nhưng nếu bạn thay đổi 1 phần tử của slide thì nó sẽ thay đổi phần tử của array gốc
```
func main() {
	 names := [4]string{
		 "John",
		 "Paul",
		 "Geogre",
		 "Ringo",
	 }
	 fmt.Println(names)

	 a := names[0:2]
	 b := names[1:3]
	 fmt.Println(a, b)

	 b[0] = "XXX"
	 fmt.Println(a, b)
	 fmt.Println(names)
}
```
#### Slice literals
Là kiểu slice không thông báo về số lượng phần tử trong nó bạn cứ thêm vào thôi

```
func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```

#### Slice default
Với array a[10]
thì những câu truy vấn trên bằng nhau
```
a[0:10]
a[:10]
a[0:]
a[:]
```

#### Slice length and capacity
length là số phần tử nó chứa `len(a)`

capacity là sô phần tử của array chứa tính từ vị trí đầu tiên của cap trở đi mà không cần quan tâm đến số high đứng sau `cap(a)`

```
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

zero value của slice là nil

#### Creating slide with make
bạn có thể tạo 1 dynamically-sized arrays với hàm `make` 

```
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
```

#### Slice of slice
Slice có thể chứa nhiều loại và có thể chứa cả slice
```
func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

```

#### Appending to a slice
hàm append để thêm các phần tử vào 

Cú pháp
```
var s []int
printSlice(s)

s = append(s, 0)
printSlice(s)
```

### Range
cái này kết hợp với cái for để iterate qua slice
```
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

bạn có thể bỏ value hay index bằng cách
```
for i, _ := range pow
for _, value := range pow
```

## Maps 
map gần giống dict trong python key và value

zero value của map là nil. Nil map không có key, và cũng không có key
```
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```
### Map literals
để gọi phần tử trong map ra cần có key của nó

bạn cũng có thể khai báo ngắn gọn như này
```
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
```

### Mutating map
Insert or update map
```
m[key] = elem
```
Lấy key về
```
elem = m[key]
```
Xóa key 
```
delete(m, key)
```
Check xem key có exist
```
elem, ok = m[key]
```

## Function values
Function cũng là 1 giá trị
```
import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```



































