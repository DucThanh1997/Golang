package main
import "fmt"
import "time"
 
// Đẩy dữ liệu cho channel và chờ 4 giây
func data1(ch chan string) {  
    // time.Sleep(4 * time.Second)
    ch <- "from data1()"
}
 
// Đẩy dữ liệu cho channel và chờ 2 giây
func data2(ch chan string) {  
    time.Sleep(2 * time.Second)
    ch <- "from data2()"
}
 
func main() {
    // Tạo các biến channel để truyền giá trị string  
    chan1 := make(chan string)
    chan2 := make(chan string)
    
    // Gọi các goroutine cùng với các biến channel
    go data1(chan1)
    go data2(chan2)
 
    // Cả hai câu lệnh case kiểm tra dữ liệu trong chan1 or chan2.
    // Nhưng dữ liệu không sẵn sàng (Cả 2 routines đều tạm dừng 2 giây và 4 giây)
    // Nên đoạn code trong default sẽ được chạy mà không chờ dữ liệu trong các channel.
    select {
    case x := <-chan1:
        fmt.Println(x)
    case y := <-chan2:
        fmt.Println(y)
    default:
    	fmt.Println("Default case executed")
    }
}