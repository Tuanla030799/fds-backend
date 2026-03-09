package health

import (
	"fds-backend/internal/bootstrap"
	"fds-backend/internal/http/handlers"
)

type Module struct{}

func NewModule() *Module { return &Module{} }

func (m *Module) Name() string { return "health" }

func (m *Module) Register(reg *bootstrap.Registry) error {
	h := handlers.NewHealthHandler()
	reg.Public.GET("/health", h.Check)
	return nil
}
