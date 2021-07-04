## Goroutines
- Là 1 cái không gian chứa thread được quản lý bởi go runtime
- 1 goroutine rất nhẹ vởi chỉ 2kb và low CPU
- Channels được dùng để liên lạc giữa các g
- Go runtime tạo OS thread và Go routine chạy trong cái context của OS thread đó
- Nhiều Goroutine có thể chạy trong 1 thread

## Go scheduler
- Go runtime có cơ chế chạy MN Scheduler có nghĩa là
- N goroutines sẽ được chạy trên M OS threads dựa theo cái GOMAXPROCS
- Từ go 1.14 trở đi thì Go scheduler cho mỗi goroutine khoảng 10ms để chạy sau đó nó lại rơi vào hàng chờ cuối

