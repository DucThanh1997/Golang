## Map
- Hash table là một trong những cấu trúc dữ liệu khéo léo và linh hoạt nhất. Nó là một tập hợp các cặp key / value không có thứ tự trong đó tất cả các key đều khác nhau và value được liên kết với một key nhất định để có thể được truy xuất, cập nhật hoặc xóa

- Trong go, map là 1 cái giống hash table và map được viết bởi **map[k]V**. Trong đó k là key và v là value. 

- Tất cả các key phải có cùng kiểu, tất cả các value cũng phải có cùng kiểu. Tuy nhiên key và value không nhất thiết phải cùng kiểu với nhau

- Kiểu của key phải so sánh được để biết được rằng key đó đã tồn tại hay chưa

- Chúng ta sử dụng `make` để tạo map
```
ages := make(map[string]int)
```

- Chúng ta cũng có thể sử dụng map literal
```
ages := map[string]int{
    "alice": 31,
    "charlie": 34,
}
```
Cái này bằng với
```
ages := make(map[string]int)
ages["alice"] = 31
ages["charlie"] = 34
```

- Cách access 1 value
```
ages["alice"] = 32
fmt.Println(ages["alice"]) // "32"
```

- Cách xóa 1 cặp key value
```
delete(ages, "alice") // remove element ages["alice"]
```

- Để loop qua toàn bộ key và value
```
for name, age := range ages {
    fmt.Printf("%s\t%d\n", name, age)
}
```

- Lướt kiểu trên này không có order nếu muốn có order thì phải làm như này
```
import "sort"
var names []string
for name := range ages {
    names = append(names, name)
}
sort.Strings(names)
for _, name := range names {
    fmt.Printf("%s\t%d\n", name, ages[name])
}
```

- Kiểm tra xem map có key bob hay không
```
age, ok := ages["bob"]
if !ok { /* "bob" is not a key in this map; age == 0. */ }
```

- Map cũng không compare được với nhau, chúng ta phải làm bằng tay

```
func equal(x, y map[string]int) bool {
    if len(x) != len(y) {
        return false
    }
    for k, xv := range x {
        if yv, ok := y[k]; !ok || yv != xv {
            return false
        }
    }
    return true
}