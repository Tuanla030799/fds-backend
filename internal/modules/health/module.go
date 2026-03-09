package health

import (
	"fds-backend/internal/http/handlers"
	"fds-backend/internal/modulekit"
)

type Module struct{}

func NewModule() *Module { return &Module{} }

func (m *Module) Name() string { return "health" }

func (m *Module) Register(reg *modulekit.Registry) error {
	h := handlers.NewHealthHandler()
	reg.Public.GET("/health", h.Check)
	return nil
}
