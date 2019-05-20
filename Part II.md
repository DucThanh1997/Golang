# Biến
## Cấu trúc
Gồm 3 phần 
- Tên biến
- Kiểu dữ liệu
- Giá trị
```
package main
 
import "fmt"
 
func main() {
    var x string = "Hello World"
    fmt.Println(x)
}
```

Để khai báo chúng ta dùng `var` + tên biến + kiểu dữ liệu = giá trị 

## Đặc điểm
- Gán giá trị: Bạn có thể định dạng cấu trúc dữ liệu trước rồi gán giá trị vào sau cũng được
```
var x string
    x  = "Hello World"
    fmt.Println(x);
```

- Giá trị của biến có thể thay đổi được
```
var x string
    x = "first"
    fmt.Println(x)
    x = "second"
    fmt.Println(x)
```

- Gán giá trị nhanh
```
x := "Hello World"
```

- Tên biến bạn có thể đặt thoải mái
Các biến được khai báo ở ngoài func trong func có thể truy cập được nhưng biến được khai báo ở trong func thì chỉ sử dụng được 
ở trong func đó thôi 
```
package main
 
import "fmt"
 
var x string = "Hello World"
 
func main() {
    fmt.Println(x)
}
```

- Hằng số
Khi bạn thêm `const` trước khi tạo biến thì biến đó sẽ không thể bị thay đổi giá trị được nữa
```
package main
 
import "fmt"
 
func main() {
    const x string = "Hello World"
    fmt.Println(x)
}
```
- Khai báo nhiều biến
```
var (
a = 1,
b = 2, 
c = 3,
```














