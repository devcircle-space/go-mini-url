package routes

import (
	"devcircle.space/mini-url/controllers"
	"github.com/gin-gonic/gin"
)

func LinkRoutes(rg *gin.RouterGroup) {
	// public routes
	rg.GET("/:id", controllers.GetFromMinifiedUrl)
	// protected routes
	rg.POST("", controllers.CreateMinifiedUrl)
	rg.GET("", controllers.GetAllMinifiedUrls)
	rg.PUT("/:id", controllers.UpdateMinifiedUrl)
	rg.DELETE("/:id", controllers.DeleteMinifiedUrl)
}
