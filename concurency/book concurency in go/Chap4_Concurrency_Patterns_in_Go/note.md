## Confinement
Thường thì khi dùng goroutine chúng ta hay có 2 kiểu
- Synchronization primitives for sharing memory (e.g., sync.Mutex)
- Sharing memory by communicating  (e.g., channels)

Ngoài ra thì còn có thể share bằng cách
- Immutable data: Rất an toàn vì không ai thay đổi được nó nhưng cũng có nhược điểm là nó ko modify được
- Data protected by confinement: Confinement là ý tưởng đơn giản nhưng mạnh mẽ để đảm bảo thông tin chỉ có sẵn từ một concurent process(goroutine). Khi đạt được điều này, a concurrent program hoàn toàn an toàn và không cần đồng bộ hóa. Có hai loại confinement có thể xảy ra: đặc biệt và theo từ vựng.

Ad hoc confinement là khi bạn đạt được confinement thông qua convention cho dù nó được thiết lập bởi cộng đồng ngôn ngữ hay nhóm mà bạn làm việc hay codebase mà bạn làm việc trong đó. Theo ý kiến của tôi, rất khó để tuân theo quy ước đối với các dự án ở mọi quy mô trừ khi bạn có các công cụ để thực hiện phân tích trên code của mình mỗi khi ai đó commits code. Đây là một ví dụ về Ad hoc confinement:

```
data := make([]int, 4)
loopData := func(handleData chan<- int) {
    defer close(handleData)
    for i := range data {
        handleData <- data[i]
    }
}
handleData := make(chan int)
go loopData(handleData)
for num := range handleData {
    fmt.Println(num)
}
```

Chúng ta có thể thấy rằng data slice of integers available ở cả hàm loopData và vòng lặp trên kênh handleData; tuy nhiên, theo quy ước, chúng ta chỉ truy cập nó từ hàm loopData. Nhưng khi nhiều người code và deadline sắp đến, lỗi có thể xảy ra và confiment có thể bị break. Vậy nên thường prefer lexical hơn


```
chanOwner := func() <-chan int {
    results := make(chan int, 5) // SỐ 1
    go func() {
        defer close(results)
        for i := 0; i <= 5; i++ {
            results <- i
        }
    }()
    return results
}
consumer := func(results <-chan int) {
    for result := range results {
        fmt.Printf("Received: %d\n", result)
    }
    fmt.Println("Done receiving!")
}
results := chanOwner()
consumer(results)
```

Ở chỗ số 1 chúng ta tạo ra 1 result nhưng hàm chanOwner chỉ trả ra 1 chan chỉ nhận dữ liệu. Điều này ngăn chặn việc các hàm hoặc đoạn code ở đằng sau nó viết vào cái chan trả ra đây (gây sai lệch)

## The for-select Loop
Chúng ta sẽ sử dụng nó ở 1 số case như
```
for _, s := range []string{"a", "b", "c"} {
    select {
    case <-done:
        return
    case stringStream <- s:
    }
}
```

hoặc Looping innitely waiting to be stopped
```
for {
    select {
    case <-done:
        return
    default:
    }
 // Do non-preemptable work
}
```

## Preventing Goroutine Leaks
they do cost resources, and goroutines are not garbage collected by the runtime, so regardless of how small their memory footprint is, we don’t want to leave them lying
about our process

Goroutine chạy sẽ cần tài nguyên và goroutine không được `garbage collected` bởi go runtime, vậy nên dù chúng có tốn ít tài nguyên như thế nào, chúng ta cũng ko muốn chúng tồn tại trong chương trình khi đã không cần đến nữa. Vậy khi nào thì những tài nguyên của goroutine cần được loại bỏ

- Khi goroutine đã chạy xong
- Khi chúng không thể tiếp tục được nữa vì một vấn đề gì đó xảy ra
- Khi chúng được ra lệnh là phải dừng lại

2 cái đầu thì có thể tự xử lý đơn giản bởi go. Cái sau thì thường truyền 1 done channel vào, bh done thì main goroutine đóng cái close lại để thằng goroutine con biết đường mà return nhằm tránh cái việc tài nguyên ko bị dọn dẹp 
```
func main() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} { // số 1
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// Do something interesting
					fmt.Println(s)
				case <-done: // số 2
					return
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{})
	terminated := doWork(done, nil)
	go func() { // số 3
		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()
	<-terminated // số 4
	fmt.Println("Done.")
}
```

## Or channel pattern
Đôi khi, bạn muốn kết hợp một hoặc nhiều done channel thành một done channel duy nhất và done channel này sẽ đóng nếu bất kỳ channel thành phần nào của nó đóng. Hoàn toàn có thể chấp nhận được, mặc dù dài dòng, để viết một câu lệnh chọn thực hiện khớp nối này; tuy nhiên, đôi khi bạn không thể biết số lượng done channel mà bạn đang làm việc trong thời gian chạy. Trong trường hợp này, bạn có thể kết hợp các channel này với nhau bằng cách sử dụng the or-channel pattern.

```
var or func(channels ...<-chan interface{}) <-chan interface{}
or = func(channels ...<-chan interface{}) <-chan interface{} { // số 1
    switch len(channels) {
    case 0: // số 2
        return nil
    case 1: // số 3
        return channels[0]
    }
    orDone := make(chan interface{})
    go func() { // số 4
        defer close(orDone)
        switch len(channels) {
        case 2: // số 5
            select {
                case <-channels[0]:
                case <-channels[1]:
            }
            default: // số 6
                select {
                case <-channels[0]:
                case <-channels[1]:
                case <-channels[2]:
                case <-or(append(channels[3:], orDone)...): // số 6
            }
        }
    }()
    return orDone
```

## Fan-out, Fan-in
Fan-out là thuật ngữ mô tả quá trình bắt đầu nhiều goroutine để xử lý input từ pipeline và fan-in là thuật ngữ mô tả quá trình kết hợp nhiều kết quả vào một channel