package domain

import "errors"

var (
	ErrNotFound               = errors.New("not found")
	ErrForbidden              = errors.New("forbidden")
	ErrInvalidToken           = errors.New("invalid token")
	ErrTokenNotFound          = errors.New("token not found")
	ErrInvalidLoginOrPassword = errors.New("invalid logon or password")
	ErrEmailAlreadyExists     = errors.New("email already exists")
)
