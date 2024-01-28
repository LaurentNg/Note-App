package models

type Decrypt struct {
    EncryptedText string  `json:"text" binding:"required"`
    Passphrase    string  `json:"passphrase" binding:"required"`
}
