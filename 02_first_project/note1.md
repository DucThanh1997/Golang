## 1. variable

```
var bien1 string = "thanh"
```
- Trên đây là cách khởi tạo 1 biến với 

    + `var` viết tắt cho variable để chỉ chúng ta sẽ
    bắt đầu tạo ra 1 biến mới
    + `bien1` là tên của biến (chúng ta sẽ đặt tên cho biến)
    + `string` là kiểu của biến. Go là 1 ngôn ngữ static type nên khi đặt biến xong xuôi rồi thì chúng ta sẽ không đặt được kiểu biến khác cho nó

```
bien1 := "thanh"
```
- Trên đây là 1 cách để tạo biến nhanh cho golang mà không cần gõ var hay nêu ra kiểu biến, go nó tự hiểu

- Lưu ý chỉ dùng `:=` khi tạo 1 biến mới còn `=` là để đặt lại giá trị cho nó thôi



## 2. function
```
func newCard() string {
    return "Thanh"
}
```

- Khi tạo 1 function từ đầu tiên cần gần là `func`
- Sau đó chúng ta sẽ đặt tên cho function ở đây tên của function là `newCard`
- Và khi tạo 1 function chúng ta cũng phải đặt luôn kiểu dữ liệu mà chúng ta sẽ `return` hay nói cách khác là trả về. Ở ví dụ trên là kiểu `string`



## Slice và vòng for
- **Array** là 1 giữ 1 list với số lượng được fix cứng và không thay đổi được

- **Slice** là 1 kiểu array nhưng không bị fix cứng

- Slice và array chỉ chứa được 1 kiểu dữ liệu

Cách tạo 1 slice
```
cards := []string{"a", "b"}
hoặc
cards := []string{"a", newCard()}

```
- `[]string` dùng để bảo go là hãy tạo 1 slice chứa các giá trị kiểu string

- {} dùng để đưa các phần tử của slice vào trong đó

Cách thêm phần tử vào slice
```
cards = append(cards, "Hiền")
```

- `append` để add thêm các giá trị của slice vào nó. Tuy nhiên nó sẽ trả ra 1 slice khác nhưng chúng ta cũng có thể reassign lại cái slice cũ cũng được

Cách để tạo 1 vòng for dùng để lướt qua slice
```
for i, card := range cards {
    fmt.Println("card: ", card)
}

- `range card` với từ range sẽ bảo go là mày phải lướt qua cái slice đằng sau từ range này

- `index` là chỉ số của nó

- `card` là giá trị của từng index trong slice cards
