## interface Value
- Về mặt khái niệm 1 giá trị của kiểu interface có 2 thành phần
    + 1 kiểu nhất định
    + và giá trị của kiểu đó

```
var w io.Writer
w=os.Stdout
w=new(bytes.Buffer)
w=nil
```

Ở dòng đầu tiên `var w io.Writer` khai báo w. Trong Go, biến luôn luôn được khởi tạo và gán vào 1 giá trị, interfaces cũng thế. Zero-value cỉa interface là `nil` ở các type và value

![image](https://user-images.githubusercontent.com/45547213/103144045-390ede00-4755-11eb-9147-7b5e719e605b.png)

Ở dòng thứ hai, w được gán với giá trị của kiểu *os.File:
```
w=os.Stdout
```

Phép gán này liên quan đến việc chuyển đổi ngầm định từ kiểu cụ thể sáng kiểu interface. Kiểu chuyển đổi này dù là ngầm định hay rõ ràng, nó đều sẽ lấy kiểu và giá trị của toán tử được gán cho. Interface dynamic type được đặt theo theo toán tử gán vào là `os.File`. Và giá trị động của nó `dynamic type` sẽ giữ 1 bản copy của `os.Stdout` mà có con trỏ trỏ đến `os.File`

![image](https://user-images.githubusercontent.com/45547213/103144222-6315cf80-4758-11eb-882c-851adb6488ff.png)

- Việc gọi method write trên 1 interface value chứa con trỏ `*os.File`

```
w.Write([]byte("hello")) // "hello"
```

Chúng ta không thể biết vào thời điểm compile kiểu của interface value sẽ là gì, nên khi động đến interface bạn cần **dynamic dispatch**. Thay vì gọi thẳng, compiler sẽ gen code ra để lấy địa chỉ của method tên là write từ `type descriptor`. Sau đó gọi gián tiếp đến method qua cái addres.

- Interface value có thể so sánh kiểu == và !=. 2 interface value bằng nhau nếu cả 2 bằng nil, hoặc kiểu của chúng giống nhau và các giá trị bằng nhau . Vì interface value có thể so sánh nên có thể dùng làm key của map. Nhưng nếu kiểu của nó là slice chẳng hạn thì sẽ không so sánh được


### Caveat: An Int erface Containing a Nil Pointer Is Non-Nil