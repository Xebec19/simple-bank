package routes

import (
	"github.com/Xebec19/simple-bank/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAccountRoutes(router *gin.RouterGroup, accountController *controllers.AccountController) {
	accounts := router.Group("/accounts")
	{
		accounts.POST("", accountController.CreateAccount)
		accounts.GET("", accountController.GetAccounts)
	}
}
