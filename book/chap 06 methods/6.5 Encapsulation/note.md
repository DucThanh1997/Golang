## Encapsulation
- 1 biến hoặc 1 method của 1 object được gọi là đóng gói (encapsulation) nếu nó không access vào được từ client. Encápulation, thi thoảng còn được biết đến để giấu dữ liệu (information hiding), là 1 mảng quan trọng trong lập trình hướng đối tượng

- Encapsulation có 3 lơi ích
    + Đầu tiên, vì client không thể trực tiếp modify biến object, thứ mà cần tốn ít công sức để dễ hiểu
    + Thứ hai, giấu đi các chi tiết implementation ngăn chặn client khỏi việc phụ thuộc vào những thứu có thể thay đổi, điều này giúp cho chúng ta có nhiều triên khai implementation mà không cần phá vỡ cấu trúc của API
    + Thứ 3, nó ngăn chặn client khỏi việc setting các biến object linh tinh. Ví dụ ở cái counter dưới đây, chúng ta chỉ cho phép người dùng tăng hoặc đưa cái biến n về không chứ không cho phép set biến n

    ```
    type Counter struct { n int }
    func (c *Counter) N() int { return c.n }
    func (c *Counter) Increment() { c.n++ }
    func (c *Counter) Reset() { c.n = 0 }
    ```
    