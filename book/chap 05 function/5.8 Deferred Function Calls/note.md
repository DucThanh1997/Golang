## Defer
- Về mặt cú pháp, defer là 1 hàm thông thường hoặc là 1 method được đặt defer. Khi gọi nó thì hàm đi cùng nó được thực thi nhưng lệnh gọi bị hoãn lại cho đến khi function chứa lệnh defer đã hoàn thành cho dù function đó kết thúc theo kiểu return bình thường hoặc dính lỗi đến mức panic

- Defer thường được dùng trong các hoạt động mở đóng, connect hoặc disconnect, lock hoặc unlock để đảm bảo resource được released đi trong mọi trường hợp

```
func bigSlowOperation() {
    defer trace("bigSlowOperation")() // don't forget the extra parentheses
    // ...lots of work...
    time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
    start := time.Now()
    log.Printf("enter %s", msg)
    return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
```
Function Trace để đo thời gian

Lưu ý nè
```
for _, filename := range filenames {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close() // NOTE: risky; could run out of file descriptors
    // ...process f...
}
```
Defer sẽ không được chạy cho đến khi nó chạy hết vòng for, để khắc phục

```
for _, filename := range filenames {
    if err := doFile(filename); err != nil {
        return err
    }
}
func doFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()
// ...process f...
}
```

Như thế này thì sau khi mà nó hoàn thành cái doFile nó chạy defer luôn