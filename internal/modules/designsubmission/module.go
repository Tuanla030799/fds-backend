package designsubmissionmodule

import (
	"fds-backend/internal/domain/admin"
	"fds-backend/internal/domain/designsubmission"
	"fds-backend/internal/http/handlers"
	"fds-backend/internal/http/middleware"
	"fds-backend/internal/modulekit"
)

type Module struct{}

func NewModule() *Module       { return &Module{} }
func (m *Module) Name() string { return "design-submission" }

func (m *Module) Register(reg *modulekit.Registry) error {
	repo := designsubmission.NewGormRepository(reg.DB)
	service := designsubmission.NewService(repo)
	h := handlers.NewDesignSubmissionHandler(
		service,
		reg.Providers.Storage,
		reg.Providers.AuditLog,
		reg.Config,
	)

	reg.Public.POST("/design-submissions", h.Create)
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
