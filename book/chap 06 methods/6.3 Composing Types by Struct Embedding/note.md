## Embeded 
```
type Point struct{ X, Y float64 }

type ColoredPoint struct {
    Point
    Color color.RGBA
}
```

- Như tìm hiểu trước chúng ta truy cập từ ColoredPoint vào các field của point thì đối với method cũng vậy

```
red := color.RGBA{255, 0, 0, 255}
blue := color.RGBA{0, 0, 255, 255}
var p = ColoredPoint{Point{1, 1}, red}
var q = ColoredPoint{Point{5, 4}, blue}
fmt.Println(p.Distance(q.Point)) // "5"
p.ScaleBy(2)
q.ScaleBy(2)
fmt.Println(p.Distance(q.Point)) // "10"
```
- Các method của point đã được nâng cấp thành ColoredPoint. Ở cách này, việc embeding cho phép 1 kiểu phức tạp có nhiều method được góp lại từ các trường của kiểu phức tạp đó và mỗi trường cung cấp 1 ít method

- Colorpoint không phải là 1 point mà nó có point và nó còn có thêm 2 method là distance và scaleby được promote lên từ point. Các trường nhúng (embedded field) sẽ bảo compiler gen ra thêm các method kiểu như này

```
func (p ColoredPoint) Distance(q Point) float64 {
    return p.Point.Distance(q)
}
func (p *ColoredPoint) ScaleBy(factor float64) {
    p.Point.ScaleBy(factor)
}
```

Khi mà `Point.Distance` được gọi bởi các wrapper methods này, receiver value của nó sẽ là p.Point không phải p và sẽ không có cách nào có thể access được ColoredPoint ở chỗ thằng Point được embedded

- Kiểu của 1 trường có thể là pointer đến 1 kiểu khác, khi đó các trường của kiểu được trỏ đến sẽ được promote 1 cách gián tiếp. Điều này dẫn đến việc tạo ra 1 level gián tiếp (indirect) mới cho phép chúng ta share các structure cơ bản và đa dạng các relationship giữa các object động (dynamically)

```
type ColoredPoint struct {
    *Point
    Color color.RGBA
}
p := ColoredPoint{&Point{1, 1}, red}
q := ColoredPoint{&Point{5, 4}, blue}
fmt.Println(p.Distance(*q.Point)) // "5"
q.Point = p.Point // p and q now share the same Point
p.ScaleBy(2)
fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
```
