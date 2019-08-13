## method
Go không có lớp. Tuy nhiên bạn có thể định nghĩa các method bằng `type`

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
