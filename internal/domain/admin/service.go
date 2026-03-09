package admin

import (
	"errors"
	"strings"
	"time"

	"fds-backend/internal/domain/refreshtoken"
	"fds-backend/internal/platform/auth"
	"fds-backend/internal/shared/apperrors"
	"fds-backend/internal/shared/tx"
	"fds-backend/internal/shared/validation"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repo       Repository
	jwtManager *auth.JWTManager
	refresh    *refreshtoken.Service
	db         *gorm.DB
	accessTTL  time.Duration
}

type AuthResult struct {
	AccessToken  string         `json:"accessToken"`
	RefreshToken string         `json:"refreshToken,omitempty"`
	Admin        map[string]any `json:"admin"`
}

func NewAuthService(repo Repository, jwtManager *auth.JWTManager, refresh *refreshtoken.Service, db *gorm.DB, accessTTL time.Duration) *AuthService {
	return &AuthService{repo: repo, jwtManager: jwtManager, refresh: refresh, db: db, accessTTL: accessTTL}
}

func (s *AuthService) Register(name, email, password string, role Role) (*AuthResult, error) {
	if err := validation.Required(name, "name"); err != nil {
		return nil, err
	}
	if err := validation.Email(email); err != nil {
		return nil, err
	}
	if err := validation.Password(password, 8); err != nil {
		return nil, err
	}
	if !role.IsValid() {
		role = RoleSuperAdmin
	}
	_, err := s.repo.FindByEmail(strings.TrimSpace(email))
	if err == nil {
		return nil, apperrors.Conflict("email already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	admin := &User{Name: strings.TrimSpace(name), Email: strings.TrimSpace(email), PasswordHash: string(hash), Role: role}
	if err := s.repo.Create(admin); err != nil {
		return nil, err
	}
	return s.buildAuthResult(admin)
}

func (s *AuthService) Login(email, password string) (*AuthResult, error) {
	if err := validation.Email(email); err != nil {
		return nil, apperrors.Unauthorized("invalid email or password")
	}
	admin, err := s.repo.FindByEmail(strings.TrimSpace(email))
	if err != nil {
		return nil, apperrors.Unauthorized("invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(password)); err != nil {
		return nil, apperrors.Unauthorized("invalid email or password")
	}
	return s.buildAuthResult(admin)
}

func (s *AuthService) Refresh(raw string) (*AuthResult, error) {
	item, err := s.refresh.Consume(raw)
	if err != nil {
		return nil, err
	}
	admin, err := s.repo.FindByID(item.AdminID.String())
	if err != nil {
		return nil, apperrors.Unauthorized("admin not found")
	}
	return s.buildAuthResult(admin)
}

func (s *AuthService) Logout(adminID string) error { return s.refresh.RevokeAllByAdminID(adminID) }

func (s *AuthService) buildAuthResult(admin *User) (*AuthResult, error) {
	var result *AuthResult
	err := tx.Within(s.db, func(_ *gorm.DB) error {
		accessToken, err := s.jwtManager.Sign(auth.Claims{Subject: admin.ID.String(), Email: admin.Email, Name: admin.Name, Role: string(admin.Role), Expiry: time.Now().Add(s.accessTTL)})
		if err != nil {
			return err
		}
		refreshToken, _, err := s.refresh.Issue(admin.ID.String())
		if err != nil {
			return err
		}
		result = &AuthResult{AccessToken: accessToken, RefreshToken: refreshToken, Admin: map[string]any{"id": admin.ID, "name": admin.Name, "email": admin.Email, "role": admin.Role}}
		return nil
	})
	return result, err
}
