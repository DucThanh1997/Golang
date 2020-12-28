## Multiple Return Values
- 1 function có thể return nhiều hơn 1 giá trị.
```
func findLinks(url string) ([]string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
    }
    return visit(nil, doc), nil
}
```

- Chúng ta có thể ignore 1 giá trị trả về bằng cách thay vào dấu `_`
```
links, _ := findLinks(url) // errors ignored
```

- Chúng ta cũng có thể trả về dữ liệu được trả về từ 1 function khác
```
func findLinksLog(url string) ([]string, error) {
    log.Printf("findLinks %s", url)
    return findLinks(url)
}
```

- Trong 1 function mà đầu ra được đặt tên thì những toán từ trả ra có thể bỏ qua, đây gọi là **bare return**
```
// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        err = fmt.Errorf("parsing HTML: %s", err)
        return
    }
    words, images = countWordsAndImages(doc)
    return
}
Ở code trên nó sẽ trả ra word và images vả err


