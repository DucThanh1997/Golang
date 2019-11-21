Chúng ta có đoạn code trong file `main.go` như này

```
package main

import "fmt"

func main() {
    fmt.Println("Hi there")
}
```
Chúng ta sẽ trả lời lần lượt 5 câu hỏi dưới đây

## 1.Chạy như nào
- Vào terminal gõ `go run main.go`
- `go` kiểu như là 1 cái cli để sử dụng nó có vài câu lệnh hay ho như

- `go run`: để compile code và chạy. lệnh này có thể chạy nhiều file go 1 lúc. Ví dụ `go run main.go func1.go`
- `go build`: để compile code, kết quả thu được nếu không có lỗi sẽ là 1 file cùng tên với file `.go`. Nhấn vào xong rồi chạy thôi
- `go fmt`: để format lại code cho chuẩn và đẹp
- `go get` và `go install`: để lấy source code từ trên mạng về
- `go test`: để test cái file đó 


## 2. Package main có nghĩa là gì
- Package kiểu dạng như là 1 cái workspace có thể chứa 1 hoặc nhiều file go và mỗi file trong đó phải khai báo xem nó thuộc package nào ở dòng đầu tiên

![image](https://user-images.githubusercontent.com/45547213/69328314-9d718b00-0c81-11ea-808a-6d960d5c9b3e.png)

- Thế tại sao phải tên là main? Vì trong go có 2 loại package **excutable package** và **reuseable package**

- excutable package: Dùng để có thể build ra 1 file excutable được.

- reuseable package: Dùng để vứt code vào đấy và gọi ra khi cần

- Cách phân biệt duy nhất giữa 2 loại package này là excutable package phải tên là main `package main` còn nếu đặt tên khác nó sẽ là reuseable package. Và package main phải có hàm main

## 3. import và "fmt" là gì 

- `import` là để import các package khác vào để file hiện tại chúng ta đang code có thể sử dụng được code ở trong package đó

- `fmt` ở đây là 1 package trong standard library của go đã được tải sẵn vào máy. fmt là viết tắt của format

![image](https://user-images.githubusercontent.com/45547213/69329690-225da400-0c84-11ea-97a8-5f40ca14f3dc.png)

- Import có thể import cả package chúng ta tự viết ra chứ không chỉ ở mỗi standard library của go

