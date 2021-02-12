## Import Paths
Mỗi package được identify bởi 1 đoạn string unique được gọi là `import path`. `import path` là đoạn string xuất hiện trong phần import
```
import (
    "fmt"
    "math/rand"
    "encoding/json"
    "golang.org/x/net/html"
    "github.com/go-sql-driver/mysql"
)
```

Đối với các package mà bạn muốn publiosh thì phần `import path` nên unique. Để tránh conflict, `import path` phải khác tên của các standard library  
