# Các lệnh điều khiển 
## Vòng for
```
package main

import "fmt"

func main() {
	var i int = 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("Phong cách 2")
	for m := 1; m <= 10; m++ {
		fmt.Println(m)
	}
}
```
In ra từ 1 đến 10
Kiể vòng lặp này ở Go chỉ có mình for thôi không có while

## If
```
package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 0 {
			fmt.Println("không")
		} else if i == 1 {
			fmt.Println("một")
		} else if i == 2 {
			fmt.Println("hai")
		} else {
			fmt.Println("i lớn hơn 2")
		}

	}
}
```
Vòng if có cú pháp như trên
Cũng như các ngôn ngữ khác if có if, else if và else

## switch
```
package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println("Không")
		case 1:
			fmt.Println("Một")
		case 2:
			fmt.Println("Hai")
		case 3:
			fmt.Println("Ba")
		case 4:
			fmt.Println("Bốn")
		default:
			fmt.Println("Không rõ")
		}

	}
}
```

# Array, Slice, Map
## Array
Array (hay mảng) là một tập hợp các phần tử có cùng kiểu dữ liệu nằm liên tiếp nhau. Chúng ta khai báo một array trong Go như sau
```
var x [5]int
x[4] = 100
fmt.Println(x)
```
Nếu chúng ta không gán giá trị cho các phần tử trong mảng nó sẽ auto bằng 0
```
func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total = total + x[i]
	}
	fmt.Println(total / float64(len(x)))
}
```
1 ví dụ khác tính giá trị khác. Ở dòng cuối bạn có thể thấy t phải ép kiểu về float 64 vì len(x) là kiểu int.
Trong go thì khác kiểu dữ liệu thì không tính toán được

```
var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i, value := range x {
		fmt.Println("i:", i)
		fmt.Println("value:", value)
		total = total + value
	}
	fmt.Println(total / float64(len(x)))
  ```
  1 ví dụ khác ở đây t gọi hàm for khác nhanh
  i chỉ số thứ tự còn value = giá trị của phần tử

Ngoài ra chúng ta có thể khai báo nhanh như sau
```
x := [5]float64{ 98, 93, 77, 82, 83 }
```

## Slice
- Định nghĩa: Slice cũng giốngnhư array, các phần tử trong slice cũng được đánh chỉ số. 
Điểm khác biệt giữa slice và array là số phần tử trong slice có thể thay đổi được. 

Chúng ta khai báo một slice như sau:
```
var x []float64
```
Dòng trên tạo 1 slide có 0 phần tử
```
x := make([]float64, 5)
```
Dòng trên tạo 1 slide có 5 phần tử
Ngoài ra còn 1 cách nữa là tạo ra slide từ [low:high]
```
arr := [5]float64{98, 93, 77, 82, 83}
x := arr[0:5]
```
nó sẽ lấy phần tử từ 0 đến 4 còn nếu là [1:4] nó sẽ lấy từ 1 đến 3 nếu muốn lấy hết [:] hoặc [0:]

- Hàm append
```
func main() {
    slice1 := []int{1, 2, 3}
    slice2 := append(slice1, 4, 5)
    fmt.Println(slice1, slice2)
}
```
slice 2 sẽ có các phần tử của slide1 và có thêm phần tử 4 và 5

- Hàm copy
```
func main() {
    slice1 := []int{1, 2, 3}
    slice2 := make([]int, 2)
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)
}
```
Dòng copy(slice2, slice1) sẽ sao chép các phần tử trong slice1 vào slice2,
tuy nhiên khi tạo slice2 chỉ có 2 phần tử nên hàm này chỉ sao chép 2 phần tử đầu tiên của slice1 vào slice2, 
do đó slice2 sẽ có các phần tử là [1, 2].

## Map
Giống dictionary

Cách khai báo
```
x := make(map[string]int)
```
string là kiểu dữ liệu của key 

còn int là kiểu dữ liệu của value

```
func main() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
}
```
In ra giá trị của key là 10

```
func main() {
    x := make(map[string]int)
    x["key"] = 10
  
    value, ok := x["key"]
    fmt.Println(value, ok)
   
    value2, ok2 := x["key2"]
    fmt.Println(value2, ok2)
}
```
Trong đoạn code trên, map x không có khóa key2, khi truy xuất giá trị sẽ trả về 0, ngoài ra giá trị boolean sẽ là false.
Còn khóa key có tồn tại nên sẽ trả về số 10 và giá trị boolean của nó là true. 
Ngoài ra ở đây chúng ta còn biết được là Go cho phép chúng ta gán nhiều giá trị vào nhiều biến trên một dòng thông qua dấu phẩy ",".

Chúng ta cũng có thể vừa tạo vừa gán map
```
elements := map[string]string{
    "H": "Hydrogen",
    "He": "Helium",
    "Li": "Lithium",
    "Be" : "Beryllium",
    "B" : "Boron",
    "C" : "Carbon",
    "N" : "Nitrogen",
    "O" : "Oxygen",
    "F" : "Fluorine",
    "Ne" : "Neon",
}
```

# Hàm
## Khái báo
```
package main
 
import "fmt"
 
func average(input []float64) float64 {
    total := 0.0
    for _, v := range input {
        total += v
    }
    return total / float64(len(input))
}
 
func main() {
    xs := []float64{98, 93, 77, 82, 83}
    fmt.Println(average(xs))
}
```

Đầu tiên gọi func + tên hàm + (kiểu dữ liệu đầu vào) + kiểu dữ liệu đầu ra

- Các hàm có thể gọi chồng lên nhau
```
func main() {
    fmt.Println(f1())
}
 
func f1() int {
    return f2() 
}
 
func f2() int {
    return 1
}
```

## Đặc điểm
- Có thể trả về nhiều giá trị
```
func f() (int, int) {
    return 5, 6
}
 
func main() {
  x, y := f()
  fmt.println(x,y)
}
```

- Có thể chuyền nhiều biến đầu vào 
```
package main
 
import "fmt"
 
func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    } 
    return total
}
 
func main() {
    fmt.Println(add(1, 2, 3))
}
```

- Hàm lồng nhau
Trong hàm có thể tạo ra 1 hàm nữa
```
package main
 
import "fmt"
 
func main() {
    add := func(x, y int) int {
        return x + y
    }
    fmt.Println(add(1, 1))
}
```

## defer, panic, recover
- defer
```
package main
 
import "fmt"
 
func first() {
    fmt.Println("1st")
}
 
func second() {
    fmt.Println("2nd")
}
 
func main() {
    defer second()
    first()
}
```
Trong ví dụ trên, hàm first() sẽ thực hiện đầu tiên, sau đó đến hàm second() mặc dù hàm second() đứng trước, 
lý do là bởi vì chúng ta thêm từ khóa defer vào trước second(), 
do đó hàm second() sẽ được thực hiện cuối cùng khi tất cả các công việc khác đã hoàn tất.

- panic
```
package main
 
func main() {
    panic("Co loi xay ra") 
}
```
panic để dừng chương trình và phát sinh lỗi. 

Chẳng hạn như khi chúng ta viết chương trình tính phép chia mà người dùng nhập vào mẫu số là 0 thì chúng ta có thể dùng hàm panic để báo lỗi.


















