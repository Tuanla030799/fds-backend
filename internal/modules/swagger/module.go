package swagger

import (
	"net/http"

	"fds-backend/internal/modulekit"

	"github.com/gin-gonic/gin"
)

type Module struct{}

func NewModule() *Module       { return &Module{} }
func (m *Module) Name() string { return "swagger" }

func (m *Module) Register(reg *modulekit.Registry) error {
	reg.Public.GET("/docs/openapi.yaml", func(c *gin.Context) {
		c.File("./internal/http/docs/openapi.yaml")
	})
	reg.Public.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/api/docs/openapi.yaml")
	})
	return nil
}
