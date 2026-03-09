package pagination

import (
	"testing"

	"fds-backend/internal/config"
)

func TestNormalize(t *testing.T) {
	p := Normalize(0, 999, config.PaginationConfig{DefaultLimit: 20, MaxLimit: 100})
	if p.Page != 1 || p.Limit != 100 || p.Offset != 0 {
		t.Fatalf("unexpected params: %+v", p)
	}
}
