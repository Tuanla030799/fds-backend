package designsubmissionmodule

import (
	"time"

	"fds-backend/internal/domain/admin"
	"fds-backend/internal/domain/designsubmission"
	"fds-backend/internal/domain/fileasset"
	"fds-backend/internal/http/handlers"
	"fds-backend/internal/http/middleware"
	"fds-backend/internal/modulekit"
)

type Module struct{}

func NewModule() *Module       { return &Module{} }
func (m *Module) Name() string { return "design-submission" }

func (m *Module) Register(reg *modulekit.Registry) error {
	repo := designsubmission.NewGormRepository(reg.DB)
	fileService := fileasset.NewService(fileasset.NewGormRepository(reg.DB), reg.Providers.Storage, 10*time.Minute)
	service := designsubmission.NewService(repo, reg.DB, fileService)
	h := handlers.NewDesignSubmissionHandler(
		service,
		reg.Providers.Storage,
		reg.Providers.AuditLog,
		reg.Config,
	)

	reg.Public.POST("/design-submissions", middleware.NewOptionalJWTMiddleware(reg.Providers.JWTManager), h.Create)
	reg.Admin.GET("/design-submissions", h.AdminList)
	reg.Admin.PATCH(
		"/design-submissions/:id/status",
		middleware.RequireRoles(admin.RoleSuperAdmin, admin.RoleOperator),
		h.AdminUpdateStatus,
	)
	reg.Admin.DELETE(
		"/design-submissions/:id",
		middleware.RequireRoles(admin.RoleSuperAdmin, admin.RoleOperator),
		h.AdminDelete,
	)

	return nil
}
