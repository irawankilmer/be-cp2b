package router

import (
	"be-cp2b/docs"
	"be-cp2b/internal"
	"be-cp2b/internal/handler"
	"be-cp2b/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(r *gin.Engine, app *internal.AppContainer) {
	authHandler := handler.NewAuthHandler(app.Authusecase)
	accountHandler := handler.NewAccountHandler(app.AccountUsecase)
	categoryHandler := handler.NewCategoryHandler(app.CategoryUsecase)
	transactionHandler := handler.NewTransactionHandler(app.TransactionUsecase)

	r.Use(middleware.CORSMiddleware())
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	api.POST("/login", authHandler.Login)

	auth := api.Group("")
	auth.Use(middleware.AuthMiddleware())

	auth.POST("/logout", authHandler.Logout)

	auth.GET("/account", accountHandler.GetAllAccounts)
	auth.POST("/account", accountHandler.CreateAccount)
	auth.GET("/account/:id", accountHandler.GetAccountByID)
	auth.PUT("/account/:id", accountHandler.UpdateAccount)
	auth.DELETE("/account/:id", accountHandler.DeleteAccount)

	auth.GET("/category", categoryHandler.GetAllCategories)
	auth.POST("/category", categoryHandler.CreateCategory)
	auth.GET("/category/:id", categoryHandler.GetCategoryByID)
	auth.PUT("/category/:id", categoryHandler.UpdateCategory)
	auth.DELETE("/category/:id", categoryHandler.DeleteCategory)

	auth.GET("/transaction", transactionHandler.GetAllTransactions)
	auth.POST("/transaction", transactionHandler.CreateTransaction)
	auth.GET("/transaction/:id", transactionHandler.GetTransactionByID)
	auth.PUT("/transaction/:id", transactionHandler.UpdateTransaction)
	auth.DELETE("/transaction/:id", transactionHandler.DeleteTransaction)
}
