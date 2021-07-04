## Channel
- Dùng để giao tiếp giữa các channel
- Synchronise giữa các goroutine
- Dữ liệu đi vào cần có 1 kiểu rõ ràng
- Thread-safe

- `ch := make(chan T)` tạo 1 channel với kiểu dữ liệu T
- `ch <- v`: gửi v vào channel
- `v = <- ch`: nhận v vào channel