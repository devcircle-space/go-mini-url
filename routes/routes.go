package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(en *gin.Engine) {
	AuthRoutes(en.Group("/api/auth"))
	LinkRoutes(en.Group("/api/minify"))

	en.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
}
