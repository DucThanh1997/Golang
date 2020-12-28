## 7.3 Interface Satisfacation
- 1 kiểu thỏa mãn interface nếu nó có tất cả các method mà interface đó requires

```
var w io.Writer
w=os.Stdout // OK: *os.File has Write method
w=new(bytes.Buffer) // OK: *bytes.Buffer has Write method
w=time.Second // compile error: time.Duration lacks Write method

var rwc io.ReadWriteCloser
rwc = os.Stdout // OK: *os.File has Read, Write, Close methods
rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method
```

```
w=rwc // OK: io.ReadWriteCloser has Write method
rwc = w // compile error: io.Writer lacks Close method
```

Ở các section trước bạn có thể truyền vào biên a dù method đòi *a nhưng ở interface thì như thế không được bạn phải truyền vào đúng *a thì mới thỏa mãn interface nhé

- Vậy interface{} làm gì nó chả có method nào cả, vì nó không có điều kiện là method nào nên kiểu nào nó cũng chấp nhận được và bạn có thể assign bất kì kiểu giá trị nào vào nó cũng được

