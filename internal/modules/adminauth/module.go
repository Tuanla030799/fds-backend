package adminauth

import (
	"fds-backend/internal/domain/admin"
	"fds-backend/internal/domain/refreshtoken"
	"fds-backend/internal/http/handlers"
	"fds-backend/internal/http/middleware"
	"fds-backend/internal/modulekit"
)

type Module struct{}

func NewModule() *Module       { return &Module{} }
func (m *Module) Name() string { return "admin-auth" }

func (m *Module) Register(reg *modulekit.Registry) error {
	repo := admin.NewGormRepository(reg.DB)
	refreshRepo := refreshtoken.NewGormRepository(reg.DB)
	refreshService := refreshtoken.NewService(refreshRepo, reg.Config.JWT.RefreshTokenTTL)
	service := admin.NewAuthService(
		repo,
		reg.Providers.JWTManager,
		refreshService,
		reg.DB,
		reg.Config.JWT.AccessTokenTTL,
	)

	h := handlers.NewAuthHandler(service, reg.Providers.AuditLog)

	reg.AdminAuth.POST("/register", h.Register)
	reg.AdminAuth.POST("/login", h.Login)
	reg.AdminAuth.POST("/refresh", h.Refresh)

	reg.Admin.Use(middleware.RequireRoles(admin.RoleSuperAdmin, admin.RoleOperator, admin.RoleViewer))
	reg.Admin.POST("/auth/logout", h.Logout)

	return nil
}
