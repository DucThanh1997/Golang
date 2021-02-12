## Testing function
Function Test phải import package testing và có cái signature như này
```
func TestName(t *testing.T) {
    // ...
}
```

Và nó phải được bắt đầu bởi chữ Test đầu tiên

```
func TestSin(t *testing.T) { /* ... */ }
func TestCos(t *testing.T) { /* ... */ }
func TestLog(t *testing.T) { /* ... */ }
```

param t dùng để cung cấp các method như báo cáo các lỗi và log ra các thông tin.

- Khi chúng ta chạy `go test -v` thì nó sẽ in ra các hàm chạy, hàm nào fail hàm nào pass 
```
=== RUN   TestPalindrome
--- PASS: TestPalindrome (0.00s)
=== RUN   TestFrenchPalindrome
--- FAIL: TestFrenchPalindrome (0.00s)
    word_test.go:19: IsPalindrome("été") = false
=== RUN   TestCanalPalindrome
--- FAIL: TestCanalPalindrome (0.00s)
    word_test.go:26: IsPalindrome("A man, a plan, a canal: Panama") = false
=== RUN   TestNonPalindrome
--- PASS: TestNonPalindrome (0.00s)
FAIL
```

- Chạy lệnh này `go test -v -run="French|Canal"` thì nó sẽ chỉ chạy những function mà match theo kiểu regular expression thôi. Ví dụ ở đây nó sẽ chỉ chạy hàm `TestFrenchPalindrome` và `TestCanalPalindrome`

- Ở các ví dụ đầu tiên này thì chúng ta đang dùng `t.Errorf` nó ko gây panic app và dừng app. Điều này giúp chúng ta tránh được việc test bị hủy nếu có 1 test case xuất hiện lỗi. Và giúp chúng ta chạy được toàn bộ các test case trong 1 lần chạy bất kể các kết quả của chúng ra sao

- Nếu muốn hủy thì dùng lệnh `t.Fatalf` app sẽ dừng ngay lập tức                                                                         

### Random testing
Tạo ra random input để test. Vậy làm thế nào để biết output đó là đúng khi input là random. Có 2 cách

+ 1 là chúng ta viết 1 cái function thay thế mà dùng 1 thuật toán rõ ràng và đơn giản hơn rồi check xem cả 2 cái có cùng ra 1 kết quả ko

+ 2 là chúng ta tạo ra 1 input pattern và random trên nó