package model

import "fmt"

// TEMP: For simple design
// TODO: Define error code and error message map
var (
	Success            = "Success"
	InvalidUsername    = "Invalid username"
	EmptyPassword      = "Password cannot be empty"
	AtLeastOne         = "Please submit at least username or phone number or email"
	InvalidPhoneNumber = "Invalid phone number"
	InvalidEmail       = "Invalid email"
	InvalidBirthday    = "Invalid birthday"
)

var (
	ErrUnknownError      = fmt.Errorf("unexpected error")
	ErrUserNotFound      = fmt.Errorf("user not found")
	ErrUserAlreadyExists = fmt.Errorf("user already exists")
)
