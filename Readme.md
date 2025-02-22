# Chat Application with ScyllaDB and Elasticsearch

## Giới thiệu

Đây là một ứng dụng chat nhỏ sử dụng **ScyllaDB** để lưu trữ tin nhắn và **Elasticsearch** để tìm kiếm nhanh. Ứng dụng được viết bằng **Golang** và có giao diện frontend bằng **ReactJS**.

## Công nghệ sử dụng

- **Golang**: Ngôn ngữ lập trình chính.
- **ScyllaDB**: Cơ sở dữ liệu NoSQL chịu tải cao để lưu trữ tin nhắn.
- **Elasticsearch**: Công cụ tìm kiếm mạnh mẽ để tìm kiếm tin nhắn theo nội dung.
- **Docker Compose**: Quản lý các container chạy ScyllaDB và Elasticsearch.
- **Gocql**: Thư viện Golang để kết nối với ScyllaDB.
- **Go-Zero**: Router HTTP cho Golang.
- **ReactJS**: Frontend framework để tạo giao diện người dùng.
- **Tailwind CSS**: Thư viện CSS để tạo giao diện nhanh chóng và đẹp mắt.

## Cài đặt

### Yêu cầu hệ thống

- Docker và Docker Compose
- Golang 1.20+
- Node.js 16+

### Bước cài đặt

1. Clone repository:
   ```bash
   git clone https://github.com/your-repo/chat-app.git
   cd chat-app
   ```
2. Khởi chạy cụm ScyllaDB bằng Docker Compose:
   ```bash
   docker-compose up -d
   ```
3. Cài đặt các thư viện Golang:
   ```bash
   go mod tidy
   ```
4. Chạy backend:
   ```bash
   go run main.go
   ```
5. Cài đặt frontend:
   ```bash
   cd frontend
   npm install
   npm start
   ```

## API Endpoints

- `POST /messages` - Gửi tin nhắn
- `GET /messages/search` - Tìm kiếm tin nhắn

## Giao diện ReactJS

Giao diện frontend gồm các tính năng:
- Hiển thị danh sách cuộc trò chuyện
- Gửi và nhận tin nhắn theo thời gian thực
- Tìm kiếm tin nhắn theo nội dung

## Liên hệ

Nếu bạn có bất kỳ câu hỏi nào, vui lòng liên hệ qua email: [**thoahlgbg2002@gmail.com**](mailto:thoahlgbg2002@gmail.com)

