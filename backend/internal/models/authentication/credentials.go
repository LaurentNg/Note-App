package authentication_models

type Credentials struct {
    Email     string  `json:"email" binding:"required"`
    Password  string  `json:"password" binding:"required"`
}