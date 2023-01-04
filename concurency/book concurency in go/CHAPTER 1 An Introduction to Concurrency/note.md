## Race Conditions
Race Conditions xảy ra khi 2 hoặc nhiều hoạt động phải được thực hiện đúng thứ tự nhưng thực chất khi chạy thì lại không đúng thứ tự

```
var data int
go func() {
    data++
}()

if data == 0 {
    fmt.Printf("the value is %v.\n", data)
}
```

Ở đây dòng `data++` và dòng `if data==0` đều cố gắng truy cập biến data nhưng ko có gì đảm bảo được là cta sẽ chạy theo thứ tự nào. Có 3 trường hợp có thể xảy ra ở đây
- Nothing is printed. In this case, line 3 was executed before line 5.
- “the value is 0” is printed. In this case, lines 5 and 6 were executed before line 3.
- “the value is 1” is printed. In this case, line 5 was executed before line 3, but line 3
was executed before line 6.

## Atomic
Trong 1 bối cảnh, nếu 1 operation (hành động) là atomic thì chứng tỏ nó sẽ xảy ra mặc kệ bối cảnh ra sao nó cũng không bị ảnh hưởng

## Deadlock
Khi mọi goroutine đều đang chờ đợi nhau và đều không có action gì 

## Livelocks
Livelocks là 1 chương trình mà đang chạy concurentn, nhưng các concurent này ngăn cản tiến trình của program

Have you ever been in a hallway walking toward another person? She moves to one side to let you pass, but you’ve just done the same. So you move to the other side, but she’s also done the same. Imagine this going on forever, and you understand livelocks.

## Starvation
Starvation là tình huống mà 1 goroutine ko có đủ resource để chạy vì 1 hoặc nhiều goroutine khác đã lấy mất hết. Có 1 cách để detect và giải quyết starvation là logging khi công việc đã hoàn thành rồi xác định xem tần xuất chạy của hàm có đúng như mức bạn mong muốn

## Determining Concurrency Safety
