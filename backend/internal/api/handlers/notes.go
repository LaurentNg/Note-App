package handlers

import (
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
	userId := c.Param("userId")

	notes, err := mongodb.GetNotesByUserId(userId)
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