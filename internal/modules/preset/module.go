package presetmodule

import (
	"fds-backend/internal/bootstrap"
	"fds-backend/internal/domain/admin"
	"fds-backend/internal/domain/preset"
	"fds-backend/internal/http/handlers"
	"fds-backend/internal/http/middleware"
)

type Module struct{}

func NewModule() *Module       { return &Module{} }
func (m *Module) Name() string { return "preset" }
func (m *Module) Register(reg *bootstrap.Registry) error {
	repo := preset.NewGormRepository(reg.DB)
	service := preset.NewService(repo)
	h := handlers.NewPresetHandler(service, reg.Providers.Storage, reg.Providers.AuditLog, reg.Config)
	reg.Public.GET("/presets", h.PublicList)
	reg.Admin.GET("/presets", h.PublicList)
	reg.Admin.POST("/presets", middleware.RequireRoles(admin.RoleSuperAdmin, admin.RoleOperator), middleware.RateLimit("admin-upload", reg.Config.Security.UploadRateLimitPerMinute), h.AdminCreate)
	reg.Admin.DELETE("/presets/:id", middleware.RequireRoles(admin.RoleSuperAdmin, admin.RoleOperator), h.AdminDelete)
	return nil
}
