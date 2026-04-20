package presetmodule

import (
	"time"

	"fds-backend/internal/domain/admin"
	"fds-backend/internal/domain/fileasset"
	"fds-backend/internal/domain/preset"
	"fds-backend/internal/http/handlers"
	"fds-backend/internal/http/middleware"
	"fds-backend/internal/modulekit"
)

type Module struct{}

func NewModule() *Module       { return &Module{} }
func (m *Module) Name() string { return "preset" }

func (m *Module) Register(reg *modulekit.Registry) error {
	repo := preset.NewGormRepository(reg.DB)
	fileService := fileasset.NewService(fileasset.NewGormRepository(reg.DB), reg.Providers.Storage, 10*time.Minute)
	service := preset.NewService(repo, reg.DB, fileService)
	h := handlers.NewPresetHandler(
		service,
		reg.Providers.Storage,
		reg.Providers.AuditLog,
		reg.Config,
	)

	reg.Public.GET("/presets", h.PublicList)
	reg.Admin.GET("/presets", h.PublicList)
	reg.Admin.POST(
		"/presets",
		middleware.RequireRoles(admin.RoleSuperAdmin, admin.RoleOperator),
		middleware.RateLimit("admin-upload", reg.Config.Security.UploadRateLimitPerMinute),
		h.AdminCreate,
	)
	reg.Admin.DELETE(
		"/presets/:id",
		middleware.RequireRoles(admin.RoleSuperAdmin, admin.RoleOperator),
		h.AdminDelete,
	)

	return nil
}
