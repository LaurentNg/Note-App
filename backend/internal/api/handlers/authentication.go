package handlers

import (
	"Note-App/internal/models"
	mongodb_models "Note-App/internal/models/mongodb"
	"Note-App/internal/services/logger"
	"Note-App/internal/services/mongodb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)


func SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logger.Error("Error binding user struct")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userMongo mongodb_models.User
	if err := copier.Copy(&userMongo, &user); err != nil {
		logger.Error("Error copying user struct")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := mongodb.CreateUser(&userMongo); err != nil {
		logger.Error("Error creating user")
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}