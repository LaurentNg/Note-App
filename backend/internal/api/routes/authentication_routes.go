package routes

import (
	"Note-App/internal/api/handlers"
	"Note-App/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(router *gin.RouterGroup) {
	router.POST("/signup", handlers.SignUp)
	router.POST("/signin", handlers.SignIn)
	router.GET("/validate", middlewares.TokenVerification, handlers.Validate)
}