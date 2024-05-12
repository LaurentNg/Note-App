package routes

import (
	"Note-App/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(router *gin.RouterGroup) {
	router.POST("/signup", handlers.SignUp)
	router.POST("/signin", handlers.SignIn)
}