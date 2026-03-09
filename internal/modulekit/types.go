package modulekit

import (
	"fds-backend/internal/config"
	"fds-backend/internal/domain/auditlog"
	"fds-backend/internal/platform/auth"
	appLogger "fds-backend/internal/platform/logger"
	"fds-backend/internal/platform/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MiddlewareSet struct {
	JWT gin.HandlerFunc
}

type Providers struct {
	JWTManager *auth.JWTManager
	Storage    storage.Storage
	Logger     *appLogger.Logger
	AuditLog   *auditlog.Service
	Middleware MiddlewareSet
}

type Module interface {
	Name() string
	Register(*Registry) error
}

type Registry struct {
	Config    *config.Config
	DB        *gorm.DB
	Router    *gin.RouterGroup
	Public    *gin.RouterGroup
	Admin     *gin.RouterGroup
	AdminAuth *gin.RouterGroup
	Providers *Providers
}
