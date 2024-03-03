package model

import (
	"net/mail"
	"regexp"
	"time"
)

type Validator interface {
	Valid() string
}

type LoginReq struct {
	// Username could be username/phone_number/email address
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *LoginReq) Valid() string {
	if req.Username == "" {
		return InvalidUsername
	}

	_, err := mail.ParseAddress(req.Username)
	if regexp.MustCompile(`\d`).MatchString(req.Username) || err != nil {
		return InvalidUsername
	}

	if req.Password == "" {
		return EmptyPassword
	}

	return Success
}

type NewUserReq struct {
	Username    string `json:"username"`
	Fullname    string `json:"fullname"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	StrBirthday string `json:"birthday"`
	CampaignID  uint64 `json:"campaign_id"` // User signup for
	Birthday    time.Time
}

func (req *NewUserReq) Valid() string {
	if req.Username == "" && req.Email == "" && req.PhoneNumber == "" {
		return AtLeastOne
	}

	if req.PhoneNumber != "" && regexp.MustCompile(`\d`).MatchString(req.PhoneNumber) {
		return InvalidPhoneNumber
	}

	_, err := mail.ParseAddress(req.Username)
	if req.Email != "" && err != nil {
		return InvalidEmail
	}

	if req.StrBirthday == "" {
		return InvalidBirthday
	}

	req.Birthday, err = time.Parse("02/01/2006", req.StrBirthday)
	if err != nil {
		return InvalidBirthday
	}

	return Success
}
