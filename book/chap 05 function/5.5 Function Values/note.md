## Function Values
Nó là first class value tức là như mọi value khác nó có kiểu, nó có thể gán vào 1 biến khác hoặc trả về bởi function

```
func square(n int) int { return n * n }
func negative(n int) int { return -n }
func product(m, n int) int { return m * n }

f := square
fmt.Println(f(3)) // "9"

f=negative
fmt.Println(f(3)) // "-3"
fmt.Printf("%T\n", f) // "func(int) int"

f=product // compile error: can't assign f(int, int) int to f(int) int
```

Nhưng nó không so sánh được và không làm key trong map được
