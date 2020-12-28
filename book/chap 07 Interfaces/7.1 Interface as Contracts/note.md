## Interface as contract
- Interface là 1 kiểu trừu tưởng. Nó chỉ chứa 1 vài methods. Khi bạn có 1 giá trị của kiểu interface, bạn không biết nó là gì, bạn chỉ biết nó làm gì

- String formating trong go không yêu cầu các kiểu đầu vào mà thay vào đó là 1 interface
```
package fmt
func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)

func Printf(format string, args ...interface{}) (int, error) {
    return Fprintf(os.Stdout, format, args...)
}

func Sprintf(format string, args ...interface{}) string {
    var buf bytes.Buffer
    Fprintf(&buf, format, args...)
    return buf.String()
}

```
- F đầu tiên trong Fprintf đại diện cho file chỉ ra rằng output đầu ra nên được viết vào file 

- Ở trường hợp của Sprintf, argument không phải là 1 file mà là 1 cái &buf được trỏ đến 1 vùng nhớ nơi byte có thể viết vào

- Param đầu tiên Fprintf cũng không phải là file, nó là `io.Writer` đó là 1 kiểu interface
```
package io
// Writer is the interface that wraps the basic Write method.
type Writer interface {
// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
//
// Implementations must not retain p.
Write(p []byte) (n int, err error)
}
``
- `io.Writer` interface định nghĩa mối ràng buộc giữa `Fprintf` và cái gọi chúng. 

- Ở 1 mặt, mối ràng buộc này yêu cầu khi gọi nó phải cung cấp 1 kiểu cố định như `*os.File or *bytes.Buffer` mà có method Write với đầu vào đầu ra phù hợp 

- Ở mặt khác mối ràng buộc đảm bảo rằng Fprintf sẽ thực hiện công việc của nó với bất kì đầu vào nào thỏa mãn yêu cầu của io.Writer

- Vì `fmt.Fprintf` không giả định gì ở dữ liệu đầu ra nên chúng ta chỉ có thể phụ thuộc vào hành vi được đảm bảo bởi `io.Writter`. Chúng ta có thể truyền vào bất kì kiểu nào thỏa mãn interface `io.Writter` như là tham số đầu tiên của `fmt.Fprintf`