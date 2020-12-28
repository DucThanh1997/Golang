## template HTML
- Template là 1 chuỗi hoặc tệp chứa 1 hoặc nhiều thành phần đặt trong dấu ``{{...}}``

- Hầu hết chuỗi được in ra nguyên xi, nhưng các hành động kích hoạt theo hướng khác.
```
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}
```
Đầu tiên in số lượng các issue, sau đó in number, user, title và tuổi của mỗi vấn đề cái dấu `.` để chỉ đến cái item.Number (viết cho nhanh là như thế). `{{range .Items}} and {{end}}` để tạo 1 vòng lặp

Cái dấu `|` để làm cho kết quả của 1 thao tác thành argument của 1 cái khác. Trong trường hợp Title, thao tác thứ hai là printf function (fmt.Sprintf) trong tất cả các mẫu.

- Để hiển thị template cần 2 bước
    + Đầu tiên chúng ta parse cái template ra thành 1 kiểu biểu diễn phù hợp và rồi parse nó ra vào 1 đầu ra thực sự. Đoạn code dưới đây định nghĩa và parse 1 templ
    ```
    report, err := template.New("report").
    Funcs(template.FuncMap{"daysAgo": daysAgo}).
    Parse(templ)
    if err != nil {
        log.Fatal(err)
    }
    ```

Bởi vì template thường được cố định tại complie time, việc parse không thành công 1 template trong quá trình compile dẫn đến một lỗi nghiêm trọng trong chương trình. Hàm template.Must làm cho việc xử lý lỗi thuận tiện hơn: nó chấp nhận template và error, kiểm tra xem lỗi đó có là nil không (và panic nếu là những lỗi khác), rồi trả về mẫu.




### html/package
- Giống text/package nhưng có thêm các tính năng như tự động ngắt hoặc thêm dấu cách của các chuỗi xuất hiện trong html, JS hoặc CSS. Các tính năng này có thể giúp chúng ta tránh vấn đề bảo bật lâu năm trong việc tạo HTML **an injection attack** 

- Ví dụ 
```
var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
    <th>#</th>
    <th>State</th>
    <th>User</th>
    <th>Title</th>
</tr>
{{range .Items}}
<tr>
    <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
    <td>{{.State}}</td>
    <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
    <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))
```