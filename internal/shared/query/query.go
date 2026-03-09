package query

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Sort struct {
	Field string `json:"field"`
	Order string `json:"order"`
}
type Filters struct {
	Keyword string `json:"keyword,omitempty"`
	Status  string `json:"status,omitempty"`
}

func ReadFilters(c *gin.Context) Filters {
	return Filters{Keyword: strings.TrimSpace(c.Query("keyword")), Status: strings.TrimSpace(c.Query("status"))}
}
func ReadSort(c *gin.Context, defaultField string, allowed map[string]string) Sort {
	field := strings.TrimSpace(c.DefaultQuery("sort", defaultField))
	if mapped, ok := allowed[field]; ok {
		field = mapped
	} else {
		field = allowed[defaultField]
	}
	order := strings.ToLower(strings.TrimSpace(c.DefaultQuery("order", "desc")))
	if order != "asc" {
		order = "desc"
	}
	return Sort{Field: field, Order: order}
}
func Meta(filters Filters, sort Sort) map[string]any {
	return map[string]any{"filters": filters, "sort": sort}
}
