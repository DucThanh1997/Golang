## Looping

Để biết khi nào cái goroutine cuối cùng kết thúc, chúng ta cần tăng 1 bộ đếm trước khi mỗi quy trình bắt đầu và giảm nó khi mỗi goroutine kết thúc. Điều này đòi hỏi 1 bộ đếm đặc biệt, có thể điều khiển 1 cách an toàn từ nhiều tuyến goroutines và cung cấp 1 cách để đợi cho đến khi nó trở thành 0. Loại bộ đếm này được gọi là sync.WaitGroup


```
// makeThumbnails6 makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames <-chan string) int64 {
    sizes := make(chan int64)
    var wg sync.WaitGroup // number of working goroutines
    for f := range filenames {
        wg.Add(1)
        // worker
        go func(f string) {
            defer wg.Done()
            thumb, err := thumbnail.ImageFile(f)
            if err != nil {
                log.Println(err)
                return
            }
            info, _ := os.Stat(thumb) // OK to ignore error
            sizes <- info.Size()
        }(f)
    }
    // closer
    go func() {
        wg.Wait()
        close(sizes)
    }()
    var total int64
    for size := range sizes {
        total += size
    }
    return total
}
```

- Tạo WaitGroup `wg sync.WaitGroup` 
- Khi chạy 1 cái goroutine mới chúng ta sẽ add 1 giá trị vào `wg.Add(1)`. Bao giờ xong thì `wg.Done()` và chúng ta phải đợi trong wg không còn gì nữa thì mới close `wg.Wait()`