## 2.1 Names
1. Quy tắc

- Bắt đầu bằng chữ
- Có phân biệt hoa và thường
- Không được đặt trùng với các key word kiểu như `if, for, return, var, type`

2. Scope
- Nếu 1 thực thể được đặt ở trong 1 function thì nó là `local` với function đó, nếu đặt ở ngoài function thì nó sẽ được visible ở cả package
- Nếu 1 thực thể được đặt với chữ cái viết hoa ở đầu thì các package khác cũng có thể được truy cập nó và ngược lại
- Scope là vùng hoạt động của thực thể còn vòng đời là thời gian tồn tại của thực thể đó
- Khi gọi 1 biến nó sẽ tìm ở local rồi ra ngoài tìm

## 2.2 Declarations
- Có 4 kiểu: `var, const, type, func`
- Khai báo ngắn gọn `x := 0` kiểu này hay được dùng khi khai báo biến local
- `:=` là khai báo cả giá trị cả kiểu còn `=` khai báo kiểu thôi 

## 2.3 pointer
- Là địa chỉ của biến
- Với pointer chúng ta có thể đọc và sửa giá trị của biến mà ko cần biết tên biến
