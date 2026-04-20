package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	App        AppConfig
	Database   DatabaseConfig
	JWT        JWTConfig
	CORS       CORSConfig
	Upload     UploadConfig
	Pagination PaginationConfig
	Security   SecurityConfig
	Seed       SeedConfig
}

type AppConfig struct {
	Env     string
	Port    string
	BaseURL string
	Name    string
}

type DatabaseConfig struct {
	URL             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	AutoMigrate     bool
}

type JWTConfig struct {
	Secret               string
	AccessTokenTTL       time.Duration
	RefreshTokenTTL      time.Duration
	RefreshTokenAudience string
}

type CORSConfig struct{ Origins []string }

type UploadConfig struct {
	Dir            string
	MaxBytes       int64
	AllowedMime    []string
	AllowedExts    []string
	PresetsSubdir  string
	OrdersSubdir   string
	TempSubdir     string
	PublicBasePath string
}

type PaginationConfig struct {
	DefaultLimit int
	MaxLimit     int
}

type SecurityConfig struct {
	LoginRateLimitPerMinute  int
	UploadRateLimitPerMinute int
	TrustedProxies           []string
}

type SeedConfig struct {
	AdminEmail    string
	AdminPassword string
	AdminName     string
	CreateDemo    bool
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	cfg := &Config{
		App:        AppConfig{Env: getEnv("APP_ENV", "development"), Port: getEnv("PORT", "8080"), BaseURL: getEnv("BASE_URL", "http://localhost:8080"), Name: getEnv("APP_NAME", "fds-backend")},
		Database:   DatabaseConfig{URL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/fds?sslmode=disable"), MaxIdleConns: getEnvInt("DB_MAX_IDLE_CONNS", 10), MaxOpenConns: getEnvInt("DB_MAX_OPEN_CONNS", 50), ConnMaxLifetime: time.Duration(getEnvInt("DB_CONN_MAX_LIFETIME_MINUTES", 30)) * time.Minute, AutoMigrate: getEnvBool("DB_AUTO_MIGRATE", true)},
		JWT:        JWTConfig{Secret: getEnv("JWT_SECRET", "change-me-now"), AccessTokenTTL: time.Duration(getEnvInt("JWT_ACCESS_TTL_HOURS", 24)) * time.Hour, RefreshTokenTTL: time.Duration(getEnvInt("JWT_REFRESH_TTL_HOURS", 24*30)) * time.Hour, RefreshTokenAudience: getEnv("JWT_REFRESH_AUDIENCE", "admin-refresh")},
		CORS:       CORSConfig{Origins: parseCSV(getEnv("CORS_ORIGINS", "http://localhost:5173,http://localhost:3000"))},
		Upload:     UploadConfig{Dir: getEnv("UPLOADS_DIR", "./uploads"), MaxBytes: int64(getEnvInt("UPLOAD_MAX_BYTES", 20<<20)), AllowedMime: parseCSV(getEnv("UPLOAD_ALLOWED_MIME", "image/png,image/jpeg,image/webp")), AllowedExts: parseCSV(getEnv("UPLOAD_ALLOWED_EXTS", ".png,.jpg,.jpeg,.webp")), PresetsSubdir: getEnv("UPLOAD_PRESETS_SUBDIR", "presets"), OrdersSubdir: getEnv("UPLOAD_ORDERS_SUBDIR", "design-submissions"), TempSubdir: getEnv("UPLOAD_TEMP_SUBDIR", "tmp"), PublicBasePath: getEnv("UPLOAD_PUBLIC_BASE_PATH", "/uploads")},
		Pagination: PaginationConfig{DefaultLimit: getEnvInt("PAGINATION_DEFAULT_LIMIT", 20), MaxLimit: getEnvInt("PAGINATION_MAX_LIMIT", 100)},
		Security:   SecurityConfig{LoginRateLimitPerMinute: getEnvInt("RATE_LIMIT_LOGIN_PER_MINUTE", 10), UploadRateLimitPerMinute: getEnvInt("RATE_LIMIT_UPLOAD_PER_MINUTE", 30), TrustedProxies: parseCSV(getEnv("TRUSTED_PROXIES", ""))},
		Seed:       SeedConfig{AdminEmail: getEnv("ADMIN_SEED_EMAIL", "admin@fds.local"), AdminPassword: getEnv("ADMIN_SEED_PASSWORD", "12345678"), AdminName: getEnv("ADMIN_SEED_NAME", "System Admin"), CreateDemo: getEnvBool("SEED_DEMO_DATA", true)},
	}
	return cfg, cfg.Validate()
}

func (c *Config) Validate() error {
	var issues []string
	if c.JWT.Secret == "" || c.JWT.Secret == "change-me-now" {
		issues = append(issues, "JWT_SECRET must be set to a non-default value")
	}
	if c.Database.URL == "" {
		issues = append(issues, "DATABASE_URL is required")
	}
	if c.Pagination.DefaultLimit <= 0 || c.Pagination.MaxLimit <= 0 || c.Pagination.DefaultLimit > c.Pagination.MaxLimit {
		issues = append(issues, "pagination config is invalid")
	}
	if c.Upload.MaxBytes <= 0 {
		issues = append(issues, "UPLOAD_MAX_BYTES must be positive")
	}
	if len(issues) > 0 {
		return fmt.Errorf("invalid config: %s", strings.Join(issues, "; "))
	}
	return nil
}

func getEnv(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}
func getEnvInt(key string, fallback int) int {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}
func getEnvBool(key string, fallback bool) bool {
	v := strings.ToLower(strings.TrimSpace(os.Getenv(key)))
	if v == "" {
		return fallback
	}
	return v == "1" || v == "true" || v == "yes"
}
func parseCSV(value string) []string {
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
