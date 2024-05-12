package authentication

import (
	"Note-App/internal/services/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("secret-key")

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePasswords(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ 
		   "username": username, 
		   "exp": time.Now().Add(time.Hour * 24 * 30).Unix(), 
		})
   
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		logger.Error("Error generating token")
		return "", err
	}
   
	return tokenString, nil
}