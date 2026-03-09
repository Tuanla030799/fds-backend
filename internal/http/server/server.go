package server

import (
	"time"

	"fds-backend/internal/config"
	"fds-backend/internal/http/middleware"
	appLogger "fds-backend/internal/platform/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Groups struct {
	Engine    *gin.Engine
	API       *gin.RouterGroup
	Admin     *gin.RouterGroup
	AdminAuth *gin.RouterGroup
	Public    *gin.RouterGroup
}

func New(cfg *config.Config, logger *appLogger.Logger, jwtMiddleware gin.HandlerFunc) *Groups {
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	_ = r.SetTrustedProxies(cfg.Security.TrustedProxies)
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID())
	r.Use(middleware.SecurityHeaders())
	r.Use(middleware.AccessLogger(logger))
	r.MaxMultipartMemory = cfg.Upload.MaxBytes
	r.Use(cors.New(cors.Config{AllowOrigins: cfg.CORS.Origins, AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}, AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "X-Request-ID"}, ExposeHeaders: []string{"Content-Length", "X-Request-ID"}, AllowCredentials: true, MaxAge: 12 * time.Hour}))
	if cfg.Storage.Driver == "local" {
		r.Static(cfg.Upload.PublicBasePath, cfg.Upload.Dir)
	}
	api := r.Group("/api")
	admin := api.Group("/admin")
	adminAuth := admin.Group("/auth")
	adminAuth.Use(middleware.RateLimit("admin-login", cfg.Security.LoginRateLimitPerMinute))
	adminProtected := api.Group("/admin")
	adminProtected.Use(jwtMiddleware)
	return &Groups{Engine: r, API: api, Admin: adminProtected, AdminAuth: adminAuth, Public: api}
}
