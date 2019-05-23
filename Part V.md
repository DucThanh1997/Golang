# String

## Các hàm của string
```
package main
 
import (
    "fmt"
    s "strings"
)
 
func main() {
    fmt.Println("Contains:  ", s.Contains("test", "es"))                  trả về true nếu test chứa es
    fmt.Println("Count:     ", s.Count("test", "t"))                      đếm số lần t xuất hiện trong test
    fmt.Println("HasPrefix: ", s.HasPrefix("test", "te"))                 trả về true nếu prefix là te
    fmt.Println("HasSuffix: ", s.HasSuffix("test", "st"))                 trả về true nếu suffix là st
    fmt.Println("Index:     ", s.Index("test", "e"))                      trả về vị trí của e trong test  
    fmt.Println("Join:      ", s.Join([]string{"a", "b"}, "-"))           join chuỗi a với chuỗi b a - b
    fmt.Println("Repeat:    ", s.Repeat("a", 5))                          in ra a 5 lần
    fmt.Println("Replace:   ", s.Replace("foo", "o", "0", -1))            in ra f00
    fmt.Println("Replace:   ", s.Replace("foo", "o", "0", 1))             in ra f0o
    fmt.Println("Split:     ", s.Split("a-b-c-d-e", "-"))                 đưa a b c d e vào array
    fmt.Println("ToLower:   ", s.ToLower("TEST"))
    fmt.Println("ToUpper:   ", s.ToUpper("test"))  
 
    fmt.Println("Len: ", len("hello"))
    fmt.Println("Char: ", "hello"[1])
}
```

## Các chuỗi của string
```
package main
 
import "fmt"
import "os"
 
type point struct {
    x, y int
}
 
func main() {
    p := point{1, 2}
 
    fmt.Printf("%v\n", p)
    fmt.Printf("%v+\n", p)
    fmt.Printf("%#v\n", p)
    fmt.Printf("%T\n", p)
  
    fmt.Printf("%t\n", true)
 
    fmt.Printf("%d\n", 123)
    fmt.Printf("%b\n", 14)
    fmt.Printf("%c\n", 33)
    fmt.Printf("%x\n", 456)
 
    fmt.Printf("%f\n", 78.9)
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)
 
    fmt.Printf("%s\n", "\"string\"")
    fmt.Printf("%q\n", "\"string\"")
 
    fmt.Printf("%x\n", "hex this")
    fmt.Printf("%p\n", &p)
    fmt.Printf("|%6d|%6d|", 12, 345)
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
    fmt.Printf("|%-6.2f|%-6.2f|", 1.2, 3.45)
    fmt.Printf("|%6s|%6s|\n", "foo", "b")
    fmt.Printf("|%-6s|%-6s|", "foo", "b")
  
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)
}
```

- %v: in các giá trị của một biến struct
- %+v: in các giá trị kèm với tên trường của biến struct
- %#v: giống %+v kèm theo tên kiểu dữ liệu của struct và tên hàm đã gọi nó
- %T: in tên struct và tên hàm đã gọi nó
- %t: in giá trị boolean
- %d: in số nguyên (hệ 10)
- %b: in số nguyên dưới dạng số nhị phân (hệ 2)
- %c: in kí tự dựa theo mã ASCII
- %x: in số nguyên dưới dạng số thập lục phân (hệ 16) hoặc chuyển một chuỗi thành số thập lục phân
- %f: in số thập phân
- %e và %E: in số thập phân dưới dạng số mũ
- %s: in một chuỗi
- %q: in một chuỗi có 2 cặp dấu nháy kép “”
- %6d: in một số nguyên, nếu số đó không đủ 6 kí tự thì tự động thêm các khoảng trống vào bên trái cho đủ 6 kí tự
- %6.2f: in số thập phân, làm tròn đến 2 chữ số thập phân, sau đó nếu phần thập phân và phần nguyên cùng với dấu chấm không đủ 6 kí tự thì tự động thêm các khoảng trống vào bên trái cho đủ 6 kí tự
- %-6.2f: tương tự với %6.2f nhưng các khoảng trống được thêm vào bên phải
- %6s: in một chuỗi, nếu chuỗi không đủ 6 kí tự thì thêm các khoảng trống vào bên trái cho đủ
- %-6s: tương tự %6s nhưng thêm các khoảng trống vào bên phải


# File và folder
```
package main
 
import (
    "fmt"
    "os"
)
 
func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    defer file.Close()
  
    stat, err := file.Stat()
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
  
    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
  
    str := string(bs)
    fmt.Println(str)
} 
```

- Mở file dùng `file.Open`
- Đọc file dùng `file.Read`
- Tạo file dùng `file.Create`
- Ghi file dùng `file.WriteString`




















