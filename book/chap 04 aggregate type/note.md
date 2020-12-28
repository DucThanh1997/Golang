## Array
- Là 1 list có độ dài cố định khi đã khai báo độ dài rồi thì không sửa được
```
var a [3]int
```

- Đây là cách khai báo luôn giá trị
```
var q [3]int = [3]int{1, 2, 3}
```

- Nếu mà khai báo như này thì len của array phụ thuộc vào số phần tử được khai báo
```
q := [...]int{1, 2, 3}
```

- Các array có thể so sánh với nhau nhưng nó phải có cùng len với nhau cơ, khác len sẽ báo lỗi
```
a := [2]int{1, 2}
b := [...]int{1, 2}
c := [2]int{1, 3}
fmt.Println(a == b, a == c, b == c) // "true false false"
d := [3]int{1, 2}
fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
```

## slices
- Là 1 phiên bản thu nhỏ của array và có thể thay đổi được slide

- Không như array thì slice không so sánh được

- Khi vượt quá số cap nó sẽ tạo ra 1 slice với với cap gấp đôi cap cũ rồi copy các giá trị mới sang cap cũ


## map
- Cách để kiểm tra có key ở trong map không
