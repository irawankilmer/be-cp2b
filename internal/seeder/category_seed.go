package seeder

import (
	"be-cp2b/internal/domain"
	"gorm.io/gorm"
)

func SeedCategory(db *gorm.DB) {
	var category domain.Category

	category = domain.Category{
		Name:         "Kebutuhan rumah tangga",
		Type:         "pengeluaran",
		Descriptions: "Semua kebutuhan rumah tangga ada disini",
	}

	db.Create(&category)
}
