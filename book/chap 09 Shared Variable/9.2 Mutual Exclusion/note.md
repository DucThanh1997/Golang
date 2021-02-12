## Mutual 
```
import "sync"
var (
    mu sync.Mutex // guards balance
    balance int
)

func Deposit(amount int) {
    mu.Lock()
    balance = balance + amount
    mu.Unlock()
}

func Balance() int {
    mu.Lock()
    b := balance
    mu.Unlock()
    return b
}
```

Mỗi khi mà goroutine access 1 biến của bank( balance). Nó phải gọi `mutex's Lock` để tạo 1 lock. Nếu có 1 goroutine khác đã dùng cái lock này vào biến đó rồi thì goroutine hiện tại sẽ bị block lại cho đến khi goroutine kia gọi unlock và hàm lock lại available. Hàm guard cho phép chia sẻ biến

Chúng ta nên nhét mu.Unlock vào defer