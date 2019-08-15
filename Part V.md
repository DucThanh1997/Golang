## Rune
Là 1 kiểu dữ liệu gần giống với int32. Rune đại diện cho 1 điểm mã trong go
```
package main

import (  
   "fmt"
)

func printCharsAndBytes(s string) {  
   for index, rune := range s {
       fmt.Printf("%c starts at byte %d\n", rune, index)
   }
}

func main() {  
   name := "Señor"
   printCharsAndBytes(name)
}
```

String là kiểu dữ liệu bất biến không thể thay đổi được nên nếu bạn muốn thay đổi nó thì bạn phải đưa về rune rồi ép kiểu trả lại về string
```
package main

import (  
   "fmt"
)

func mutate(s []rune) string {  
   s[0] = 'a' 
   return string(s)
}
func main() {  
   h := "hello"
   fmt.Println(mutate([]rune(h)))
}
```
