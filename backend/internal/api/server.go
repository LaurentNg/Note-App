package server

import (
	"Note-App/internal/api/routes"
	"Note-App/internal/services/mongodb"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.Use(cors.Default())
	routes.EncryptorRoutes(&router.RouterGroup)
	routes.AuthenticationRoutes(&router.RouterGroup)
	mongodb.ConnectMongoDb()
    router.Run()
}