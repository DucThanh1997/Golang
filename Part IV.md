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

# Struct và Interface

## Struct 
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






















