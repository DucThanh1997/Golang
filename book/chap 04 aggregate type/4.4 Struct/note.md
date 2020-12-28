## Struct 
- Là 1 kiểu dữ liệu tổng hợp, nó nhóm các giá trị với nhiều kiểu dữ liệu thành 1 thực thể. Mỗi 1 giá trị được gọi là 1 field. Dưới đây là 1 ví dụ về struct

```
type Employee struct {
    ID          int
    Name        string
    Address     string
    DoB         time.Time
    Position    string
    Salary      int
    ManagerID   int
}

var dilbert Employee

fmt.Println(dilbert.Name)
```

- Thứ tự khai báo các trường là quan trọng, nếu chúng ta thay đổi chúng ta sẽ tạo ra 1 struct mới

- Các trường được export ra ngoài nếu nó có chữ cái đầu viết hoa 

- Struct không thể có 1 field mang chính nó

### Struct Literal
- Chúng ta có thể tạo 1 struct và tạo 1 biến struct như này
```
type Point struct{ X, Y int }

p := Point{1, 2}
```

- Có 2 kiểu literal của struct
    + Kiểu đầu tiên được hiển thị ở trên bắt người dùng phải nhớ chính xác vị trí của field => ít được dùng
    + Kiểu thứ 2 `anim: = gif.GIF {LoopCount: nframes}`. Nếu một trường bị bỏ qua trong loại chữ này, nó sẽ được đặt thành giá trị 0 cho kiểu của nó. Vì tên được cung cấp nên thứ tự của các trường không quan trọng. Hai hình thức không thể được trộn lẫn

- Struct cũng có thể được dùng làm argument trong function. Struct thường được đưa vào function bằng pointer nếu function bắt buộc phải modify cái struct đó

- Có thể so sánh các struct với nhau nhé

### Struct Embedding and Anonymous Fields
Giả sử chúng ta có 2 hàm
```
type Circle struct {
    X, Y, Radius int
}
type Wheel struct {
    X, Y, Radius, Spokes int
}   
```

Chúng ta thấy có sự tưởng đồng giữa các hàm ở đây, nên viết lại thế này:
```
type Point struct {
    X, Y int
}
type Circle struct {
    Center Point
    Radius int
}
type Wheel struct {
    Circle Circle
    Spokes int
}
```
Khi đó accessed 1 field của wheel sẽ kiểu như này
```
var w Wheel
w.Circle.Center.X = 8
w.Circle.Center.Y = 8
w.Circle.Radius = 5
w.Spokes = 20
```

- Go cho phép chúng ta tạo 1 field mà không có tên chỉ có kiểu. Gọi là anonymous fields. Kiểu của field đó phải
là tên của kiểu hoặc pointer đến 1 tên kiểu

- Gọn hơn thì có thể viết như này
```
var w Wheel
w.X = 8 // equivalent to w.Circle.Point.X = 8
w.Y = 8 // equivalent to w.Circle.Point.Y = 8
w.Radius = 5 // equivalent to w.Circle.Radius = 5
w.Spokes = 20
```
Nhưng như này lại không được nhé

```
w=Wheel{8, 8, 5, 20} // compile error: unknown fields
w=Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields
```

- Phải như này mới được
```
w=Wheel{Circle{Point{8, 8}, 5}, 20}
w=Wheel{
    Circle: Circle{
        Point: Point{X: 8, Y: 8},
        Radius: 5,
    },
    Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
}
fmt.Printf("%#v\n", w)
// Output:
// Wheel{Circle:Circle{Point:Point{X:8, Y:8}, Radius:5}, Spokes:20}
w.X = 42
```