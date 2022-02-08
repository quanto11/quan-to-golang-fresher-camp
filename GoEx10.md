# Câu hỏi: 
Vì sao không nên chứa file upload vào ngay chính bên trong service mà nên dùng Cloud. Vì sao không chứa binary ảnh vào DB?
# Trả lời:
- File upload bên trong service sẽ làm cho service phải chứa nhiều file và trở nên nặng hơn.
- chứa binary ảnh vào DB làm cho DB nặng, ta chỉ nên lưu url của ảnh.
