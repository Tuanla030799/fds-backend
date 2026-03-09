package response

import "github.com/gin-gonic/gin"

type Envelope struct {
	Data    any            `json:"data,omitempty"`
	Message string         `json:"message,omitempty"`
	Meta    map[string]any `json:"meta,omitempty"`
}

func JSON(c *gin.Context, status int, data any, message string) {
	c.JSON(status, Envelope{Data: data, Message: message})
}

func Error(c *gin.Context, status int, message string, meta map[string]any) {
	c.JSON(status, Envelope{Message: message, Meta: meta})
}
