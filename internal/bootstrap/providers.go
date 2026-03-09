package bootstrap

import (
	"fds-backend/internal/config"
	"fds-backend/internal/domain/auditlog"
	appmiddleware "fds-backend/internal/http/middleware"
	"fds-backend/internal/modulekit"
	"fds-backend/internal/platform/auth"
	appLogger "fds-backend/internal/platform/logger"
	"fds-backend/internal/platform/storage"
)

func NewProviders(
	cfg *config.Config,
	jwtManager *auth.JWTManager,
	fileStorage storage.Storage,
	logger *appLogger.Logger,
	auditService *auditlog.Service,
) *modulekit.Providers {
	return &modulekit.Providers{
		JWTManager: jwtManager,
		Storage:    fileStorage,
		Logger:     logger,
		AuditLog:   auditService,
		Middleware: modulekit.MiddlewareSet{
			JWT: appmiddleware.NewJWTMiddleware(jwtManager),
		},
	}
}
