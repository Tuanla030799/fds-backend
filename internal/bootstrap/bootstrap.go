package bootstrap

import (
	"fmt"

	"fds-backend/internal/app"
	"fds-backend/internal/config"
	"fds-backend/internal/database"
	"fds-backend/internal/domain/auditlog"
	"fds-backend/internal/http/server"
	"fds-backend/internal/modulekit"
	"fds-backend/internal/platform/auth"
	appLogger "fds-backend/internal/platform/logger"
	"fds-backend/internal/platform/storage"
)

func BuildApplication() (*app.Application, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	logger := appLogger.New(cfg.App.Env)

	var fileStorage storage.Storage = storage.NewLocalStorage(cfg.Upload.Dir, cfg.Upload.PublicBasePath)

	if err := fileStorage.EnsureDirs(cfg.Upload.OrdersSubdir, cfg.Upload.PresetsSubdir, cfg.Upload.TempSubdir); err != nil {
		return nil, err
	}

	db, err := database.Connect(cfg, logger)
	if err != nil {
		return nil, err
	}

	if cfg.Database.AutoMigrate {
		if err := database.Migrate(db); err != nil {
			return nil, err
		}
	}

	if err := database.Seed(db, cfg, logger); err != nil {
		return nil, err
	}

	jwtManager := auth.NewJWTManager(cfg.JWT.Secret)
	auditService := auditlog.NewService(auditlog.NewGormRepository(db))
	providers := NewProviders(cfg, jwtManager, fileStorage, logger, auditService)
	groups := server.New(cfg, providers.Logger, providers.Middleware.JWT)

	registry := &modulekit.Registry{
		Config:    cfg,
		DB:        db,
		Router:    groups.API,
		Public:    groups.Public,
		Admin:     groups.Admin,
		AdminAuth: groups.AdminAuth,
		Providers: providers,
	}

	for _, module := range DefaultModules() {
		if err := module.Register(registry); err != nil {
			return nil, fmt.Errorf("register module %s: %w", module.Name(), err)
		}
	}

	return &app.Application{
		Config: cfg,
		Router: groups.Engine,
	}, nil
}
