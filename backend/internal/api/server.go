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
	router.Use(cors.Default())
	api := router.Group("/api")
	{
		routes.EncryptorRoutes(api)
		routes.AuthenticationRoutes(api)
		routes.NotesRoutes(api)
	}
	mongodb.ConnectMongoDb()
	router.Run()
}