package handlers

import (
	authentication_models "Note-App/internal/models/authentication"
	noteMongodb_models "Note-App/internal/models/mongodb"
	note_models "Note-App/internal/models/note"
	"Note-App/internal/services/logger"
	"Note-App/internal/services/mongodb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func CreateNote(c *gin.Context) {
	var note note_models.Note

	if err := c.BindJSON(&note); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note.UserId = c.MustGet("user").(authentication_models.User).ID
	var noteMongodb noteMongodb_models.Note
	if err := copier.Copy(&noteMongodb, &note); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	err := mongodb.CreateNote(&noteMongodb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Note created successfully"})
}

func GetNotesByUserId(c *gin.Context) {
	user, _ := c.Get("user")

	notes, err := mongodb.GetNotesByUserId(user.(authentication_models.User).ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var notesResponse []note_models.Note
	if err := copier.Copy(&notesResponse, &notes); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notesResponse)
}