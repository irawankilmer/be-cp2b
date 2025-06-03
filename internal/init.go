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
	Authusecase        usecase.AuthUsecase
	AccountUsecase     usecase.AccountUsecase
	CategoryUsecase    usecase.CategoryUsecase
	TransactionUsecase usecase.TransactionUsecase
}

func InitApp() *AppContainer {
	config.LoadEnv()
	config.InitDB()
	db := config.DB

	if os.Getenv("APP_ENV") == "local" {
		migration.AutoMigrate(db)
		seeder.SeedMain()
	}

	authRepo := repository.NewAuthRepository(db)
	accountRepo := repository.NewAccountRepository(db)
	balanceRepo := repository.NewBalanceRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	return &AppContainer{
		Authusecase:        usecase.NewAuthUsecase(authRepo),
		AccountUsecase:     usecase.NewAccountUsecase(accountRepo, balanceRepo),
		CategoryUsecase:    usecase.NewCategoryUsecase(categoryRepo),
		TransactionUsecase: usecase.NewTransactionUsecase(transactionRepo),
	}
}
