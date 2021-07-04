- Tại sao cần concurrency:
    + Để chạy nhanh hơn thì chúng ta cần app chia nhỏ các task ra rồi chạy song song chứ không phải là làm từng thứ một

- Paralesism là song song:
    + Thằng A làm xong xong task A, thằng B làm song song với thằng A nhưng làm task B

- COncurrency
    + Không phải song song mà là đồng thời, ví dụ thằng A gửi request đến thằng B trong lúc đợi response từ thằng B nó làm tiếp các việc khác rồi đợi có response nó mới hoàn thành công việc