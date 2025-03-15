package routes

import (
	"github.com/BinayRajbanshi/go-auth/controllers"
	"github.com/BinayRajbanshi/go-auth/middleware"
	"github.com/gin-gonic/gin"
)

// user routes are protected so use middleware to protect them
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:id", controllers.GetUser())
}
