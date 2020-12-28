## Anonymous Function
- Nó là kiểu 1 function được viết ngay trong hàm xử lý của 1 cái khác và nó không có tên
```
func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}

func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}
```

- Function kiểu này có thể lấy biến của function chứa nó. Như ở đây việc gọi function square sẽ tạo 1 biến local và trả về 1 anonymous function mà mỗi lần function đó gọi sẽ tăng x lên 1 đơn vị

- Function ngoài code ra nó có cả state

- Đối với anonymous function mà có truy hồi thì phải khai báo nó trước bằng var sau đó gọi lại còn nếu như này thì sẽ lỗi ngay

### Caveat: Capturing Iteration Variables
```
var rmdirs []func()

for _, d := range tempDirs() {
    dir := d // NOTE: necessary!
    os.MkdirAll(dir, 0755) // creates parent directories too
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir)
    })
}

// ...do some work...
for _, rmdir := range rmdirs {
    rmdir() // clean up
}
```

- Cái `d` này sẽ chỉ được init 1 lần thôi muốn nó được init nhiều lần thì phải gán vào dir

