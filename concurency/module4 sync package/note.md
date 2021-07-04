- sync.Muxtex: Dùng để bảo vệ các tài nguyên được chia sẻ. Ví dụ

Nếu cả 2 thằng routine cùng 1 lúc access vào thì nó sẽ dẫn đến lỗi nên thằng sync ra đời để tại 1 thời điểm chỉ có 1 thằng truy cập được vào cái biến đó thôi

- sync.Cond: Ví dụ như này chúng ta có 1 cái go routine đợi 1 cái condition nào đó rồi mới chạy, và cái condition được gỡ 1 một routine khác vậy làm sao để 2 cái giao tiếp được với nhau 1 cách hiệu quả. Có thể dùng channel nhưng nếu có nhiều goroutine thì sẽ hết sức phức tạp nên thằng sys.Cond ra đời

    + c.Wait() sẽ bảo cái goroutine đó dừng lại và đợi
    + c.Signal() sẽ bảo cái goroutine đợi lâu nhất là bạn ơi bạn chạy được rồi
    + c.Broadcast sẽ bảo cho tất cả các thằng goroutine đang đợi là okke hết rồi chúng mày chạy đi

- Go hỗ trợ 1 phương án test race detector: `go run -race mysrc.go`. Nhưng nó sẽ ngốn gấp 10 lần cả về tốc độ lẫn bộ nhớ
