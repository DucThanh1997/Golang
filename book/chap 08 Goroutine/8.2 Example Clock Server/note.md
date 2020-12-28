## Clock Server
- Hàm listen tạo 1 net.Listener (1 object) listen tới các connection ở 1 cổng network. Trong trường hợp này là `localhost:8000`.

- `listener.Accept()` method bị block cho đến khi có 1 connect đến và trả ra 1 `net.Conn`

- `handleConn` sẽ xử lý connection đến. Trong vòng lặp nó trả ra thời gian hiện tại tới client

- Tuy nhiên nếu có 1 người đến sau cũng request vào thì người này phải đợi đến khi người trước request xong, như thế đéo chấp nhận được nên phải thêm go routine vào