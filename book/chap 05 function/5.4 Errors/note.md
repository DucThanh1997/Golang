## Error
- Là 1 kiểu interface khi chạy hàm, nó được tạo ra để xử lý khi hàm xảy ra lỗi
Có 4 chiến thuật để xử lý `error`

1. Miêu tả rõ ràng cái lỗi ra ở hàm nào, bị làm sao, 

2. Với những lỗi tạm thời hoặc các vấn đề không kiểm soát được, sẽ okke hơn nếu bạn thử lại thêm lẫn nữa

3. Nếu không thể chối bỏ cái lỗi đấy hãy in ra lỗi và dừng program lại 1 cách nhẹ nhàng
```
if err := WaitForServer(url); err != nil {
    fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
    os.Exit(1)
}
```
- Nếu bạn dùng log.Fatal sẽ okke hơn vì nó có thêm ngày giờ nữa

4. Đôi khi bạn chỉ cần in ra thôi