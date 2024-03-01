package model

import "time"

type Validator interface {
	Valid() string
}

type LoginReq struct {
	// Username could be username/phone_number/email address
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *LoginReq) Valid() string {
	return Success
}

type NewUserReq struct {
	Username    string `json:"username"`
	Fullname    string `json:"fullname"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	StrBirthday string `json:"birthday"`
	Birthday    time.Time
}

func (req *NewUserReq) Valid() string {
	return Success
}
