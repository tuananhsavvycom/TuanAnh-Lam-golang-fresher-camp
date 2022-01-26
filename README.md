# Công dụng của Index trong DB :
-Index là một cấu trúc dữ liệu được dùng để định vị và truy cập nhanh nhất vào dữ liệu trong các bảng database.
-Index là một cách tối ưu hiệu suất truy vấn database bằng việc giảm lượng truy cập vào bộ nhớ khi thực hiện truy vấn.
-Example :
-Giả sử ta có một bảng User lưu thông tin của người dùng, ta muốn lấy ra thông tin của người dùng có trường tên (Name) là “HienNguyen” . Ta có truy vấn SQL sau: SELECT * FROM User WHERE Name = 'HienNguyen';
-Nếu không có Index cho cột Name, truy vấn sẽ phải chạy qua tất cả các Row của bảng User để so sánh và lấy ra những Row thỏa mãn. Vì vậy, khi số lượng bản ghi lớn, chuyện gì sẽ xảy ra ??  Index được sinh ra để giải quyết vấn đề này
# Nếu không có index thì khác biệt trong truy vấn là gì?
-Sử dụng index tốc độ truy vấn sẽ tăng lên nhưng đi kèm với một cái giá: việc tạo index sẽ yêu cầu dung lượng ổ đĩa và có thể làm chậm các hoạt động khác. Đặc biệt, khi insert hay update column sử dụng index, index cần được điều chỉnh (reindex) sẽ tiêu tốn một khoảng thời gian. Việc đánh index bừa bãi, lộn xộn, không những không tăng tốc độ truy vấn mà còn làm giảm performance.
