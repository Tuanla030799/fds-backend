package refreshtoken

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"fds-backend/internal/shared/apperrors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repo Repository
	ttl  time.Duration
}

func NewService(repo Repository, ttl time.Duration) *Service { return &Service{repo: repo, ttl: ttl} }
func Hash(raw string) string                                 { sum := sha256.Sum256([]byte(raw)); return hex.EncodeToString(sum[:]) }
func (s *Service) Issue(adminID string) (string, *Token, error) {
	uid, err := uuid.Parse(adminID)
	if err != nil {
		return "", nil, apperrors.BadRequest("invalid admin id")
	}
	raw := uuid.NewString() + "." + uuid.NewString()
	item := &Token{AdminID: uid, TokenHash: Hash(raw), ExpiresAt: time.Now().Add(s.ttl)}
	return raw, item, s.repo.Create(item)
}
func (s *Service) Consume(raw string) (*Token, error) {
	item, err := s.repo.FindByHash(Hash(raw))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperrors.Unauthorized("invalid refresh token")
		}
		return nil, err
	}
	if item.RevokedAt != nil || time.Now().After(item.ExpiresAt) {
		return nil, apperrors.Unauthorized("refresh token expired or revoked")
	}
	now := time.Now()
	item.RevokedAt = &now
	return item, s.repo.Update(item)
}
func (s *Service) RevokeAllByAdminID(adminID string) error { return s.repo.RevokeAllByAdminID(adminID) }
