## pointer

![image](https://user-images.githubusercontent.com/45547213/69689479-a89e3e00-10fb-11ea-93ea-5a0bd1d9b4f6.png)

- Khi bạn gọi hàm khác thì nó sẽ sao chép cái biến hay giá trị đó sang 1 địa chỉ khác rồi sửa ở trong đó, đó gọi là **pass by value**

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

    Một Slice sẽ có 2 thuộc tính là length (len) và capacity (cap). Length là số phần tử chứa trong Slice, còn capacity là số phần tử chứa trong Array mà Slice tham chiếu đến (bắt đầu tính từ phần tử đầu tiên của Slice). Để lấy ra length của Slice ta dùng hàm len(), còn để lấy ra capacity thì ta dùng hàm cap()


    ```
    s := []int{2, 3, 5, 7, 11, 13}

    s = s[0:0]   // s = [], len(s) = 0, cap(s) = 6
    s = s[0:4]   // s = [2, 3, 5, 7], len(s) = 4, cap(s) = 6
    s = s[2:4]   // s = [5, 7], len(s) = 2, cap(s) = 4, cap được tính từ vị trí số 2 trở đi
    s = s[2:]   // s = [5, 7, 11, 13], len(s) = 4, cap(s) = 4

    ```

    + capacity: là số phần tử chứa được

    + length: là số phần tử hiện nó có



- Khi bạn gọi hàm update như mọi khi nó cũng sao chép cái đấy nhưng ảnh dưới đây sẽ chỉ rõ

![image](https://user-images.githubusercontent.com/45547213/69693383-c1acec00-1107-11ea-924c-db1fcccbad1d.png)

- vì nó copy nên mọi thứ vẫn giữ nguyên nên cái giá trị được copy vẫn trỏ đến string



![image](https://user-images.githubusercontent.com/45547213/69693659-b8704f00-1108-11ea-92cd-e92ef0fdc2bf.png)

- Reference type không cần pointer

- Còn value type thì vẫn cần