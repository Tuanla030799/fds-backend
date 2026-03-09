package auditlog

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Service struct{ repo Repository }

type RecordInput struct {
	AdminID, Action, Resource, ResourceID, RequestID, IP, UserAgent string
	Before, After                                                   any
}

func NewService(repo Repository) *Service { return &Service{repo: repo} }
func (s *Service) Record(input RecordInput) error {
	var adminID *uuid.UUID
	if input.AdminID != "" {
		if parsed, err := uuid.Parse(input.AdminID); err == nil {
			adminID = &parsed
		}
	}
	item := &AuditLog{AdminID: adminID, Action: input.Action, Resource: input.Resource, ResourceID: input.ResourceID, RequestID: input.RequestID, IP: input.IP, UserAgent: input.UserAgent}
	if input.Before != nil {
		if raw, err := json.Marshal(input.Before); err == nil {
			item.BeforeState = raw
		}
	}
	if input.After != nil {
		if raw, err := json.Marshal(input.After); err == nil {
			item.AfterState = raw
		}
	}
	return s.repo.Create(item)
}
