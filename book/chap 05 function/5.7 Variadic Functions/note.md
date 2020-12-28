## Variadic Functions
- Bạn có thể truyền vào bao nhiêu tham số tùy muốn

```
func sum(vals ...int) int {
    total := 0
    for _, val := range vals {
        total += val
    }
    return total
}

fmt.Println(sum()) // "0"
fmt.Println(sum(3)) // "3"
fmt.Println(sum(1, 2, 3, 4)) // "10"

```

- Thằng caller tạo 1 array rồi copy các argument vào nó và truyền slice của cả cái array đó vào function

- Mặc dù thằng varadic function nó truyền slice vào nhưng nó vẫn khác cái function có đầu vào là slice
```
func f(...int) {}
func g([]int) {}
```
