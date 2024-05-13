package middlewares

import (
	authentication_models "Note-App/internal/models/authentication"
	"Note-App/internal/services/authentication"
	"Note-App/internal/services/logger"
	mongodbService "Note-App/internal/services/mongodb"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func TokenVerification(c *gin.Context) {
	token, err := c.Cookie("Authorization")
	if err != nil {
		logger.Warn("Unauthorized access to resource")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, err := authentication.ValidateToken(token)

	if err != nil {
		logger.Warn(err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		logger.Warn("Token expired")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userMongo, err := mongodbService.GetUserByEmail(claims["userId"].(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var user authentication_models.User
	if err := copier.Copy(&user, &userMongo); err != nil {
		logger.Error("Error copying user struct")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.Set("user", user)

	c.Next()
}