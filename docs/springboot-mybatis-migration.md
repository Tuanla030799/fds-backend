# Kế hoạch chuyển đổi FDS Backend sang Spring Boot + MyBatis

## 1) Mục tiêu kiến trúc
- Runtime: Java 21 + Spring Boot 3.
- Data access: MyBatis XML (ưu tiên SQL tường minh cho use case phức tạp).
- DB: PostgreSQL, connection pool HikariCP.
- Quan sát hiệu năng: Actuator + Micrometer/Prometheus + slow query log.

## 2) Chiến lược chuyển đổi an toàn (Strangler Pattern)
1. **Giữ Go service đang chạy** cho production.
2. Dựng song song `springboot-backend` và mở endpoint mới `/api/v2/*`.
3. Chuyển từng module theo thứ tự:
   - `preset` (đọc nhiều, dễ benchmark)
   - `design-submissions`
   - `admin-auth`
4. So sánh output 2 hệ thống bằng contract tests trước khi cutover.
5. Chuyển traffic theo tỷ lệ (canary) rồi tắt endpoint Go tương ứng.

## 3) SQL nâng cao nên áp dụng ngay
- Dynamic filter/sort bằng MyBatis XML (`<if>`, `<choose>`, `<trim>`).
- Pagination ổn định + `COUNT(*) OVER()` để trả về `total_count` cùng 1 query.
- Full text search cho tên/ghi chú qua `GIN index` + `to_tsvector` khi dữ liệu lớn.
- Batch update/insert bằng `foreach` để giảm round-trip.

## 4) Performance tuning checklist
- **Index bắt buộc**:
  - `presets(status, sort_order, created_at desc)`.
  - `design_submissions(status, created_at desc)`.
  - `admin_users(email)` unique.
- **Query hygiene**:
  - Tránh `SELECT *`, luôn select cột cần dùng.
  - Chặn page size tối đa 100.
  - Timeout statement MyBatis: 3s (đã cấu hình mẫu).
- **Pool tuning Hikari**:
  - max pool khởi điểm: 20.
  - Theo dõi `hikaricp.connections.active` để điều chỉnh.
- **Đo lường**:
  - p95 latency theo endpoint.
  - DB time / request.
  - Tỷ lệ cache hit (nếu thêm Redis).

## 5) Kết quả đã scaffold trong repo
- `springboot-backend/` chứa khung app Spring Boot + MyBatis.
- Có sẵn ví dụ endpoint `/api/v2/presets` dùng SQL động + window function.
- Có sẵn cấu hình connection pool và timeout cơ bản.

## 6) Lệnh chạy thử
```bash
cd springboot-backend
mvn spring-boot:run
```

> Gợi ý: Sau khi chạy ổn, bước tiếp theo là thêm Flyway/Liquibase để quản lý migration schema thay cho AutoMigrate hiện tại.
