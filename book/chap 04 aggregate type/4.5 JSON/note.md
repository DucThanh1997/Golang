## JSON
- Là 1 kiểu dữ liệu tiêu chuẩn để gửi và nhận thông tin. Go có hỗ trợ nó bởi thư viện `encoding/json`

- JSON là một chuỗi giá trị có thứ tự, được viết dưới dạng danh sách được phân tách bằng dấu phẩy được đặt trong dấu ngoặc vuông;

- Convert dữ liệu của go sang json gọi là **marshalling**. Ta có hàm `json.Marshal`

```
type Movie struct {
    Title   string
    Year    int `json:"released"`
    Color   bool `json:"color,omitempty"`
    Actors  []string
}
var movies = []Movie{
    {Title: "Casablanca", Year: 1942, Color: false,
        Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
    {Title: "Cool Hand Luke", Year: 1967, Color: true,
        Actors: []string{"Paul Newman"}},
    {Title: "Bullitt", Year: 1968, Color: true,
        Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
// ...
}

data, err := json.Marshal(movies)
if err != nil {
log.Fatalf("JSON marshaling failed: %s", err)
}
fmt.Printf("%s\n", data)
```

Kiểu này in ra dữ liệu sẽ bị xấu, để làm đẹp nó ta dùng `json.MarshalIndent`
```
data, err := json.MarshalIndent(movies, "", " ")
if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
}
fmt.Printf("%s\n", data)
```

- Chỉ những field được export ra ngoài mới có thể marshal

- Ngược lại từ json sang data structure chúng ta dùng json.Unmarshal