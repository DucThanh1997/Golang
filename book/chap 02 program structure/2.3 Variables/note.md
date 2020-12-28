## Variable
- `var` tạo 1 biến của 1 type cụ thể nào đó, đặt tên cho nó và đặt giá trị khoeir tạo cho nó
```
var name type = expression
```
Phần `= expression` có thể bỏ qua, nếu bị bỏ qua biến sẽ có giá trị **zero value** (0 với số, false cho bool và "" nếu là string). **zero value** đảm bảo rằng 1 biến luôn có 1 giá trị nào đó

- Khai báo nhiều biến
```
var i, j, k int // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string
```

- Cách khai báo nhanh mà không cần specified type
```
anim := gif.GIF{LoopCount: nframes}
i, j := 0, 1
```

- `:=` là khởi tạo còn `=` là khai báo

- `:=` chỉ hợp lệ khi bên trái của nó có ít nhất 1 biến mới

### Pointer
- Pointer là địa chỉ của cái biến đó hoặc pointer là địa chỉ mà cái biến đó được lưu trữ
- Không phải giá trị nào cũng có địa chỉ nhưng biến nào cũng có
- Với pointer chúng ta có thể đọc và update giá trị của biến 1 cách không trực tiếp mà không cần dùng hoặc biết tên của biến

```
var x int,
```
`&x` là địa chỉ của biến x

```
import (
	"fmt"
)

func main() {
	v := 1
	incr(&v) // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
}

func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}
```

### new
`new(T)` tạo 1 biến không tên thuộc loại T, gán biến zero value của T cho nó rồi trả ra địa chỉ của biến

```
p := new(int) // p, of type *int, points to an unnamed int variable
fmt.Println(*p) // "0"
*p = 2 // sets the unnamed int to 2
fmt.Println(*p) // "2"
```

### vòng đời của biến
Là khoảng thời gian mà biến tồn tại khi mà chương trình đang chạy
- Biến package-level tồn tại suốt quá trình chạy của chương trình

- Ngược lại biến local có thời gian sống khác nhau, 1 biến local được tạo ra mỗi khi mà hàm nó được chạy và nó sẽ sống đến khi không thể động đến được nữa

- Trình biên dịch sẽ chỉ định biến local vào heap hay vào stack. CHúng ta sẽ rõ hơn ở ví dụ dưới

```
var global *int
func f() {
	var x int
	x=1
	global = &x
}

func g() {
	y := new(int)
	*y = 1
}
```
x sẽ được chỉ định vào heap vì nó vẫn có thể được reachable từ biến global kể cả khi f đã return mặc dù được khai báo là biến local trong hàm f

còn với y thì khi hàm g kết thúc chúng ta không truy cập được nó nữa nên nó sẽ được chỉ định vào stack

- Stack là vùng nhớ được quản lý bởi hệ điều hành (các biến trong function thường được khai báo trong đó)
- Heap là vùng nhớ lưu trữ con trỏ được quản lý bởi lập trình viên
