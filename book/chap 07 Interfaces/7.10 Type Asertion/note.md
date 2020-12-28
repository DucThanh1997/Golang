## Type Assertions
- Là kiểu `x.(T)`, x là kiểu hiện tại còn T là kiểu mà muốn ép vào. Type assertion check xem kiểu dynamic có match được với kiểu assert vào không

- Có 2 khả năng: 
    + Đầu tiên là kiểu T assert vào sẽ là kiểu cố định, sau đó `type assertion` sẽ kiểm tra xem cái dynamic type đó có giống với T không. Nếu thành công, kết quả của `type assertion` là giá trị của x nhưng mang kiểu T. Nếu failed thì panic xảy ra

    + Nếu kiểu T là interface, nó sẽ check xem có thỏa mãn T không nếu thành công nó sẽ vẫn giữ nguyên cái kiểu interface và value interface đó nhưng các method của interface T vẫn sử dụng được, nói cách khác là gộp 2 cái lại với nhau

- Không cần biết là loại nào được asserted, nếu x là nil interface thì `type assertion` sẽ failed. Nếu bạn muốn test xem `type assertion` có thành công hay không chúng ta sẽ làm như này

```
var w io.Writer = os.Stdout
f, ok := w.(*os.File) // success: ok, f == os.Stdout
b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil
```

nếu thành công thì là true còn không thành công thì là failed
