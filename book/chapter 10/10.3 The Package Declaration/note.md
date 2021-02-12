## The Package Declaration
Khởi tạo package luôn được yêu cầu ở phần đầu tiên của file go. Mục tiêu chính của nó là xác định package mà file đó thuộc về

Ví dụ mỗi file của package `math/rand` đều start với dòng `package rand` để khi mà bạn import cái package này bạn có thể access các phần tử trong nó như là `rand.Int, rand.Float64`

```
package main
import (
    "fmt"
    "math/rand"
)
func main() {
    fmt.Println(rand.Int())
}
```
Theo convention, tên package sẽ là mảnh cuối của import path, nên có thể tên package có thể trùng nhau dù import path về mặt tổng thể là khác nhau. Ví dụ `math/rand` và `crypto/rand`

- Có 3 rule cho việc đặt tên package này
    + package mà định nghĩa lệnh chạy phải luôn có tên là main
    + đối với các file mà có đuổi `_test` ở phần tên package nếu tên file có đuôi `_test.go`. Khi đó thư mục có thể định nghĩa 2 package: 1 cái thường và 1 cái gọi là `external test package`. `go test` sẽ phát hiện các file có đuôi `_test`. Các `external test package` dùng để tránh dính import cycle
    + vài package có đuổi version ở sau nhưng khi gọi bạn có thể bỏ phần version đó đi. Ví dụ: `gopkg.in/yaml.v2` thì chỉ cần gọi `yaml` là đủ
