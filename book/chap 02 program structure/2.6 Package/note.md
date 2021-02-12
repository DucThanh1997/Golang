## Package
Package trong go có cùng mục đích giống với library hoặc module trong các ngôn ngữ khác, hỗ trợ đóng gói, tái sử dụng

Package còn giúp chúng ta che giấu các thông tin bằng cách kiểm soát biến hoặc hàm nào sẽ có thể sử dụng được bên ngoài nó

Mỗi file .go đều bắt đầu bằng cách khai báo tên của package

### 2.6.1 Import
Mỗi package sẽ có 1 cái tên riêng ko cái nào giống cái nào

Theo convention thì tên của package nằm ở phần cuối `gopl.io/ch2/tempconv` là tempconv

### 2.6.2 Package Init
go sẽ chạy từ dưới lên trên để tránh việc p import q nhưng p đã được khởi tạo rồi nhưng nó vẫn chưa biết gì về q