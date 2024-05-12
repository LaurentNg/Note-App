package models

type User struct {
    Email     string  `json:"email" binding:"required"`
    Password  string  `json:"password" binding:"required"`
    Username  string  `json:"username" binding:"required"`
}