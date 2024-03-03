package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/nqvinh00/CakeAssignment/model"
	"github.com/pkg/errors"
)

var (
	sqlInsertVoucher = `
	INSERT INTO user_voucher (user_id, campaign_id, voucher, created_at, updated_at)
	VALUES (
		?,?,?,?,?
	)
	`

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

	// TODO: voucher expiration
)

type IUserVoucherDAO interface {
	Insert(ctx context.Context, voucher *model.UserVoucher) (err error)
	SelectByUserID(ctx context.Context, userID uint64) (vouchers []*model.UserVoucher, err error)
	SelectByCampaignID(ctx context.Context, campaignID uint64) (vouchers []*model.UserVoucher, err error)
}

type userVoucherDAO struct {
	db *sql.DB
}

func NewVoucherDAO(db *sql.DB) IUserVoucherDAO {
	return &userVoucherDAO{
		db: db,
	}
}

func (u *userVoucherDAO) Insert(ctx context.Context, voucher *model.UserVoucher) (err error) {
	now := time.Now()
	result, err := u.db.ExecContext(ctx, sqlInsertVoucher,
		voucher.UserID, voucher.CampaignID, voucher.Voucher, now, now,
	)
	if err != nil {
		return
	} else if result == nil {
		err = errors.New("invalid result from database server")
		return
	}

	return
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
