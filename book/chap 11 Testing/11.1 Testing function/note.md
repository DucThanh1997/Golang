## Go test tool
- Go test sẽ tìm những file có đuôi là `_test.go` để chạy. Những file này không thuộc package chính, khi build nó ko bị dính vào

- Với file có đuổi là `_test.go` có 3 kiểu function được chạy
    + test: bắt đầu bởi chữ `Test`, function này thể hiện 1 số logic của hàm để kiểm tra tính đúng đắn của hàm. Go test gọi các hàm này rồi báo cáo kết quả (PASS hoặc FAIL)
    + benchmarks: bắt đầu bằng chữ `Benchmark` để đo hiệu năng của hàm. Go test sẽ báo cáo thời gian chạy trung bình
    + example function: bắt đầu bằng `Example` cung cấp doc của máy tính toán rồi đưa cho
## Testing function
Function Test phải import package `testing` và có cái signature như này
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

- Cái `t.Errof` dùng giống `fmt.Printf`

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

- Concept của error message sẽ thường là `f(x) = y, want z` chạy hàm f với biến đầu vào là x nhưng có kết quả đầu ra là y nhưng thực ra phải là z                     

### Random testing
Tạo ra random input để test. Vậy làm thế nào để biết output đó là đúng khi input là random. Có 2 cách

+ 1 là chúng ta viết 1 cái function thay thế mà dùng 1 thuật toán rõ ràng và đơn giản hơn rồi check xem cả 2 cái có cùng ra 1 kết quả ko

+ 2 là chúng ta tạo ra 1 input pattern và random trên nó

- Vì đầu vào của random test là không xác định, sẽ rất quan trọng khi điền cái đầu vào bị lỗi ra nếu ko thì bạn sẽ đ reproduce lại được

Go test tool test được cả command nhé

### White-Box Testing
Một cách để phân loại test là theo mức độ đầu vào của thông tin mà chúng yêu cầu để chạy được hoạt động bên trong của package được kiểm tra

- `black-box` giả sử là nó đ biết gì về cái package ngoài những gì được show ra ở api và trong doc (tương đối mờ mịt)

- Ngược lại white-box có quyền truy cập các hàm bên trong và các cấu trúc dữ liệu của package, có thể quan sát và thay đổi. Ví dụ white-box test có thể kiểm tra độ bất biến của kiểu dữ liệu sau mỗi hành động

2 phương thức này bổ sung cho nhau. Black box dễ thực hiện, dễ thay đổi còn white-box sẽ cung cấp những thứ chi tiết hơn

### External Test Packages
Ví dụ trong package `net/url` cung cấp khả năng parse url còn `net/http` cung cấp khả năng làm việc với web server. thằng `net/http` depend trên thằng `net/url`. Tuy nhiên nếu 1 trong các test của `net/url` muốn thử sự tương tác của `net/url` và `net/http` thì nó sẽ tạo ra 1 cái import cycle

Để loại bỏ cái này, chúng ta tạo 1 `extenal test package` mà nó sẽ ở cùng chỗ với `net/url` và có package name `url_test`. Vì các test ở 1 package khác nên nó có thể import các helper package khác mà ko sợ import cycle

## Benchmark
- Cũng lấy từ testing package ra. Hàm có tên bắt đầu bằng `benchmark`. Nó dùng `testing.B` với biến N expose ra ngoài chỉ số lần chạy

Hàm để chạy `go test -bench=.`. `-bench` để chọn hàm benchmark nào sẽ được chạy. Nếu bạn điền `.` nó sẽ chạy hết

## Proofling
Khi mà chúng ta muốn xem speed phần mềm của chúng ta. SỬ DỤNG PROFFLING

- PROFFLING là một cách tiếp cận tự động để đo lường hiệu suất dựa trên việc lấy mẫu một số sự kiện trong quá trình thực thi, sau đó in ra chúng trong một bản tóm tắt thống kê kết quả được gọi là một hồ sơ.

Go hỗ trợ nhiều loại PROFFLING, mỗi loại liên quan đến một khía cạnh khác nhau của hiệu suất, nhưng tất cả chúng đều liên quan đến việc ghi lại một chuỗi các sự kiện quan tâm, mỗi loại đều có stack trace đi kèm. Công cụ kiểm tra go đã được tích hợp sẵn hỗ trợ cho một số loại cấu hình.

- PROFFLING CPU xác định các chức năng mà việc thực thi đòi hỏi nhiều thời gian CPU nhất. Luồng đang chạy hiện tại trên mỗi CPU được hệ thống vận hành thực hiện trong vài mili giây một lần, với mỗi lần gián đoạn ghi lại một sự kiện hồ sơ trước khi tiếp tục thực thi bình thường.

PROFFLING heap xác định các câu lệnh chịu trách nhiệm cấp phát bộ nhớ trong hầu hết các bộ nhớ. Các
Các mẫu thư viện cấu hình gọi đến các quy trình cấp phát bộ nhớ trong để trung bình, một sự kiện cấu hình được ghi lại trên 512KB bộ nhớ được cấp phát.

PROFFLING blocking xác định các tác nhân chịu trách nhiệm chặn các goroutines lâu nhất, chẳng hạn như các cuộc gọi hệ thống, kênh gửi và nhận và các khóa chuyển đổi

```
$ go test -cpuprofile=cpu.out
$ go test -blockprofile=block.out
$ go test -memprofile=mem.out
```


