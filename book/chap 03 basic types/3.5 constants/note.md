## Constants
- Là 1 biểu thức mà giá trị được biết bởi compiler và sự khởi tạo được xảy ra trong quá trình compile chứ không phải ở run-time. Kiểu thực sự của constant là 1 kiểu dữ liệu cơ bản: boolean, string và number

- Một khai báo const xác định các giá trị được đặt tên trông giống như các biến về mặt cú pháp nhưng có giá trị không đổi, điều này ngăn chặn các thay đổi ngẫu nhiên (hoặc bất chính) trong quá trình thực thi chương trình.

- Nhiều phép tính toán có thể chạy trên constant khi ở compile time, giảm thiểu các công việc cần làm ở run-time đồng thời giúp tối ưu các compiler khác. Các lỗi thông thường được phát hiện trong thời gian chạy có thể được báo cáo tại thời điểm biên dịch khi toán hạng của chúng là Constant, chẳng hạn như phép chia số nguyên cho 0.


### Untyped Constant 
Các Constant trong go hơi bất thường. Mặc dù một Constant có thể có bất kỳ kiểu dữ liệu cơ bản nào như int hoặc float64, bao gồm cả các kiểu cơ bản được đặt tên như time.Duration, nhưng nhiều Constant không cam kết với một kiểu cụ thể. compiler biểu diễn các hằng số không được cam kết này với độ chính xác số lớn hơn nhiều so với các giá trị của các kiểu cơ bản, và số học trên chúng chính xác hơn số học máy.

Các constant không được định kiểu này có thể có sử dụng với nhiều kiểu dữ liệu khác mà không cần biến đổi 