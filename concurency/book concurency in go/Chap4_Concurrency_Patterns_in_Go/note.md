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

