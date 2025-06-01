package router

import (
	"be-cp2b/internal"
	"be-cp2b/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, app *internal.AppContainer) {
	accountHandler := handler.NewAccountHandler(app.AccountUsecase)

	api := r.Group("/api")
	api.GET("/account", accountHandler.GetAllAccounts)
	api.POST("/account", accountHandler.CreateAccount)
	api.GET("/account/:id", accountHandler.GetAccountByID)
	api.PUT("/account/:id", accountHandler.UpdateAccount)
	api.DELETE("/account/:id", accountHandler.DeleteAccount)
}
