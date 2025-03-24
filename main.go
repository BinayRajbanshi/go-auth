package main

import (
	"os"

	"github.com/BinayRajbanshi/go-auth/controllers"
	"github.com/BinayRajbanshi/go-auth/database"
	"github.com/BinayRajbanshi/go-auth/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadEnvVariables()
	database.ConnectToDb()
	utils.Migrate()
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// routes.AuthRoutes(router)
	// routes.UserRoutes(router)

	router.POST("/api/v1/signup", controllers.Signup)
	router.POST("/api/v1/login", controllers.Login)

	router.Run(":" + port)
}
