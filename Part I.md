# Viết chương trình đầu tiên
Mở đầu 1 file thường sẽ là `package clause` và việc chỉ rõ tên
của package để cho go biết file bạn đang code thuộc package nào
và 1 file chỉ được thuộc 1 package

Sau đó sẽ là hàm, hàm trong go có thể tái sử dụng và chạy được.
Trong hàm sẽ chưa 1 đống code

Hàm main (`func main`) là 1 hàm đặc biệt để chỉ xem chỗ nào Go sẽ
thực thi

```
package main

import "fmt"
func main() {
  fmt.Println("Hello")
}
```

fmt là 1 thư viên có những function cơ bản. Ở trên bạn gọi ra function `println`

```
Để chạy code ta gõ `go run main.go`


## Sự khác biệt nhỏ nhỏ giữa print, printf, println
```
Printf --> "Print Formatter" Hàm này cho phép ta gán chữ hoặc số và dãy

VD: `fmt.Printf("%s is cool", "Bob") `

Println --> "Print Line" In ra bình thường và có cái `/n` ở cuối

Print --> "Print" giống `println` nhưng không có `/n ơ dưới

## Package
- Đầu tiên tất cả các file trong package phải ở trong cùng 1 thư mục trên hệ điều hành
- Tiếp theo tuy không bắt buộc nhưng 1 thư mục chỉ nên có 1 package
- Tất cả mọi chương trình Go đều phải được bắt đầu bởi câu lệnh khai báo gói. Giúp cho việc tái sử dụng


## Import
- Yêu cầu sử dụng các gói khác từ trong chương trình gói `fmt` ở đây để đọc dữ liệu ra vào 

## Func
- Là thành phần xây dựng nên 1 chương trình. Xử lí dữ liệu, đọc dữ liệu, xuất dữ liệu.
Hàm main là hàm đặc biệt vì nó sẽ được excute khi chạy

# Các kiểu dữ liệu
## Số nguyên
Số nguyên trong go được biểu diễn bởi
![image](https://user-images.githubusercontent.com/45547213/58007239-10a01600-7b14-11e9-92cf-a0ae7485ec8d.png)
- Cứ cái nào có u thì cái đấy ko lưu được số âm và số 8,16,32 để chỉ số bit cần để lữu trữ những số được khai báo

## Số thực 
- Để dùng số thực chúng ta dùng `float64` là đủ

## String
- String (chuỗi) là các kí tự được bọc trong cặp dấu nháy kép hoặc nháy đơn được dùng để biểu diễn văn bản. 
String nằm trong dấu nháy kép có thể sử dụng các kí tự điều khiển đặc biệt như `\n` là xuống dòng, `\t` là dấu tab.
```
package main
 
import "fmt"
 
func main() {
    fmt.Println(len("Hello World")) # Chúng ta có thể in ra độ dài của chuỗi
    fmt.Println("Hello World"[1]) # In ra 1 giá trị trong chuỗi
    fmt.Println("Hello " + "World") # Cộng 2 chuỗi
} 
```

## Boolean
Có 2 giá trị True và False
- Các phép toán go có `or`,`and`,`not`
```
package main
 
import "fmt"
 
func main() {
    fmt.Println(true && false)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)
}
```
