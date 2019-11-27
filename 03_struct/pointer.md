## pointer

![image](https://user-images.githubusercontent.com/45547213/69689479-a89e3e00-10fb-11ea-93ea-5a0bd1d9b4f6.png)

- Khi bạn gọi hàm khác thì nó sẽ sao chép cái biến hay giá trị đó sang 1 địa chỉ khác rồi sửa ở trong đó

```
func main() {
    mySlice := []{"Hi", "There", "How"}

    updateSlice(mySlice)

    fmt.Println(mySlice)
}

func updateSlice(s []string) {
    s[0] = "Bye"
}

```

đoạn code này vẫn có khá năng thay đổi cái gốc. Tại sao?

- Thứ nhất khi bạn tạo 1 slice nó sẽ tạo ra những thứ như này

![image](https://user-images.githubusercontent.com/45547213/69693229-48ad9480-1107-11ea-9f15-67d93214d456.png)

    + pointer to head sẽ trỏ đến 1 cái array và array chứa các phần tử của slice và tạo ra các thứ có cùng slice

    + capacity: là số phần tử chứa được

    + length: là số phần tử hiện nó có

- Khi bạn gọi hàm update như mọi khi nó cũng sao chép cái đấy nhưng ảnh dưới đây sẽ chỉ rõ

![image](https://user-images.githubusercontent.com/45547213/69693383-c1acec00-1107-11ea-924c-db1fcccbad1d.png)

- vì nó copy nên mọi thứ vẫn giữ nguyên nên cái giá trị được copy vẫn trỏ đến string



![image](https://user-images.githubusercontent.com/45547213/69693659-b8704f00-1108-11ea-92cd-e92ef0fdc2bf.png)

- Reference type không cần pointer

- Còn value type thì vẫn cần