package pagination

import (
	"math"
	"strconv"
	"strings"

	"fds-backend/internal/config"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
type Meta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"totalPages"`
}

func Normalize(page, limit int, cfg config.PaginationConfig) Params {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = cfg.DefaultLimit
	}
	if limit > cfg.MaxLimit {
		limit = cfg.MaxLimit
	}
	return Params{Page: page, Limit: limit, Offset: (page - 1) * limit}
}
func FromRequest(c *gin.Context, cfg config.PaginationConfig) Params {
	page, _ := strconv.Atoi(strings.TrimSpace(c.DefaultQuery("page", "1")))
	limit, _ := strconv.Atoi(strings.TrimSpace(c.DefaultQuery("limit", strconv.Itoa(cfg.DefaultLimit))))
	return Normalize(page, limit, cfg)
}
func NewMeta(params Params, total int64) Meta {
	totalPages := 0
	if params.Limit > 0 && total > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(params.Limit)))
	}
	return Meta{Page: params.Page, Limit: params.Limit, Total: total, TotalPages: totalPages}
}
