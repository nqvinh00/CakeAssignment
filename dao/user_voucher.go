package dao

import "database/sql"

type IVoucherDAO interface {
}

type voucherDAO struct {
	db *sql.DB
}

func NewVoucherDAO(db *sql.DB) IVoucherDAO {
	return &voucherDAO{
		db: db,
	}
}
