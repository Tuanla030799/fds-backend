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
10. storage strategy abstraction (local + s3 scaffold)
11. OpenAPI/Swagger yaml
12. tests scaffold
13. seed data
14. transaction helper
15. rate limit + security headers

## Run with Docker

```bash
cp .env.example .env
docker compose up --build
```

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
- S3 đã có abstraction để mở rộng, file driver `s3.go` hiện là scaffold để nối SDK sau.
- AutoMigrate đang bật để dev nhanh; production nên đổi sang migration runner riêng.
# fds-backend
