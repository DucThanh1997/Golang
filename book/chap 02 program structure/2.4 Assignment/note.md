## Assignments
Mỗi giá trị được lưu trữ bởi 1 biến được update bằng **assignment**

```
x=1 // named variable
*p = true // indirect variable
person.name = "bob" // struct field
count[x] = count[x] * scale // array or slice or map element
```

### tuple assignment
assign giá trị cho nhiều biến
```
x, y = y, x

a[i], a[j] = a[j], a[i]

i, j, k = 2, 3, 5

f, err = os.Open("foo.txt")
```

### Assignability
- Assignment luôn đúng miễn là biến bên trái cùng vị trí với biến bên phải phải cùng kiểu
