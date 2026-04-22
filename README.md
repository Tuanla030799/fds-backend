# FDS Backend (Spring Boot + MyBatis)

Backend đã chuyển hoàn toàn sang Java 21 với Spring Boot, MyBatis và Flyway theo chuẩn tách lớp Spring:
- `Controller` (HTTP)
- `Service` (business/use-case)
- `Repository` interface + `MyBatis...Repository` adapter (persistence)

## Công nghệ
- Java 21
- Spring Boot 3
- MyBatis XML
- Flyway
- PostgreSQL
- HikariCP

## Cấu hình
Project dùng cấu hình chuẩn Spring Boot qua:
- `src/main/resources/application.properties` (mặc định/local)
- `src/main/resources/application-docker.properties` (khi chạy Docker)

## Chạy local
1. Chạy PostgreSQL:
```bash
docker compose up -d db
```
2. Chạy API:
```bash
./gradlew bootRun
```

## Build bằng Gradle
```bash
./gradlew clean bootJar
```

## Logging
- Log ghi ra console và rolling file qua Logback.
- Local: `./logs/fds-backend.log`
- Docker/server: `/app/logs/fds-backend.log`
- Khi chạy bằng Docker Compose, thư mục này được mount ra host tại `./logs`.
- File rolling theo ngày và dung lượng: `logs/archive/fds-backend.yyyy-MM-dd.i.log.gz`
- Mức log mặc định: `root=INFO`, `com.fds.backend=INFO`, `org.springframework=INFO`, `org.mybatis=INFO`.
- Mỗi HTTP request được log dạng `METHOD /path -> status (duration ms)`.

Tail log local:
```bash
tail -f logs/fds-backend.log
```

Tail log trên server/container:
```bash
docker compose exec api tail -f /app/logs/fds-backend.log
```

Hoặc tail trực tiếp trên host server:
```bash
tail -f logs/fds-backend.log
```

## CORS
- Local mặc định cho phép frontend từ `http://localhost:3000`, `http://localhost:5173`, `http://127.0.0.1:3000`, `http://127.0.0.1:5173`.
- Docker/server cấu hình bằng biến môi trường `APP_CORS_ALLOWED_ORIGINS`, phân tách nhiều origin bằng dấu phẩy.

Ví dụ:
```bash
APP_CORS_ALLOWED_ORIGINS=https://admin.example.com,https://www.example.com docker compose up -d
```

## Chạy full bằng Docker
```bash
docker compose up --build
```
`docker-compose` set `SPRING_PROFILES_ACTIVE=docker` để dùng profile Docker.

## Endpoint chính
- `GET /api/health`
- `POST /api/admin/auth/register`
- `POST /api/admin/auth/login`
- `POST /api/admin/auth/refresh`
- `POST /api/admin/auth/logout`
- `POST /api/design-submissions`
- `GET /api/admin/design-submissions` (Bearer token)
- `PATCH /api/admin/design-submissions/{id}/status` (Bearer token)
- `DELETE /api/admin/design-submissions/{id}` (Bearer token)
- `GET /api/presets`
- `GET /api/admin/presets` (Bearer token)
- `POST /api/admin/presets` (Bearer token, body dùng `fileId` từ API upload)
- `DELETE /api/admin/presets/{id}` (Bearer token)
- `POST /api/files/upload`
- `POST /api/admin/files/upload` (Bearer token)
- `GET /files/{yyyy-MM-dd}/{file-name}`

## File upload / preset image
- Upload file trả về `fileId`, `path`, `url` và lưu file với status `INACTIVE`.
- File được lưu theo ngày dưới `app.upload-dir`, ví dụ `./uploads/2026-01-01/<file-name>`.
- DB lưu relative path, ví dụ `2026-01-01/<file-name>`.
- `url` được build từ `app.cdn.base-url`; local mặc định là `/files`, production cấu hình domain Cloudflare/CDN, ví dụ `https://cdn.example.com/files`.
- Docker nhận CDN base URL qua env `APP_CDN_BASE_URL`.
- `/api/**` trả no-cache headers.
- `/files/**` expose nội dung trong `app.upload-dir` và trả `Cache-Control: public, max-age=31536000` để Cloudflare cache mạnh.
- File upload được prefix UUID trong filename, nên mỗi upload có URL mới và tránh lỗi CDN/browser cache file cũ.
- Khi tạo preset, FE gửi `fileId`; backend lấy path file đó lưu vào `presets.image_url`, sau đó đổi file sang `ACTIVE`.
- Cron cleanup chạy định kỳ và xóa các file status `INACTIVE` có `created_at` cũ hơn 10 phút.
- `created_at` dùng default DB, `updated_at` do DB trigger tự cập nhật; `created_by/update_by` được service truyền từ admin token nếu request có token.

## Migration
- `src/main/resources/db/migration/V1__initial_schema.sql`
- `src/main/resources/db/migration/V2__seed_admin.sql`
- `src/main/resources/db/migration/V3__index_files_status_created_at.sql`
- `src/main/resources/db/migration/V4__audit_columns_and_updated_at_triggers.sql`

## Tài liệu kiến trúc
- `docs/ARCHITECTURE.md`
