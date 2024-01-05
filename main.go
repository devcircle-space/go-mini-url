package main

import (
	"fmt"
	"os"
	"time"

	"devcircle.space/mini-url/routes"
	"devcircle.space/mini-url/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// set app mode
	APP_MODE := os.Getenv("APP_MODE")
	if APP_MODE == "production" {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("Running in production mode")
	} else {
		utils.LoadEnv()
		fmt.Println("Running in debug mode")
	}

	// initialize gin
	e := gin.Default()

	// using cors middlware
	e.Use((cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	})))

	// register routes
	routes.RegisterRoutes(e)

	// run server
	PORT := os.Getenv("PORT")
	fmt.Println("🔌 Server running on port: " + PORT)

	e.Run(":" + PORT)
}
