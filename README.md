# FDS Backend (Spring Boot + MyBatis)

Backend đã chuyển hoàn toàn sang Java 21 với Spring Boot, MyBatis và Flyway.

## Công nghệ
- Java 21
- Spring Boot 3
- MyBatis XML
- Flyway
- PostgreSQL
- HikariCP

## Cấu hình
Project dùng cấu hình chuẩn Spring Boot qua:
- `springboot-backend/src/main/resources/application.properties` (mặc định/local)
- `springboot-backend/src/main/resources/application-docker.properties` (khi chạy Docker)

> Không còn dùng `.env.server` / `.env.example` cho app config.

## Chạy local
1. Chạy PostgreSQL:
```bash
docker compose up -d db
```
2. Chạy API:
```bash
cd springboot-backend
mvn spring-boot:run
```

Mặc định API chạy tại: `http://localhost:8080`.

## Chạy full bằng Docker
```bash
docker compose up --build
```
Docker compose tự set `SPRING_PROFILES_ACTIVE=docker` để dùng `application-docker.properties`.

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

## CSDL migration
- `springboot-backend/src/main/resources/db/migration/V1__initial_schema.sql`
- `springboot-backend/src/main/resources/db/migration/V2__seed_admin.sql`

## Kiến trúc mở rộng
- Tài liệu: `docs/ARCHITECTURE.md`
