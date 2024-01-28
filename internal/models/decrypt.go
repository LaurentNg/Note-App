package models

type Decrypt struct {
    EncryptedText string  `json:"text"`
    Passphrase    string  `json:"passphrase"`
}
