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
- `springboot-backend/src/main/resources/application.properties` (mặc định/local)
- `springboot-backend/src/main/resources/application-docker.properties` (khi chạy Docker)

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
- `GET /api/admin/design-submissions`
- `PATCH /api/admin/design-submissions/{id}/status`
- `DELETE /api/admin/design-submissions/{id}`
- `GET /api/presets`
- `GET /api/admin/presets`
- `POST /api/admin/presets`
- `DELETE /api/admin/presets/{id}`
- `POST /api/files/upload`

## Migration
- `springboot-backend/src/main/resources/db/migration/V1__initial_schema.sql`
- `springboot-backend/src/main/resources/db/migration/V2__seed_admin.sql`

## Tài liệu kiến trúc
- `docs/ARCHITECTURE.md`
