##  Flag
Trong phần này chúng ta sẽ xem làm thế nào 1 standard interface `flag.Value` giúp chúng ta define 1 notation mới cho command-line flags. Hãy xem chương trình dưới

```
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
    flag.Parse()
    fmt.Printf("Sleeping for %v...", *period)
    time.Sleep(*period)
    fmt.Println()
}
```
Ở default, sleep mất 1s nhưng nó có thể control bằng dòng -period. `flag.Duration` tạo 1 biến flag của kiểu `time.Duration` và cho phép người dùng chỉ định duration theo cách rất thân thiện với người dùng

```
$ ./sleep -period 50ms
Sleeping for 50ms...
$ ./sleep -period 2m30s
Sleeping for 2m30s...
$ ./sleep -period 1.5h
Sleeping for 1h30m0s...
$ ./sleep -period "1 day"
invalid value "1 day" for flag -period: time: invalid duration 1 day
```

Vì cái giá trị duration flag rất hữu dụng nên tính năng này được tích hợp vào package `flag`, nhưng rất dễ để tạo 1 flag notations mới cho kiểu dữ liệu của chúng ta. Chúng ta chỉ cần định nghĩa kiểu mà sẽ thỏa mãn interface `flag.Value`

```
package flag

// Value is the interface to the value stored in a flag.
type Value interface {
    String() string
    Set(string) error
}
```
- method string sẽ tạo nên các helper message để show ra cho người dùng
- set method để parse các giá trị được truyền vào


