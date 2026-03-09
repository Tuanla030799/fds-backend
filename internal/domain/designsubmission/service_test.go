package designsubmission

import (
	"testing"

	"fds-backend/internal/shared/pagination"
	"fds-backend/internal/shared/query"
)

type stubRepo struct{}

func (s stubRepo) Create(item *Submission) error { return nil }
func (s stubRepo) List(filters query.Filters, sort query.Sort, params pagination.Params) ([]Submission, int64, error) {
	return nil, 0, nil
}
func (s stubRepo) FindByID(id string) (*Submission, error) { return nil, nil }
func (s stubRepo) Update(item *Submission) error           { return nil }
func (s stubRepo) Delete(id string) error                  { return nil }

func TestStatusIsValid(t *testing.T) {
	if !StatusPending.IsValid() {
		t.Fatal("pending should be valid")
	}
	if Status("x").IsValid() {
		t.Fatal("x should be invalid")
	}
}
