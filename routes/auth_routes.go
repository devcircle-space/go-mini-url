package routes

import (
	"devcircle.space/mini-url/controllers"
	"devcircle.space/mini-url/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup) {
	// public routes
	rg.POST("/login", controllers.Login)
	rg.POST("/register", controllers.Register)
	rg.POST("/reset-password", controllers.GetResetPasswordLink)

	// protected routes
	rg.Use(middlewares.ValidateAuthToken)
	rg.POST("/logout", controllers.Logout)
	rg.GET("/me", controllers.VerifyUser)
	rg.PUT("/reset-password", controllers.ResetPassword)
	rg.DELETE("/me", controllers.DeleteAccount)
}
