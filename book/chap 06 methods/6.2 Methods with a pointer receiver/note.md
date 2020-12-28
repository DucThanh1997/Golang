## Methods with a Pointer Receiver

Vì gọi 1 function sẽ dẫn đến việc tạo 1 bản copy cho mỗi giá trị argument nên nếu function cần update hoặc modify cái argument truyền vào hoặc argument quá lớn và chúng ta không muons copy, Chúng ta sẽ chuyền vào địa chỉ của biến bằng cách sử dụng pointer. Tương tự với method mà cần update receiver variable

```
func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}
```

Trong convention thực tế, nếu 1 cái method nào mà sử dụng pointer receiver thì tất cả các method khác cũng phải dùng pointer receiver, kể cả cái mà không cần nó

Để nhanh gọn chúng ta có thể sử dụng như này
```
p.ScaleBy(2)
```

- Nhưng chúng ta không thể gọi method với pointer receiver bằng các biến chưa có địa chỉ

```
Point{1, 2}.ScaleBy(2) // compile error: can't take address of Point literal
```
