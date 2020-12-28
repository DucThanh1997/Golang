## Channels 
- Nếu `Goroutine` là các hành động đồng thời của chương trình go, thì channel là các connection giữa chúng. Channel là 1 cơ chế liên lạc mà cho phép 1 `Goroutine` gửi 1 giá trị đến 1 `Goroutine` khác. Mỗi 1 channel như 1 đường ống của 1 giá trị thuộc về 1 kiểu nhất định

- 2 channel có đầu vào cùng 1 kiểu có thể so sánh `==` với nhau.

- Channel có 2 cơ chế chính gửi và nhận. Gửi đi sẽ truyền 1 giá trị vào `Goroutine` qua channel tới 1 `Goroutine` khác đang hứng cái channel đó

- Cả 2 cơ chế đều dùng dấu `<-`

```
ch <- x // a send statement
x=<-ch // a receive expression in an assignment statement
<-ch // a receive statement; result is discarded
```

- Channel còn có 1 cơ chế nữa là `Close` để thông báo rằng sẽ không còn thằng nào được truyền vào cái channel này nữa và sẽ đóng nó lại `close(ch)`

- Cách để tạo 1 channel bằng cách dùng make sẽ tạo `unbuffered` channel, nhưng make cho phép thêm 1 tham số nữa là sức chứa của channel. Nếu sức chứa khác 0 thì make sẽ tạo ra 1 bufferd channel

```
ch = make(chan int) // unbuffered channel
ch = make(chan int, 0) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity 3
```

### Unbuffered Channels
- Send operation trong Unbuffered Channels sẽ block cái `Goroutine` gửi cho đến khi 1 cái `Goroutine` khác thực hiện việc nhận ở channel mà `Goroutine` kia gửi đi. Ngược lại, nếu receive operation được thực hiện trước, `Goroutine` thực hiện receive operation sẽ bị block cho đến khi 1 goroutine khác thực hiện việc send operation

- Giao tiếp qua Unbuffered Channels khiến các `Goroutine` gửi và nhận phải đồng bộ hóa. Do đó, các Unbuffered Channels đôi khi được gọi là synchronous channels. 

### Pipelines
- Channels dùng để connect các goroutine. Cho nên đầu ra của cái này là đầu vào của cái kia. Cái này gọi là pipeline. Ví dụ là cái pipeline dưới đây

![image](https://user-images.githubusercontent.com/45547213/103173692-33b5ae80-488f-11eb-9ff6-79b35de724fc.png)

Nếu người gửi không muốn gửi bất kì thứ gì vào channel nữa thì chúng ta sẽ đóng lại cái channel bằng hàm
```
close(naturals) // natural là tên channel
```

Sau khi mà cái channel đã đóng bất kì cái gì gửi vào channel sẽ dẫn đến panic. Sau khi channel đóng các `goroutine` mà hứng dữ liệu từ channel đó sẽ chạy mà không bị chặn nhưng sẽ trả về zero value

Không có cách nào để kiểm tra trực tiếp xem channel đó đã đóng hay chưa nhưng có 1 cách để thử kiểu như này
```
// Squarer
go func() {
    for {
        x, ok := <-naturals
        if !ok {
            break // channel was closed and drained
        }
        squares <- x * x
    }
    close(squares)
}()
```

### Unidirectional Channel Types
Chúng ta sẽ tách main thành 3 function 
```
func counter(out chan int)
func squarer(out, in chan int)
func printer(in chan int)
```

squarer funtrion đứng ở giữa nhận 2 param input channel và output channel. Cả 2 đều có cùng kiểu nhưng cách dùng của nó là hoàn toàn khác nhau. `in` chỉ dùng để nhận còn `out` chỉ dùng để gửi đi. Tên của nó là như vậy nhưng chả có gì nếu chúng ta muốn gửi đi bằng `in` và nhận về bằng `out`

Để tránh cái mis như này go cung cấp  `unidirectional channel type` mà chỉ cho phép 1 hành động được gửi đi. Kiểu `chan<- int` là chỉ cho phép send vào kiểu int và không cho phép nhận. Ngược lại với `<-chan int`

```
func counter(out chan<- int)
func squarer(out chan<- int, in <-chan int)
func printer(in <-chan int)
```

### Buffered Channels
- Buffered channel có thêm thành phần queue nữa. Cái phía dưới tạo ra 1 buffered channel

```
ch = make(chan string, 3)
```

![image](https://user-images.githubusercontent.com/45547213/103175859-cc542a80-489f-11eb-9590-afae20464af9.png)

- Khi mà channel full thì send operation sẽ bị block ngược lại nếu thằng channel rỗng thì thằng receiver sẽ bị block

- Chúng ta có thể xem sức chứa của channel bằng
```
fmt.Println(cap(ch)) // "3"
```
- Ngoài ra chúng ta có xem độ dài hiện tại của channel bằng bao nhiêu
```
fmt.Println(len(ch)) // "2"
```

Chúng ta sẽ gửi 1 cái gửi đi 3 request và chỉ in ra 1 cái request nhanh nhất
```
func mirroredQuery() string {
    responses := make(chan string, 3)
    go func() { responses <- request("asia.gopl.io") }()
    go func() { responses <- request("europe.gopl.io") }()
    go func() { responses <- request("americas.gopl.io") }()
    return <-responses // return the quickest response
}

func request(hostname string) (response string) { /* ... */ }
```

Nếu chúng ta sử dụng unbuffered channel, 2 cái `Goroutine` thấp hơn sẽ không có chỗ để gửi response đi nên bị stuck vì không có cái goRoutine nào sẽ receive. Cái này gọi là `goroutine leak`. Không như các biến rác, `leaked goroutine` không được tự động thu thập nên hãy để ý khi nào goroutine được terminate

Việc chọn unbuffered hay buffered channel có thể quyết định tính đúng đắn của phần mềm. 

- Unbuffered Channels có tính đồng bộ hơn vì mỗi send operation sẽ được đồng bộ với cái receiver

- Với buffered Channels các hành động được tách rời


