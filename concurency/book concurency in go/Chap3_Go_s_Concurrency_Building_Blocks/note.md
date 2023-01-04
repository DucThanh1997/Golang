## Goroutine
Goroutine là 1 đơn vị cơ bản trong chương trình của go. Trong mỗi chương trình của go sẽ có ít nhất 1 goroutine (main go routine, được khởi tạo và chạy khi process bắt đầu)

Nói 1 cách đơn giản thì goroutine là 1 hàm chạy concurently (đồng thời) với các dòng code khác. Để chạy thì chỉ cần gắn cú pháp `go` trước function
```
func main() {
    go sayHello()
    // continue doing other things
}
func sayHello() {
    fmt.Println("hello")
}
```

Goroutines là unique đối với Go (mặc dù một số ngôn ngữ khác có nguyên hàm đồng thời tương tự). Chúng không phải là các OS thread và chúng không chính xác là các green thread được quản lý bởi language runtime. Chúng là một cấp độ trừu tượng cao hơn được gọi là các coroutine. Coroutine chỉ đơn giản là các chương trình con đồng thời (hàm, bao đóng hoặc phương thức trong Go) không được ưu tiên — nghĩa là chúng không thể bị gián đoạn. Thay vào đó, các coroutine có nhiều điểm xuyên suốt cho phép tạm dừng hoặc truy cập lại.

Điều làm cho goroutines trở nên độc đáo đối với Go là sự tích hợp sâu của chúng với Go runtime. Goroutine không xác định điểm tạm dừng hoặc điểm chạy tiếp sau khi dừng; Go runtime quan sát hành vi của goroutine và tự động tạm dừng và chạy tiếp khi thích hợp. Theo một cách nào đó. Đó là sự hợp tác tao nhã giữa go runtime và logic của goroutine. Do đó, goroutines có thể được coi là một lớp coroutine đặc biệt

Cơ chế lưu trữ goroutines của Go được gọi là M:N scheduler. Có nghĩa là map giữa M green threads tới N OS threads. Goroutine được sắp xếp vào các green threads. Khi chúng ta có nhiều goroutine hơn là green threads. Go sẽ xử lý việc phân phối các goroutine vào các green threads và đảm bảo rằng khi các goroutine bị block thì chúng sẽ chạy các goroutine sẵn sàng khác

## The sync package
sync package chứa các concurency primitive (gốc) rất hữu dụng trong low level memory access synchronization (đồng bộ hóa truy cập bộ nhớ cấp thấp). Nếu bạn đã từng làm việc với 1 ngôn ngữ mà xử lý concurency bằng cách memory access synchronization (đồng bộ hóa truy cập bộ nhớ), thì package sync này sẽ khá là familiar với bạn

### Waitgroup
Waitgroup là 1 cách tuyệt vời để đợi các goroutine mà không cần quan tâm đến kết quả của chúng. Nếu bạn cần đợi goroutine và cũng quan tâm đến goroutine thì bạn nền dùng channel và select

```
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    fmt.Println("1st goroutine sleeping...")
    time.Sleep(1)
}()
wg.Add(1)
go func() {
    defer wg.Done()
    fmt.Println("2nd goroutine sleeping...")
    time.Sleep(2)
}()
wg.Wait()
fmt.Println("All goroutines complete.")
```

### Mutex and RWMutex
Mutex là viết tắt của mutual exclusion và nó là 1 cách chỉ cho phép 1 goroutine truy cập vào 1 biến của chúng ta tại 1 thời điểm

RWMutex được khóa bởi nhiều hơn một tác vụ đọc hoặc duy nhất một tác vụ ghi. Khi RWMutex có giá trị 0 thì khóa mở. Trong trường hợp dùng RWMutex các tác vụ đọc không cần phải đợi nhau vì bản chất đọc không thay đổi dữ liệu. Do đó RWMutex cho phép nhiều tác vụ đọc cùng giữ khóa nhưng chỉ có một tác vụ ghi giữ khóa.

### Cond
một điểm cho goroutines đang chờ đợi hoặc thông báo về sự xuất hiện của một sự kiện. Theo định nghĩa đó, một “sự kiện” là bất kỳ tín hiệu tùy ý nào giữa hai hoặc nhiều goroutine không mang thông tin nào khác ngoài thực tế là nó đã xảy ra. Rất thường xuyên, bạn sẽ muốn đợi một trong những tín hiệu này trước khi tiếp tục thực hiện trên goroutine.

### Once
sync.Once đảm bảo rằng chỉ một lệnh gọi Do gọi hàm được truyền vào — ngay cả trên các goroutine khác nhau

### Pool
Ở cấp độ cao, pool là một cách để tạo và cung cấp một số lượng cố định hoặc nhóm các thứ để sử dụng. Nó thường được sử dụng để hạn chế việc tạo ra những thứ đắt tiền (ví dụ: kết nối cơ sở dữ liệu) để chỉ một số lượng cố định trong số chúng được tạo, nhưng một số lượng hoạt động không xác định vẫn có thể yêu cầu quyền truy cập vào những thứ này. Trong trường hợp Go's sync.Pool, loại dữ liệu này có thể được nhiều goroutine sử dụng một cách an toàn

```
myPool := &sync.Pool{
    New: func() interface{} {
        fmt.Println("Creating new instance.")
        return struct{}{}
    },
}
myPool.Get()
instance := myPool.Get()
myPool.Put(instance)
myPool.Get()
```
Here we call Get on the pool. These calls will invoke the New function defined on the pool since instances haven’t yet been instantiated. Here we put an instance previously retrieved back in the pool. This increases the available number of instances to one. When this call is executed, we will reuse the instance previously allocated and put it back in the pool. The New function will not be invoked.

### Channel
When using channels, you’ll pass a value into a chan variable, and then somewhere else in your program read it off the channel.

```
stringStream := make(chan string)
go func() {
    stringStream <- "Hello channels!"
}()
salutation, ok := <-stringStream
fmt.Printf("(%v): %v", ok, salutation)
```

- Cái `ok` để check xem giá trị truyền vào là từ 1 nơi truyền vào hay là giá trị truyền đến do close channel 

### The select Statement
`The select statement`: Là cách để gắn kết các channel lại với nhau. Nhờ select chúng ta có thể sắp xếp các channel cùng với nhau trong 1 chương trình

```
var c1, c2 <-chan interface{}
var c3 chan<- interface{}
select {
case <- c1:
    // Do something
case <- c2:
    // Do something
case c3<- struct{}{}:
    // Do something
}
```

Với select tất cả các lần reads and writes của channel được xem xét đồng thời để xem liệu có bất kỳ channel nào sẵn sàng hay không Nếu không có channel nào sẵn sàng, thì toàn bộ câu lệnh select sẽ bị chặn. Sau đó, khi một channel đã sẵn sàng, hoạt động đó sẽ tiếp tục và các câu lệnh tương ứng của nó sẽ thực thi. Hãy xem một ví dụ nhanh:

```
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	case a, oke := <-c:
		fmt.Println("a: ", a)
		fmt.Println("oke: ", oke)
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}
```

Vậy điều gì xảy ra khi
### What happens when multiple channels have something to read?: 
```
c1 := make(chan interface{}); close(c1)
c2 := make(chan interface{}); close(c2)
var c1Count, c2Count int
for i := 1000; i >= 0; i-- {
    select {
    case <-c1:
        c1Count++
    case <-c2:
        c2Count++
    }
}
fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
```
Kết quả của C1 là 505 còn C2 là 496
Khi có nhiều channel có dữ liệu truyền vào Go runtime sẽ cố gắng lựa chọn ngẫu nhiên một cách cân bằng nhất, nghĩa là các case sẽ có % được chọn và chạy ngang nhau


### What if there are never any channels that become ready?
Nếu ko có gì xảy ra nữa có lẽ chúng ta sẽ muốn timeout vậy nên chúng ta có thể dùng timeout như sau
```
var c <-chan int
select {
case <-c:
case <-time.After(1 * time.Second):
    fmt.Println("Timed out.")
}
```


## The GOMAXPROCS Lever
Hàm này kiểm soát lượng OS thread mà chúng ta hay gọi là `work queues`
