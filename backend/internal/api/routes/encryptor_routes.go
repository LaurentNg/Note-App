package routes

import (
	"Note-App/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func Encryptor_routes(router *gin.RouterGroup) {
	router.POST("/encrypt", handlers.Encrypt)
	router.POST("/decrypt", handlers.Decrypt)
}