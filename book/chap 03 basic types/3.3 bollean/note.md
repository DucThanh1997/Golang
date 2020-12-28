# Booleans
- Có 2 giá trị `true` và `false`

- Không đổi thẳng true, false sang 0 với 1 được, phải qua 1 hàm btoi
```
func btoi(b bool) int {
    if b {
        return 1
    }
        return 0
}
```