## Custom type decleration

Cách tạo 1 type
```
type deck []string
```

Cạch tạo 1 function receiver
```
func (d deck) print() {
    for i, card := range d {
        fmt.Println("card: ", card)
    }
}
```
- Tất cả mọi biến có kiểu `deck` đều có thể sử dụng function `print`

- `(d deck)` ở đây với `deck` để chỉ kiểu dữ liệu funtion này thuộc về. Ở đây là kiểu `deck` còn `d` là 1 biến chỉ tới `deck`
