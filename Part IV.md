## method
Go không có class. Tuy nhiên bạn có thể định nghĩa các method bằng `type`

method là 1 hàm với reciever argument khá đặc biệt

```
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```
Cái hàm dưới này được viết riêng cho struct vertex `(v vertex)` để đánh dấu cái đó. Abs là tên hàm. còn cái `v` thay cho self trong python
```
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

Bạn chỉ có thể khai báo một method với receiver nhận có type được định nghĩa trong cùng 1 package

### Pointer receivers
```
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
```

khi bạn chạy cái hàm Scale bạn đã thay đổi v rồi vì nó trỏ đến address memory còn nên sau khi chạy xong giá trị của x là 30 và 40 rồi

### Pointer Functions
Điều tương tự cũng sẽ xảy ra với các function
```
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
```
Như trên thì cái function phải chuyền vào `ScaleFunc(&v, 5) // OK` như này còn nếu như này sẽ lỗi `ScaleFunc(v, 5)`

Trong khi method pointer thì có thể chuyền vào cái gì cũng được

`v.Scale(5)` 

hay
`p := &v
p.Scale(10)` đều okke

## interface
```
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.

func (t T) M() {
	fmt.Println(t.S)
}

 type Q struct {
 	R int
 }
 
 func (q Q) M() {
 	fmt.Println(q.R + 1)
 } 
func main() {
	var i I = T{"hello"}
	i.M()
	i = Q{1}
	i.M()
}
```

như các bạn thấy đầu tiên khai báo 1 interface I sau đó chúng ta khai báo 2 struct Q va T đi cùng nó là 2 cái function. 2 function này có tên là M() đồng nghĩa với việc nó cho M chạy vào

I kiểu Interface là nó có thể trở thành bất kì loại struct nào được cho phép


### Empty interface
Empty interface là kiểu chả có gì trong nó. Nó đơn giản chỉ như này thôi `var i interface{}`

1 empty interface có thể trở thành bất kì type nào bạn muốn 
```
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

### Type assertion
Type assertion cho phép truy cập đến giá trị của interface 

```
var i interface{} = "hello"

s, ok := i.(string)
fmt.Println(s, ok)
```
nếu i có kiểu dữ liệu string thì s sẽ hiện ra hello và ok là true còn nếu không có kiểu dữ liệu string nó sẽ trả về zero value và false

Chúng ta có thể dùng switch để bao hết các kiểu dữ liệu
```
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```

### Stringer
1 trong những interface phổ biến nhất là Stringer định nghĩa bởi package `fmt`
```
type Stringer interface {
    String() string
}
```
Ví dụ:
```
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
```

## Error
Go thể hiện error state qua error value 

Error type built bằng type Stringer
```
type error interface {
	Error() string
}
```

```
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

## Reader 
The io package specifies the io.Reader interface, which represents the read end of a stream of data.

The Go standard library contains many implementations of these interfaces, including files, network connections, compressors, ciphers, and others.

The io.Reader interface has a Read method:

The io package chỉ định io.Reader interface, đại diện cho đầu đọc của luồng dữ liệu.

Thư viện chuẩn Go chứa nhiều triển khai cho các giao diện này, bao gồm files, network connections, compressors, ciphers, and others..

Giao diện io.Reader có phương thức Đọc:
```
func (T) Read(b []byte) (n int, err error)
```




















