package services

import (
	"errors"
	"time"

	"fds-backend/internal/config"
	"fds-backend/internal/models"
	"fds-backend/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	cfg  *config.Config
	repo *repositories.AdminRepository
}

type AuthResult struct {
	AccessToken string `json:"accessToken"`
	Admin       any    `json:"admin"`
}

func NewAuthService(cfg *config.Config, repo *repositories.AdminRepository) *AuthService {
	return &AuthService{cfg: cfg, repo: repo}
}

func (s *AuthService) Register(name, email, password string) (*AuthResult, error) {
	_, err := s.repo.FindByEmail(email)
	if err == nil {
		return nil, errors.New("email already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	admin := &models.AdminUser{Name: name, Email: email, PasswordHash: string(hash)}
	if err := s.repo.Create(admin); err != nil {
		return nil, err
	}
	return s.buildAuthResult(admin)
}

func (s *AuthService) Login(email, password string) (*AuthResult, error) {
	admin, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}
	return s.buildAuthResult(admin)
}

func (s *AuthService) buildAuthResult(admin *models.AdminUser) (*AuthResult, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"sub":   admin.ID.String(),
		"email": admin.Email,
		"name":  admin.Name,
		"role":  "admin",
		"iat":   now.Unix(),
		"exp":   now.Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return nil, err
	}
	return &AuthResult{AccessToken: signed, Admin: map[string]any{"id": admin.ID, "name": admin.Name, "email": admin.Email}}, nil
}
