## Multiplexing with select
- Phần mềm đếm rocket launch
```
func main() {
    fmt.Println("Commencing countdown.")
    tick := time.Tick(1 * time.Second)
    for countdown := 10; countdown > 0; countdown-- {
        fmt.Println(countdown)
        <-tick
    }
    launch()
}
```

- Sau đó thêm tính năng hủy
```
abort := make(chan struct{})
go func() {
    os.Stdin.Read(make([]byte, 1)) // read a single byte
    abort <- struct{}{}
}()
```

Bây giờ chúng ta cần 1 cái để multiplex operations, và nó được gọi là câu lệnh select

```
select {
case <-ch1:
    // ...
case x := <-ch2:
    // ...use x...
case ch3 <- y:
    // ...
default:
    // ...
}
```

Nó cũng giống như switch statement cũng có nhiều trường hợp và các case

- Select đợi cho đến khi có 1 trong các case met exception thì nó sẽ chạy. Còn không nó đợi mãi mãi. Bây giờ sửa lại cái function nổ rocket

```
fmt.Println("Commencing countdown. Press return to abort.")
select {
case <-time.After(10 * time.Second):
    // Do nothing.
case <-abort:
    fmt.Println("Launch aborted!")
return
}
launch()
```

```
func main() {
	// ...create abort channel...
	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	fmt.Println("launch()")
}
```

`time.Tick` kiểu như nó sẽ tạo 1 goroutine mà gọi cái time.Sleep trong vòng lặp. gửi 1 event khi mà nó thức dậy. Khi hàm countdown return, nó sẽ dừng nhận event từ tick nhưng cái ticker routine vẫn ở đó và nó không biết gửi đi đâu dẫn đến `goroutine leak`
