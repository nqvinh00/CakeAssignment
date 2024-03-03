package services

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/nqvinh00/CakeAssignment/dao"
	"github.com/nqvinh00/CakeAssignment/model"
	"github.com/rs/zerolog/log"
)

type VoucherDistributor interface {
	GetUserVouchers(ctx context.Context, userID uint64) (voucher []string, err error)
	CreateVoucher(ctx context.Context, campaignID, userID uint64) (voucher string, err error)
	AddVoucherForUser(ctx context.Context, user *model.User, voucher string) (err error)
}

type distributor struct {
	campaignDAO dao.ICampaignDAO
	voucherDAO  dao.IUserVoucherDAO
}

func NewVoucherDistributor(campaignDAO dao.ICampaignDAO, voucherDAO dao.IUserVoucherDAO) VoucherDistributor {
	return &distributor{
		campaignDAO: campaignDAO,
		voucherDAO:  voucherDAO,
	}
}

func (d *distributor) CreateVoucher(ctx context.Context, campaignID, userID uint64) (voucher string, err error) {
	campaign, err := d.campaignDAO.SelectByID(ctx, campaignID)
	if err != nil {
		return
	}

	isActive, err := d.checkCampaignValid(ctx, campaignID)
	if err != nil {
		return
	}

	if !isActive {
		return
	}

	// might use different way to do this
	hash := md5.Sum([]byte(fmt.Sprintf("%s-%d", campaign.Name, userID)))
	return hex.EncodeToString(hash[:]), nil
}

func (d *distributor) GetUserVouchers(ctx context.Context, userID uint64) (voucher []string, err error) {
	return
}

func (d *distributor) AddVoucherForUser(ctx context.Context, user *model.User, voucher string) (err error) {
	if err = d.voucherDAO.Insert(ctx, &model.UserVoucher{
		UserID:     user.ID,
		CampaignID: user.CampaignID,
		Voucher:    voucher,
	}); err != nil {
		log.Err(err).Msg("insert voucher failed")
	}

	return
}

func (d *distributor) checkCampaignValid(ctx context.Context, campaignID uint64) (ok bool, err error) {
	campaign, err := d.campaignDAO.SelectByID(ctx, campaignID)
	if err != nil {
		log.Err(err).Msg("select campaign failed")
		return false, err
	}

	// still process next transaction
	if campaign == nil {
		return false, nil
	}

	vouchers, err := d.voucherDAO.SelectByCampaignID(ctx, campaign.ID)
	if err != nil {
		log.Err(err).Msgf("select vouchers by campaign_id %d failed", campaign.ID)
		return false, err
	}

	if len(vouchers) >= campaign.VoucherCapacity {
		campaign.Status = 0 // deactive campaign
		if err = d.campaignDAO.Update(ctx, campaign); err != nil {
			log.Err(err).Msg("update campaign failed")
			// ignore error, keep processing
		}

		return false, nil
	}

	return true, nil
}
