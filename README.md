# FDS Backend (Java)

Project đã chuyển toàn bộ backend sang **Spring Boot + MyBatis + PostgreSQL** theo hướng thiết kế dễ mở rộng.

## Stack
- Java 21
- Spring Boot 3
- MyBatis XML
- Flyway migration
- HikariCP tuning
- JWT auth (access/refresh)

## Chạy local
```bash
cp .env.server .env
docker compose up -d db
cd springboot-backend
mvn spring-boot:run
```

## Endpoint chính
- `GET /api/health`
- `POST /api/admin/auth/register`
- `POST /api/admin/auth/login`
- `POST /api/admin/auth/refresh`
- `POST /api/admin/auth/logout`
- `POST /api/design-submissions`
- `GET /api/admin/design-submissions`
- `PATCH /api/admin/design-submissions/{id}/status`
- `DELETE /api/admin/design-submissions/{id}`
- `GET /api/presets`
- `GET /api/admin/presets`
- `POST /api/admin/presets`
- `DELETE /api/admin/presets/{id}`
- `POST /api/files/upload`

## Kiến trúc mở rộng
Xem chi tiết tại `docs/ARCHITECTURE.md`.

## DB migration
Flyway scripts nằm ở:
- `springboot-backend/src/main/resources/db/migration/V1__initial_schema.sql`
- `springboot-backend/src/main/resources/db/migration/V2__seed_admin.sql`

Tài khoản seed mặc định:
- email: `admin@fds.local`
- password: `admin123`
