package migration

import (
	"be-cp2b/internal/domain"
	"gorm.io/gorm"
	"log"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.User{},
		&domain.Account{},
		&domain.Balance{},
		&domain.Category{},
		&domain.Transaction{},
	)

	if err != nil {
		log.Fatalf("Migrasi Database Gagal: %v", err)
	}
}
