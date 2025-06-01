package seeder

import (
	"be-cp2b/internal/domain"
	"gorm.io/gorm"
)

func SeedAccount(db *gorm.DB) {
	var account domain.Account

	account = domain.Account{
		Name:         "BCA",
		Descriptions: "Semua yang ada di BCA",
	}

	db.Create(&account)
}
