package authentication_errors

import "errors"

var (
	ErrUnexpectedSigningMethod = errors.New("Unexpected signing method")
)