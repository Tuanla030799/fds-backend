package fileassetmodule

import (
	"time"

	"fds-backend/internal/domain/fileasset"
	"fds-backend/internal/http/handlers"
	"fds-backend/internal/http/middleware"
	"fds-backend/internal/modulekit"
)

type Module struct{}

func NewModule() *Module       { return &Module{} }
func (m *Module) Name() string { return "file-asset" }

func (m *Module) Register(reg *modulekit.Registry) error {
	repo := fileasset.NewGormRepository(reg.DB)
	service := fileasset.NewService(repo, reg.Providers.Storage, 10*time.Minute)
	service.StartCleanupJob()

	h := handlers.NewFileHandler(service, reg.Config)
	reg.Router.POST(
		"/files/upload",
		middleware.NewOptionalJWTMiddleware(reg.Providers.JWTManager),
		middleware.UploadRateLimit("files-upload", 20),
		h.Upload,
	)
	return nil
}
