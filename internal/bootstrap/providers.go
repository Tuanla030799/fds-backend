package bootstrap

import (
	"fds-backend/internal/config"
	"fds-backend/internal/domain/auditlog"
	appmiddleware "fds-backend/internal/http/middleware"
	"fds-backend/internal/platform/auth"
	appLogger "fds-backend/internal/platform/logger"
	"fds-backend/internal/platform/storage"

	"github.com/gin-gonic/gin"
)

type Providers struct {
	JWTManager *auth.JWTManager
	Storage    storage.Storage
	Logger     *appLogger.Logger
	AuditLog   *auditlog.Service
	Middleware MiddlewareSet
}

type MiddlewareSet struct{ JWT gin.HandlerFunc }

func NewProviders(cfg *config.Config, jwtManager *auth.JWTManager, fileStorage storage.Storage, logger *appLogger.Logger, auditService *auditlog.Service) *Providers {
	return &Providers{JWTManager: jwtManager, Storage: fileStorage, Logger: logger, AuditLog: auditService, Middleware: MiddlewareSet{JWT: appmiddleware.NewJWTMiddleware(jwtManager)}}
}
