package mongodb_errors

import (
	"errors"
)

func ErrCreateNote() error {
	mes := "Error while saving note in database"
	return errors.New(mes)
}