package model

type Validator interface {
	Valid() ErrorCode
}

type LoginReq struct {
	// Username could be username/phone_number/email address
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *LoginReq) Valid() ErrorCode {
	return Success
}
