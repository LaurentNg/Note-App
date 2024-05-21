package routes

import (
	"Note-App/internal/api/handlers"
	"Note-App/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func NotesRoutes(router *gin.RouterGroup) {
	router.POST("/note", middlewares.TokenVerification, handlers.CreateNote)
	router.GET("/notes", middlewares.TokenVerification, handlers.GetNotesByUserId)
	router.DELETE("/notes/:id", middlewares.TokenVerification, handlers.DeleteNoteById)
}