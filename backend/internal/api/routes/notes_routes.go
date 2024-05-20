package routes

import (
	"Note-App/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func NotesRoutes(router *gin.RouterGroup) {
	router.POST("/note", handlers.CreateNote)
	router.GET("/notes/:userId", handlers.GetNotesByUserId)
}