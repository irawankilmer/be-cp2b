package internal

import (
	"be-cp2b/internal/config"
	"be-cp2b/internal/migration"
	"be-cp2b/internal/seeder"
	"os"
)

func InitApp() {
	config.LoadEnv()
	config.InitDB()
	db := config.DB

	if os.Getenv("APP_ENV") == "local" {
		migration.AutoMigrate(db)
		seeder.SeedMain()
	}
}
