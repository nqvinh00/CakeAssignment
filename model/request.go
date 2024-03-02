package model

import (
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
	req.Birthday, _ = time.Parse("02/01/2006", req.StrBirthday)
	return Success
}
