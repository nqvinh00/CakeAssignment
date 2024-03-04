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

	if req.Password == "" {
		return InvalidPassword
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

	r, _ := regexp.Compile(`^(0|84)(56|58|59|7[67890]|3[2-9]|8[1-5]|9\d|16[2-9]|12\d|86|88|89|186|188|199)(\d{7})$`)
	if req.PhoneNumber != "" && !r.Match([]byte(req.PhoneNumber)) {
		return InvalidPhoneNumber
	}

	_, err := mail.ParseAddress(req.Email)
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
