## Function Declarations
- Khởi tạo hàm có:
    + Tên hàm
    + List các param đầu vào: Danh sách các param chỉ định tên và kiểu đầu vào của function. Là các biến local được người dùng cung cấp
    + List các kết quả trả ra: Chỉ định kiểu dữ liệu mà hàm sẽ trả ra
    + Body: Phần này chứa các hành động mà function sẽ thực hiện và nó phải kết thúc bởi return

- Các cách khai báo nhanh
```
func f(i, j, k int, s, t string) { /* ... */ }
func f(i int, j int, k int, s string, t string) { /* ... */ }

func add(x int, y int) int { return x + y }
func sub(x, y int) (z int) { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int { return 0 }
fmt.Printf("%T\n", add) // "func(int, int) int"
fmt.Printf("%T\n", sub) // "func(int, int) int"
fmt.Printf("%T\n", first) // "func(int, int) int"
fmt.Printf("%T\n", zero) // "func(int, int) int"
```

- 
-