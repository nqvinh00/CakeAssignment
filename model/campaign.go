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
