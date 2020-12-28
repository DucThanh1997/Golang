## Slices
- Slices đại diện cho các chuỗi có độ dài thay đổi mà các phần tử của chúng đều có cùng kiểu.

- Slices []T, trong đó các phần tử có kiểu T; nó trông giống như một arrays không có kích thước.

- Arrays và slice có liên quan mật thiết với nhau. Slice là 1 kiểu dữ liệu nhẹ hơn mà cho phép truy cập các array con trong 1 array

- 1 Slice có 3 thành phần cơ bản: pointer, length, capacity
    + pointer trỏ đến phần tử đầu tiên của array mà có thể truy cập được qua lát cắt
    + length chỉ số phần tử của slice, nó không bh vượt qua được capactity
    + capactity thường là số lượng giữa phần tử đầu tiên của slice cho đến phần tử cuối cùng của array

- Slicing vượt quá cap sẽ gây lỗi `panic` còn slicing quá len sẽ mở rộng slice

```
Q2 := months[4:7]
summer := months[6:9]
fmt.Println(Q2) // ["April" "May" "June"]
fmt.Println(summer) // ["June" "July" "August"]

fmt.Println(summer[:20]) // panic: out of range
endlessSummer := summer[:5] // extend a slice (within capacity)
fmt.Println(endlessSummer) // "[June July August September October]"
```

- Vì 1 slice chứa 1 pointer đến 1 phần tử của array nên việc chuyển 1 slice tới 1 hàm cho phép hàm sửa đổi các giá trị cuả array. Nói cách
khác, copy slice tạo 1 bản sao đến array phía dưới