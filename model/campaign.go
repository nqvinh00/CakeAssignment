package model

import "time"

type Campaign struct {
	ID              uint64    `db:"id"`
	Name            string    `db:"name"`
	Status          int8      `db:"status"`
	VoucherCapacity int       `db:"voucher_capacity"`
	StartDate       time.Time `db:"start_date"`
	EndDate         time.Time `db:"end_date"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type UserVoucher struct {
	ID         uint64    `db:"id"`
	UserID     uint64    `db:"user_id"`
	CampaignID uint64    `db:"campaign_id"`
	Voucher    string    `db:"voucher"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `dB:"updated_at"`
}
