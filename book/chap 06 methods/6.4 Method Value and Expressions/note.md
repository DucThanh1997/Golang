## Method Values and Expressions

Thường thì chúng ta sẽ select và gọi method trong cùng 1 expression như là p.Distance(), nhưng có thể tách riêng 2 việc này ra. Select p.Distance sẽ yields 1 method value. Sau đó function này có thể được gọi mà không cần receiver value

```
p := Point{1, 2}
q := Point{4, 6}

distanceFromP := p.Distance // method value
fmt.Println(distanceFromP(q)) // "5"
var origin Point // {0, 0}
fmt.Println(distanceFromP(origin)) // "2.23606797749979", ;5

scaleP := p.ScaleBy // method value
scaleP(2) // p becomes (2, 4)
scaleP(3) // then (6, 12)
scaleP(10) // then (60, 120)
```

Method Expression sẽ hữu dụng khi bạn cần 1 giá trị đại diện cho 1 lựa chõn trong rất nhiều method cùng thuộc về 1 kiểu từ đó bạn có thể gọi thằng method đó với nhiều receiver khác nhau