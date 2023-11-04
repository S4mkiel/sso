package errors

import (
	"errors"
)

var ErrPasswordRecoveryTokenExpired = errors.New("password recovery token expired")
var ErrPasswordAndConfirmPasswordDontMatch = errors.New("password and confirm password mismatch")
var ErrDuplicatedEmail = errors.New("email already been taken")
var ErrDuplicatedUsername = errors.New("username already been taken")
var ErrDuplicatedExternalUID = errors.New("external uid already been taken")
var ErrInvalidToken = errors.New("unauthorized Token")
var ErrInvalidTokenFormat = errors.New("invalid Token format")
