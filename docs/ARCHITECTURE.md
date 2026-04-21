# Architecture (Extensible-first)

Mục tiêu: giữ codebase dễ mở rộng khi thêm module/domain mới.

## Nguyên tắc
1. **Controller chỉ nhận/trả HTTP**, không gọi MyBatis trực tiếp.
2. **Service chứa nghiệp vụ**, orchestration cho từng use-case.
3. **Repository là abstraction** để có thể đổi MyBatis sang JPA/Redis/External API mà không đổi service.
4. **Infrastructure adapter** cài đặt repository/storage theo công nghệ cụ thể.

## Module hiện tại
- `auth`: `AuthController -> AuthService -> AuthMapper`
- `preset`: `PresetController -> PresetService -> PresetRepository -> MyBatisPresetRepository -> PresetMapper`
- `design`: `DesignSubmissionController -> DesignSubmissionService -> DesignSubmissionRepository -> MyBatisDesignSubmissionRepository -> DesignSubmissionMapper`
- `file`: `FileController -> FileService -> FileAssetRepository + FileStorage`

## Mở rộng module mới
Khi thêm module mới (vd: `orders`), tạo theo skeleton:
1. `orders/OrderController.java`
2. `orders/OrderService.java`
3. `orders/OrderRepository.java` (interface)
4. `orders/MyBatisOrderRepository.java` (adapter)
5. `orders/OrderMapper.java` + `resources/mappers/OrderMapper.xml`
6. Migration Flyway tương ứng.

## Cross-cutting
- Exception envelope chung ở `common/*`.
- Config runtime chung ở `application.yml`.
- Schema thay đổi qua Flyway để đảm bảo deploy repeatable.
