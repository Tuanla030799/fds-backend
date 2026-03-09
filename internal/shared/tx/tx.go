package tx

import "gorm.io/gorm"

func Within(db *gorm.DB, fn func(tx *gorm.DB) error) error {
	return db.Transaction(func(txx *gorm.DB) error { return fn(txx) })
}
