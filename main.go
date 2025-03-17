package main

import (
	"os"

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

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
	router.Run(":" + port)
}
