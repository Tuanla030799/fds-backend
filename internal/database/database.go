package database

import (
	"fmt"
	"time"

	"fds-backend/internal/config"
	"fds-backend/internal/domain/admin"
	"fds-backend/internal/domain/auditlog"
	"fds-backend/internal/domain/designsubmission"
	"fds-backend/internal/domain/fileasset"
	"fds-backend/internal/domain/preset"
	"fds-backend/internal/domain/refreshtoken"
	appLogger "fds-backend/internal/platform/logger"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func Connect(cfg *config.Config, logger *appLogger.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.Database.URL), &gorm.Config{Logger: gormlogger.Default.LogMode(gormlogger.Warn)})
	if err != nil {
		return nil, fmt.Errorf("connect database: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime)
	logger.Info("database connected")
	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&admin.User{}, &designsubmission.Submission{}, &preset.Preset{}, &auditlog.AuditLog{}, &refreshtoken.Token{}, &fileasset.File{})
}

func Seed(db *gorm.DB, cfg *config.Config, logger *appLogger.Logger) error {
	var count int64
	if err := db.Model(&admin.User{}).Where("email = ?", cfg.Seed.AdminEmail).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(cfg.Seed.AdminPassword), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		seed := admin.User{Name: cfg.Seed.AdminName, Email: cfg.Seed.AdminEmail, PasswordHash: string(hash), Role: admin.RoleSuperAdmin}
		if err := db.Create(&seed).Error; err != nil {
			return err
		}
		logger.Info("seeded admin user", "email", cfg.Seed.AdminEmail)
	}
	if !cfg.Seed.CreateDemo {
		return nil
	}
	if err := seedPresets(db); err != nil {
		return err
	}
	if err := seedSubmissions(db); err != nil {
		return err
	}
	return nil
}

func seedPresets(db *gorm.DB) error {
	var count int64
	if err := db.Model(&preset.Preset{}).Count(&count).Error; err != nil || count > 0 {
		return err
	}
	items := []preset.Preset{
		{Name: "Classic Varsity", Status: preset.StatusActive, Note: "Demo preset", Tags: datatypes.JSON([]byte(`["cap","satin"]`)), ImageURL: "/uploads/presets/demo-1.png", SortOrder: 1},
		{Name: "Street Puff", Status: preset.StatusDraft, Note: "Demo preset", Tags: datatypes.JSON([]byte(`["hoodie"]`)), ImageURL: "/uploads/presets/demo-2.png", SortOrder: 2},
	}
	return db.Create(&items).Error
}

func seedSubmissions(db *gorm.DB) error {
	var count int64
	if err := db.Model(&designsubmission.Submission{}).Count(&count).Error; err != nil || count > 0 {
		return err
	}
	items := []designsubmission.Submission{{FullName: "Nguyen Van A", Address: "HCM", Phone: "0900000001", Note: "Demo order", ImageURL: "/uploads/design-submissions/demo-1.png", Status: designsubmission.StatusPending, CreatedAt: time.Now()}, {FullName: "Tran Thi B", Address: "HN", Phone: "0900000002", Note: "Urgent", ImageURL: "/uploads/design-submissions/demo-2.png", Status: designsubmission.StatusConfirmed, CreatedAt: time.Now()}}
	return db.Create(&items).Error
}
