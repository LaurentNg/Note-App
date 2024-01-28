package server

import (
	"Note-App/internal/api/handlers"

	"github.com/gin-gonic/gin"
)
func encryptor_routes(router *gin.Engine) {
	router.POST("/encrypt", handlers.Encrypt)
	router.POST("/decrypt", handlers.Decrypt)
}

func Run() {
	router := gin.Default()
	encryptor_routes(router)
    router.Run("localhost:8080")
}