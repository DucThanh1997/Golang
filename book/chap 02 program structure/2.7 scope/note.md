## Scope    
- Scope là vùng code mà biến có thể được sử dụng
- Lifetime là khoảng thời gian tồn tại của biến trong tình trạng hợp lệ

- Cú pháp `block` bao bọc 1 biến được khởi tạo ở trong nó, biến đó sẽ không truy cập được từ ngoài vùng block đó. Khái quát hơn gọi là `lexical block`
- `universe block` là cái block của cả package

- order decleration không tác động đến scope
```
if f, err := os.Open(fname); err != nil { // compile error: unused: f
    return err
}
f.ReadByte() // compile error: undefined f
f.Close() // compile error: undefined f
```
Sửa lại thì như này
```
f, err := os.Open(fname)
if err != nil {
    return err
}
f.ReadByte()
f.Close()
```
