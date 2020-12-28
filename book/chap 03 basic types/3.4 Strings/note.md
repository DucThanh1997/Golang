## strings
- Là 1 chuổi byte bất biến, nó có thể chứa bất cứ thứ gì bao gồm các byte có giá trị 0 nhưng thường là các văn bản ng đọc được

- `len` là built-in function để đếm các số lượng chữ trong 1 chuỗi string

- Bạn cũng có thể truy cập từng chữ trong chuỗi 1
```
s := "hello world"
fmt.Println(s[0])
>> h
fmt.Println(s[0:5]) // "hello"
fmt.Println(s[:5]) // "hello"
fmt.Println(s[7:]) // "world"
fmt.Println(s[:]) // "hello, world"
```

- Chuỗi string dùng được +
```
fmt.Println("goodbye" + s[5:]) // "goodbye, world"
```

- Chuỗi cũng có thể sử dụng các toán tử so sánh ==, <, >

### strings literal
- Lưu ý: 
    + Các string được viết sau dấu `/` sẽ bỏ được các kí tự kết nối với chúng
    ```
    "string\"" (string")
    ```

- Khi bạn sử dụng `..` thì nó sẽ lấy raw string literal, các content sẽ được lấy hết. Nó thường được dùng để viết đoạn văn

### unicode
- Trước chỉ có ASCII (128 kí tự) của mỹ thôi bh có thêm UNICODE (120000 kí tự), nó thu thập và có thể show ra toàn bộ các loại ngôn ngữ lên máy tính 


### Strings and Byte Slices
- Có 4 thư viện chính của strings bao gồm `bytes, strings, strconv, unicode`
    + Thư viện strings cung cấp nhiều function cho việc tìm kiếm, thay thế, so sánh, cắt tỉa, phân tách và nối các biến strings với nhau
    + Thư viện bytes có các function để tính toán các mảng của bytes thuộc kiểu []byte, có nhiều thuộc tính giống string. Vì string không thay đổi được nên việc xây dựng 1 biêns string dần dần cần phân bổ và sao chép nhiều. Trong trường hợp đó sẽ hiệu quả hơn nếu sử dụng bytes.Buffer
    + Thư viện strconv sử dụng các function để chuyển đổi các kiểu dữ liệu về string và ngược lại
    + Thư viện unicode cung cấp các functions như IsDigit, IsLetter, IsUpper, and IsLower

```
s := "abc"
b := []byte(s)
s2 := string(b)
```
Theo khái niệm cái `[]byte(s)` phân bổ 1 mảng byte mới là bản sao của các byte(s) và tạo ra 1 slice tham chiếu đến cả mảng []byte(s). Một trình biên dịch tối ưu hóa có thể tránh được việc phân bổ và sao chép trong một số trường hợp, nhưng nói chung, cần phải sao chép để đảm bảo rằng các byte của s vẫn không thay đổi ngay cả khi các byte của b được sửa đổi sau đó. Việc chuyển đổi từ lát cắt byte trở lại chuỗi có chuỗi (b) cũng tạo ra một bản sao, để đảm bảo tính bất biến của chuỗi kết quả s2.

Thư viện byte cung cấp kiểu Buffer để thao tác hiệu quả với các slice byte. Buffer bắt đầu trống nhưng phát triển khi dữ liệu của các loại như chuỗi, byte và [] byte được ghi vào nó