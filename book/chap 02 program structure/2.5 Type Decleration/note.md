## Type Decleration
```
type Celsius float64

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC Celsius = 0
    BoilingC Celsius = 100
)
```
Tạo 1 kiểu giống với float64 nhưng mà có thể được gắn với các function mà mình tự tạo

- Với mọi kiểu dữ liệu luôn có hàm để biến đổi các giá trị về kiểu của nó 
    + Ví dụ: kiểu T sẽ có `T(x)` đê biến đổi x về kiểu của nó. Tuy nhiên hàm này chỉ thực hiện được khi mà biến x có kiểu dữ liệu giống với dữ liệu gốc của kiểu T

- Các kiểu dữ liệu có thể so sánh với kiểu dữ liệu gốc của nó
```
var c Celsius
var f Fahrenheit
fmt.Println(c == 0) // "true"
fmt.Println(f >= 0) // "true"
fmt.Println(c == f) // compile error: type mismatch
fmt.Println(c == Celsius(f)) // "true"!
```
