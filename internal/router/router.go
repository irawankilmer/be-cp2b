package router

import (
	"be-cp2b/internal"
	"be-cp2b/internal/handler"
	"be-cp2b/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, app *internal.AppContainer) {
	authHandler := handler.NewAuthHandler(app.Authusecase)
	accountHandler := handler.NewAccountHandler(app.AccountUsecase)
	categoryHandler := handler.NewCategoryHandler(app.CategoryUsecase)

	r.Use(middleware.CORSMiddleware())
	api := r.Group("/api")
	api.POST("/login", authHandler.Login)

	api.GET("/account", accountHandler.GetAllAccounts)
	api.POST("/account", accountHandler.CreateAccount)
	api.GET("/account/:id", accountHandler.GetAccountByID)
	api.PUT("/account/:id", accountHandler.UpdateAccount)
	api.DELETE("/account/:id", accountHandler.DeleteAccount)

	api.GET("/category", categoryHandler.GetAllCategories)
	api.POST("/category", categoryHandler.CreateCategory)
	api.GET("/category/:id", categoryHandler.GetCategoryByID)
	api.PUT("/category/:id", categoryHandler.UpdateCategory)
	api.DELETE("/category/:id", categoryHandler.DeleteCategory)
}
