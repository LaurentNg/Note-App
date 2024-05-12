package mongodb_errors

import (
	"errors"
	"fmt"
)

func ErrExistingUser(email string, username string) error {
	mes := fmt.Sprintf("Duplicate : user with email: %s and username: %s already exist", email, username)
	return errors.New(mes)
}

func ErrUserNotFound(email string) error {
	mes := fmt.Sprintf("User with email: %s not found", email)
	return errors.New(mes)
}

func ErrHashPassword(email string, username string) error {
	mes := fmt.Sprintf("Error hashing password for user with email: %s and username: %s", email, username)
	return errors.New(mes)
}

func ErrCreateUser(email string, username string) error {
	mes := fmt.Sprintf("Error creating user with email: %s and username: %s", email, username)
	return errors.New(mes)
}