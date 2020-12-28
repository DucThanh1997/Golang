## Sort
- Trong golang hàm sort.Sort không yêu cầu 1 loại nhất định nào. Thay vào đó, nó dùng `sort.Interface`, để tạo mối quan hệ cho thuật toán sắp xếp và mỗi loại mà đang được sort.

- sort ỉnterface có 3 method, length, cách so sánh và cách swap

```
package sort
type Interface interface {
    Len() int
    Less(i, j int) bool // i, j are indices of sequence elements
    Swap(i, j int)
}
```
