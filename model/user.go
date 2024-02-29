package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID           uint64    `db:"id"`
	Fullname     string    `db:"fullname"`
	PhoneNumber  string    `db:"phone_number"`
	Email        string    `db:"email"`
	Username     string    `db:"username"`
	CampaignID   uint64    `db:"campaign_id"`
	Status       int8      `db:"status"`
	LoginAttempt int8      `db:"login_attempt"`
	Checksum     uint64    `db:"checksum"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type UserSec struct {
	UserID    uint64    `db:"user_id"`
	Password  []byte    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Claim struct {
	Username string
	Email    string
	jwt.StandardClaims
}
