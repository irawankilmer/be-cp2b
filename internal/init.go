package internal

import (
	"be-cp2b/internal/config"
	"be-cp2b/internal/migration"
	"be-cp2b/internal/repository"
	"be-cp2b/internal/seeder"
	"be-cp2b/internal/usecase"
	"os"
)

type AppContainer struct {
	AccountUsecase usecase.AccountUsecase
}

func InitApp() *AppContainer {
	config.LoadEnv()
	config.InitDB()
	db := config.DB

	if os.Getenv("APP_ENV") == "local" {
		migration.AutoMigrate(db)
		seeder.SeedMain()
	}

	accountRepo := repository.NewAccountRepository(db)
	balanceRepo := repository.NewBalanceRepository(db)

	return &AppContainer{
		AccountUsecase: usecase.NewAccountUsecase(accountRepo, balanceRepo),
	}
}
