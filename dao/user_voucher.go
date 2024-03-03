package dao

import (
	"context"
	"database/sql"

	"github.com/nqvinh00/CakeAssignment/model"
)

var (
	sqlSelectVoucherByUserID = `
	SELECT id, user_id, campaign_id, voucher, created_at
		updated_at
	FROM user_voucher
	WHERE user_id = ?
	`

	sqlSelectVoucherByCampaignID = `
	SELECT id, user_id, campaign_id, voucher, created_at
		updated_at
	FROM user_voucher
	WHERE campaign_id = ?
	`
)

type IUserVoucherDAO interface {
}

type userVoucherDAO struct {
	db *sql.DB
}

func NewVoucherDAO(db *sql.DB) IUserVoucherDAO {
	return &userVoucherDAO{
		db: db,
	}
}

func (u *userVoucherDAO) SelectByUserID(ctx context.Context, userID uint64) (vouchers []*model.UserVoucher, err error) {
	vouchers = []*model.UserVoucher{}
	rows, err := u.db.QueryContext(ctx, sqlSelectVoucherByUserID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
	}
	defer rows.Close()

	for rows.Next() {
		v := &model.UserVoucher{}
		if err = rows.Scan(&v.ID, &v.UserID, &v.CampaignID, &v.Voucher, &v.CreatedAt, &v.UpdatedAt); err != nil {
			return
		}

		vouchers = append(vouchers, v)
	}

	return
}

func (u *userVoucherDAO) SelectByCampaignID(ctx context.Context, campaignID uint64) (vouchers []*model.UserVoucher, err error) {
	vouchers = []*model.UserVoucher{}
	vouchers = []*model.UserVoucher{}
	rows, err := u.db.QueryContext(ctx, sqlSelectVoucherByCampaignID, campaignID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
	}
	defer rows.Close()

	for rows.Next() {
		v := &model.UserVoucher{}
		if err = rows.Scan(&v.ID, &v.UserID, &v.CampaignID, &v.Voucher, &v.CreatedAt, &v.UpdatedAt); err != nil {
			return
		}

		vouchers = append(vouchers, v)
	}

	return
}
