## Arrays
- Là 1 list có độ dài cố định khi đã khai báo độ dài rồi thì không sửa được nên nó hiếm khi được sử dụng. Đây là cách khai báo 1 array và cách truy cập các phần tử có trong array
```
var a [3]int // array of 3 integers
fmt.Println(a[0]) // print the first element
fmt.Println(a[len(a)-1]) // print the last element, a[2]
```

- Cách lướt qua từng phần tử 1
```
// Print the indices and elements.
for i, v := range a {
fmt.Printf("%d %d\n", i, v)
}
// Print the elements only.
for _, v := range a {
fmt.Printf("%d\n", v)
}
```

- Bình thường khi mới được khai báo thì biến array mới sẽ được set về zero-value của kiểu array đó
```
var q [3]int = [3]int{1, 2, 3}
var r [3]int = [3]int{1, 2}
fmt.Println(r[2]) // "0"
```

- Nếu mà sử dụng `...` thì biến chiều dài của array được xác định bởi số phần tử khởi tạo bởi nó

```
q := [...]int{1, 2, 3}
fmt.Printf("%T\n", q) // "[3]int"
```

- Array cũng có thể chứa 1 list bao gồm index và value pair như này
```
type Currency int
const (
    USD Currency = iota
    EUR
    GBP
    RMB
)
symbol := [...]string{USD: "$", EUR: "9", GBP: "!", RMB: "0"}
fmt.Println(RMB, symbol[RMB]) // "3 0"
```

- Khởi tạo 1 mảng có 100 phần tử và phần tử cuối cùng bằng -1 còn các phần tử khác do không được khởi tạo nên = 0
```
r := [...]int{99: -1}
```

- Mảng cũng so sánh được với nhau nhưng phải cùng số lượng phần tử trong array thì mới được
```
a := [2]int{1, 2}
b := [...]int{1, 2}
c := [2]int{1, 3}
fmt.Println(a == b, a == c, b == c) // "true false false"
d := [3]int{1, 2}
fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
```

```
s := []int{0, 1, 2, 3, 4, 5}
```

- Cách khởi tạo slice khá giống với cách khởi tạo mảng, đầu tiên chúng ta tạo 1 array với số lượng được đo bằng số phần tử ở bên phải rồi khởi tạo 1 slice trỏ đến cái mảng đó

- Không như array slice không so sánh được

- Zero value của slice là 1 slice với length = 0 và cap = 0 


### Append
- Dùng để thêm các phần tử vào slice

- Các bước để bắt đầu
    + Đầu tiên check xem còn đủ cap để chứa phần tử mới không 
    + Nếu còn đủ, nó sẽ mở rộng slice bằng cách khởi tạo 1 slice mới lớn hơn (slice đó vẫn thuộc cùng array với slice cũ)
    + Copy phần tử y mới vào
    + Nếu không đủ chỗ thì sẽ phải tạo 1 array mới lớn hơn array cũ rồi copy các phần tử của slice sang rồi lấy các phần tử của slice cũ sang và add phần tử mới vào
    + Khi nới cap thường thì cap sẽ được x2 ở slice mới