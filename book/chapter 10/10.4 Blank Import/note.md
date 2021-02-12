## Blank Import
Nếu chúng ta không dùng các package mà chúng ta đã import, thì go sẽ báo lỗi. Tuy nhiên trong trường hợp chúng ta chỉ import các package đó vào để lấy các chức năng phụ như, khởi tạo các biến toàn cục của package hoặc chạy hàm init của package. Để pass cái lỗi `unused import` chúng ta sẽ import theo kiểu như này

```
package main

import (
	_ "fmt"
)
func main() {
}
```

Nó được gọi là blank import