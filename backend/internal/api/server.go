package server

import (
	"Note-App/internal/api/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	routes.Encryptor_routes(&router.RouterGroup)
    router.Run("localhost:8080")
}