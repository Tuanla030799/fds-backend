package bootstrap

import (
	"fds-backend/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
