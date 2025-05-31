package seeder

import (
	"be-cp2b/internal/domain"
	"be-cp2b/pkg/utils"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	var user domain.User
	if err := db.Where("email = ?", "irawankillmer@gmail.com").First(&user).Error; err == nil {
		return
	}

	hashed, _ := utils.HashPassword("Ir123085Aa*")
	user = domain.User{
		Name:     "Irawan Kilmer",
		Email:    "irawankillmer@gmail.com",
		Password: hashed,
	}

	db.Create(&user)
}
