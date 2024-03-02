package model

import "fmt"

const Success = "Success"

// TEMP: For simple design
// TODO: Define error code and error message map
var (
	ErrUnknownError      = fmt.Errorf("unexpected error")
	ErrUserNotFound      = fmt.Errorf("user not found")
	ErrUserAlreadyExists = fmt.Errorf("user already exists")
)
