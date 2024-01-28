package encryptor

import (
	"crypto/aes"
	"encoding/hex"
	"errors"
)

// https://www.developer.com/languages/cryptography-in-go/

var (
	ErrWrongPassphrase = errors.New("The provided passphrase is incorrect")
)

func Encrypt(key string, text string) (string, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	textByte := make([]byte, len(text))
	c.Encrypt(textByte, []byte(text))
	return hex.EncodeToString(textByte), nil
}

func Decrypt(key string, text string) (string, error) {
	txt, _ := hex.DecodeString(text)
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", ErrWrongPassphrase
	}
	textByte := make([]byte, len(txt))
	c.Decrypt(textByte, []byte(txt))

	return string(textByte), nil
}