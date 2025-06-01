package seeder

import "be-cp2b/internal/config"

func SeedMain() {
	db := config.DB
	SeedUser(db)
	SeedAccount(db)
	SeedCategory(db)
}
