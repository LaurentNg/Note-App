package encryptor

import (
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"

	"golang.org/x/crypto/hkdf"
)

// https://www.developer.com/languages/cryptography-in-go/

var (
	ErrWrongPassphrase = errors.New("The provided passphrase is incorrect")
)

func Encrypt(passphrase string, text string) (string, error) {
	key, err := generate_key(passphrase)
	if err != nil {
		return "", err
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// TODO: Check to enable any length text
	textByte := make([]byte, len(text))
	c.Encrypt(textByte, []byte(text))
	return hex.EncodeToString(textByte), nil
}

func Decrypt(passphrase string, text string) (string, error) {
	txt, _ := hex.DecodeString(text)

	key, err := generate_key(passphrase)
	if err != nil {
		return "", err
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", ErrWrongPassphrase
	}
	textByte := make([]byte, len(txt))
	c.Decrypt(textByte, []byte(txt))

	return string(textByte), nil
}

func generate_key(passphrase string) ([]byte, error) {
	hkdfSalt := make([]byte, 32)
	info := []byte("AES Key Derivation")
	key := make([]byte, 32)

	r := hkdf.New(sha256.New, []byte(passphrase), hkdfSalt, info)
	if _, err := io.ReadFull(r, key); err != nil {
		return nil, err
	}

	return key, nil
}
