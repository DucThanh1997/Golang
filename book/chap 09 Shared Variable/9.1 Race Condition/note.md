## Race condition
- Ở 1 chương trình tuần tự thì chỉ có 1 goroutine. Các bước của chương trình xảy ra theo thứ tự logic. Trong 1 chương trình có nhiều hơn 1 goroutine. Các trình tự trong mỗi goroutine vẫn xảy ra bình thường nhưng chúng ta không biết khi nào event x xảy ra ở 1 goroutine trước khi event y xảy ra ở 1 goroutine khác hoặc xảy ra sau. Khi chuyngs ta không chắc chắn về điều này thì event x và y là 

- 1 function chạy chuẩn trên 1 chương trình bình thường. Nó `concurency-safe` chỉ khi mà nó tiếp tục chạy chuẩn ở concurent. 


- `Race condittion` là 1 tình trạng mà chương trình không đưa ra kết quả đúng cho các hoạt động xen kẽ trên nhiều goroutine. `Race condittion` rất nguy hiểm vì nó xuất hiện ngầm và ít khi show ra, chỉ khi chạy load nặng hoặc chạy 1 loại compilers, platforms, or architectures. Điều này làm chúng khí tại hiện và tìm cách sửa

- Data race xảy ra khi có 2 goroutine cùng access vào 1 variable và có ít nhất 1 cái access vào để việt

```
var x []int
go func() { x = make([]int, 10) }()
go func() { x = make([]int, 1000000) }()
x[999999] = 1 // NOTE: undefined behavior; memory corruption possible!

```

Có thể cái pointer lấy lệnh đầu để make và độ dài lấy từ lệnh thứ 2. x sẽ ảo diệu, khi slice  trên danh nghĩa là 1.000.000 nhưng cái array nằm dưới cái slice đó có 10 phần tử thôi

Có 3 cách để tránh data race:

- Cách 1 là không viết vào biến khi bạn đang đọc hãy viết xong hết rồi mới đọc. Nhưng nếu có trường hợp update thì chịu nhé


- Cách 2 để tránh data race là tránh để biến bị access 1 lúc từ nhiều goroutine. Chúng ta sẽ hạn chế chỉ cho 1 biến vào 1 goroutine và nếu muốn update nó thì phải truyền tín hiệu vào channel để đưa ra tín hiệu thay đổi

```
var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int { return <-balances }

func teller() {
    var balance int // balance is confined to teller goroutine
    for {
        select {
        case amount := <-deposits:
            balance += amount
        case balances <- balance:
        }
    }
}

func init() {
    go teller() // start the monitor goroutine
}
```

Kể cả khi biến được giới hạn chỉ trong 1 goroutine trong suốt cả lifetime. Chúng share biến giữa các goroutine bằng cách pass address vào channel rồi xử lý dần dần

- Cách thứ 3 là cho phép nhiều goroutine access nhưng là dần dần cái này gọi là mutual exclusion