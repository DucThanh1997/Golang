## buffer
### 1. Định nghĩa
- Là 1 kĩ thuật kinh điển khi đọc và ghi dữ liệu

### 2. Tác dụng
- Lưu tạm thời data nhận được từ một IO Operation trong user space trước khi gửi tới kernel space (trường hợp write) hoặc trước khi chuyển tới process của bạn ( trường hợp read)
- Tối giản hóa số lần syscall

https://kipalog.com/posts/IO--Buffer-vs-non-buffer-technique

### 3. Cách hoạt động
Bình thường mỗi lần đọc chúng ta sẽ đọc 1 kb vậy nếu file đó là 5GB. thì phải đọc 5 triệu lần, và ổ địa sẽ phải quay 5 triệu lần để lấy từng kb ra một

Buffer sẽ lấy 4kb hoặc hơn trong 1 lần lấy, xử lý ví dụ là 1 kb đi, còn 3 kb còn lại đưa vào ram, sau khi xử lý xong 1kb đầu tiên nó lấy lần lượt từng 1kb trong ram ra rồi xử lý tiếp. Lấy hết trong ram thì mới đi tìm và lấy trong ổ đĩa

Như vậy sẽ tiết kiệm số lần tìm trong ổ đĩa hơn

### Nguồn

https://kipalog.com/posts/IO--Buffer-vs-non-buffer-technique