Câu hỏi:  
    Vì sao trong khoá học này các bạn được khuyên không nên dùng khoá ngoại (FK), điểm yếu của khoá ngoại là gì?  
Trả lời:  
    Khi sử dụng khóa ngoại dữ liệu được nhất quán, ràng buộc với nhau, hiệu năng đọc dữ liệu sẽ tăng.  
    Tuy nhiên hiệu năng thêm, sửa, xóa dữ liệu sẽ bị giảm đi. Nếu sử dụng khóa ngoại:  
        + Khi thêm/ sửa dữ liệu, hệ thống sẽ phải kiểm tra xem những khóa ngoại đó có tồn tại hay không. Khóa ngoại phải tồn tại thì mới có thể thêm/ sửa dữ liệu.  
        + Khi xóa dữ liệu, hệ thống sẽ phải kiểm tra xem có những bảng con nào tham chiếu tới nó.  
    Vì vậy ta không nên sử dụng khóa ngoại để tăng hiệu năng thêm, sửa, xóa dữ liệu. Chỉ sử dụng khóa ngoại trong trường hợp các dữ liệu cần phải nhất quán với nhau.