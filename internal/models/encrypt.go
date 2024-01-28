package models

type Encrypt struct {
    Text     	string  `json:"text" binding:"required"`
    Passphrase  string  `json:"passphrase" binding:"required"`
}
