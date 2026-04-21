# FDS Backend Modular

Backend Golang cho FDS theo hướng mở rộng module, gồm 15 nhóm nền tảng đã được thêm:
1. validation chung
2. config typed + fail-fast validate
3. migration/autormigrate scaffold
4. structured logger
5. audit log admin actions
6. refresh token
7. RBAC role-based access
8. soft delete
9. search/filter/sort chuẩn
10. storage abstraction
11. OpenAPI/Swagger yaml
12. tests scaffold
13. seed data
14. transaction helper
15. rate limit + security headers

## Run with Docker

```bash
cp .env.server .env
docker compose up --build
```

## Local Development

Chạy riêng database bằng Docker:

```bash
docker compose up -d db
```

Chạy API trên máy bằng `go run`:

```bash
go run ./cmd/api
```

Quy ước file:
- `.env`: bản local để chạy `go run`, dùng `DATABASE_URL=...@localhost...`
- `.env.server`: bản cho Docker/server, dùng `DATABASE_URL=...@db...`

## Files Upload

- `POST /api/files/upload`
- `multipart/form-data` gồm `file` và `folder`
- `folder` cho phép: `tmp`, `presets`, `design-submissions`
- file mới tạo record trong bảng `files` với status `UNACTIVE`
- mỗi user chỉ upload tối đa `20` ảnh/phút, riêng `super_admin` được bỏ qua limit
- khi tạo preset hoặc design submission với `fileId`, backend sẽ đổi status file sang `ACTIVE`
- cron job nền xóa file `UNACTIVE` quá `10` phút kể từ `created_at`

## Main endpoints

- `POST /api/admin/auth/register`
- `POST /api/admin/auth/login`
- `POST /api/admin/auth/refresh`
- `POST /api/admin/auth/logout`
- `POST /api/design-submissions`
- `GET /api/admin/design-submissions`
- `PATCH /api/admin/design-submissions/:id/status`
- `DELETE /api/admin/design-submissions/:id`
- `GET /api/presets`
- `GET /api/admin/presets`
- `POST /api/admin/presets`
- `DELETE /api/admin/presets/:id`
- `GET /api/docs/openapi.yaml`

## Roles
- `super_admin`: full access
- `operator`: manage orders and presets
- `viewer`: read only

## Notes
- Local storage chạy ngay được.
- File hiện được lưu local qua `Storage` abstraction để sau này vẫn có thể thay backend lưu trữ khác nếu cần.
- AutoMigrate đang bật để dev nhanh; production nên đổi sang migration runner riêng.
# fds-backend

## Spring Boot + MyBatis migration scaffold

Đã thêm skeleton tại `springboot-backend/` để bắt đầu chuyển đổi dần từ Go sang Spring Boot + MyBatis.

- Tài liệu kế hoạch: `docs/springboot-mybatis-migration.md`
- Endpoint mẫu: `GET /api/v2/presets`
- SQL động + pagination + `COUNT(*) OVER()` đã có sẵn trong mapper XML.
