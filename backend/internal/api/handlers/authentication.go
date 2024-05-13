package handlers

import (
	authentication_models "Note-App/internal/models/authentication"
	mongodb_models "Note-App/internal/models/mongodb"
	authService "Note-App/internal/services/authentication"
	"Note-App/internal/services/logger"
	mongodbService "Note-App/internal/services/mongodb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func SignUp(c *gin.Context) {
	var user authentication_models.User
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

	if err := mongodbService.CreateUser(&userMongo); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func SignIn(c *gin.Context) {
	var credentials authentication_models.Credentials
	if err := c.BindJSON(&credentials); err != nil {
		logger.Error("Error binding user struct")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userMongo, err := mongodbService.GetUserByEmail(credentials.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = authService.ComparePasswords(userMongo.Password, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := authService.GenerateToken(userMongo.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600 * 24 * 30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func SignOut(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	
	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	logger.Info("Validate middleware")
	user, _ := c.Get("user")
	
	c.JSON(http.StatusOK, gin.H{"user": user.(authentication_models.User)})
}