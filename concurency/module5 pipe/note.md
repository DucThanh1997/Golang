## Pipeline
- Định nghĩa
    + Nó là 1 loạt các function kết nối với nhau bằng channel để xử lý 1 việc cụ thể
    + Mỗi 1 function chạy bằng goroutine
    + Tối đa hóa ưu thế của đa luồng

## Fan in Fan out
- Khi có 1 công đoạn trong go routine tốn nhiều thời gian thì chúng ta sẽ break nhỏ nó ra thành nhiều thằng cùng làm 1 việc đó. Ví dụ bước 1 chỉ có 1 thằng làm nhưng bước 2 có 3 thằng làm. Sau đó chúng ta sẽ merge kết quả về 1 channel để gửi sang bước tiếp theo