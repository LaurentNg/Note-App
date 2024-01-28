package handlers

import (
	"Note-App/internal/encryptor"
	"Note-App/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// https://go.dev/doc/tutorial/web-service-gin

func Encrypt(c *gin.Context) {
    var encrypt models.Encrypt

    if err := c.BindJSON(&encrypt); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "The body does not respect format"})
        return
    }

    encryptedText, err := encryptor.Encrypt(encrypt.Passphrase, encrypt.Text)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

    c.IndentedJSON(http.StatusCreated, gin.H{"text": encryptedText})
}

func Decrypt(c *gin.Context) {
    var encrypt models.Decrypt

    if err := c.BindJSON(&encrypt); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "The body does not respect format"})
        return
    }

    decryptedText, err := encryptor.Decrypt(encrypt.Passphrase, encrypt.EncryptedText)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

    c.JSON(http.StatusCreated, gin.H{"text": decryptedText})
}