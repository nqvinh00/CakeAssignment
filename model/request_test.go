package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateLoginReq(t *testing.T) {
	req := &LoginReq{
		Username: "test",
		Password: "test",
	}

	t.Run("happy case", func(t *testing.T) {
		errMsg := req.Valid()
		assert.Equal(t, Success, errMsg)
	})

	t.Run("invalid username", func(t *testing.T) {
		req.Username = ""
		errMsg := req.Valid()
		assert.Equal(t, InvalidUsername, errMsg)
	})

	t.Run("invalid password", func(t *testing.T) {
		req.Username = "test"
		req.Password = ""
		errMsg := req.Valid()
		assert.Equal(t, InvalidPassword, errMsg)
	})
}

func TestValidateNewUserReq(t *testing.T) {
	req := &NewUserReq{
		Username:    "test",
		Fullname:    "test",
		PhoneNumber: "0966762092",
		Email:       "test@gmail.com",
		Password:    "test",
		StrBirthday: "14/12/2000",
		CampaignID:  1,
	}

	t.Run("happy case", func(t *testing.T) {
		errMsg := req.Valid()
		assert.Equal(t, Success, errMsg)
	})

	t.Run("at least one", func(t *testing.T) {
		testcase := &NewUserReq{
			Fullname:    req.Fullname,
			Password:    req.Password,
			StrBirthday: req.StrBirthday,
			CampaignID:  req.CampaignID,
		}

		errMsg := testcase.Valid()
		assert.Equal(t, AtLeastOne, errMsg)
	})

	t.Run("invalid phone number", func(t *testing.T) {
		testcase := &NewUserReq{
			Username:    req.Username,
			Fullname:    req.Fullname,
			PhoneNumber: "123",
			Email:       req.Email,
			Password:    req.Password,
			StrBirthday: req.StrBirthday,
			CampaignID:  req.CampaignID,
		}

		errMsg := testcase.Valid()
		assert.Equal(t, InvalidPhoneNumber, errMsg)
	})

	t.Run("invalid email", func(t *testing.T) {
		testcase := &NewUserReq{
			Username:    req.Username,
			Fullname:    req.Fullname,
			PhoneNumber: req.PhoneNumber,
			Email:       "test@gmail.",
			Password:    req.Password,
			StrBirthday: req.StrBirthday,
			CampaignID:  req.CampaignID,
		}

		errMsg := testcase.Valid()
		assert.Equal(t, InvalidEmail, errMsg)
	})

	t.Run("invalid birthday", func(t *testing.T) {
		testcase := &NewUserReq{
			Username:    req.Username,
			Fullname:    req.Fullname,
			PhoneNumber: req.PhoneNumber,
			Email:       req.Email,
			Password:    req.Password,
			CampaignID:  req.CampaignID,
		}

		errMsg := testcase.Valid()
		assert.Equal(t, InvalidBirthday, errMsg)
	})

	t.Run("wrong birthday format", func(t *testing.T) {
		testcase := &NewUserReq{
			Username:    req.Username,
			Fullname:    req.Fullname,
			PhoneNumber: req.PhoneNumber,
			Email:       req.Email,
			Password:    req.Password,
			StrBirthday: "14-12-2000",
			CampaignID:  req.CampaignID,
		}

		errMsg := testcase.Valid()
		assert.Equal(t, InvalidBirthday, errMsg)
	})
}
