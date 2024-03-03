package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/nqvinh00/CakeAssignment/model"
	"github.com/pkg/errors"
)

var (
	sqlSelectCampaignByID = `
	SELECT id, name, status, voucher_capacity, start_date,
		end_date, created_at, updated_at
	FROM campaign
	WHERE id = ?
	AND status = 1
	AND end_date >= ?
	`

	sqlUpdateCampaign = `
	UPDATE campaign
	SET
		status = ?, voucher_capacity = ?, updated_at = ?
	WHERE id = ?
	`
)

type ICampaignDAO interface {
	SelectByID(ctx context.Context, id uint64) (campaign *model.Campaign, err error)
	Update(ctx context.Context, campaign *model.Campaign) (err error)
}

type campaignDAO struct {
	db *sql.DB
}

func NewCampaignDAO(db *sql.DB) ICampaignDAO {
	return &campaignDAO{
		db: db,
	}
}

func (c *campaignDAO) SelectByID(ctx context.Context, id uint64) (campaign *model.Campaign, err error) {
	campaign = &model.Campaign{}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	row := c.db.QueryRowContext(ctx, sqlSelectCampaignByID, id, today)
	err = row.Scan(&campaign.ID, &campaign.Name, &campaign.Status, &campaign.VoucherCapacity, &campaign.StartDate,
		&campaign.EndDate, &campaign.CreatedAt, &campaign.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return
	}

	return campaign, nil
}

func (c *campaignDAO) Update(ctx context.Context, campaign *model.Campaign) (err error) {
	result, err := c.db.ExecContext(ctx, sqlUpdateCampaign,
		campaign.Status, campaign.VoucherCapacity, time.Now(),
		campaign.ID,
	)
	if err != nil {
		return
	} else if result == nil {
		err = errors.New("invalid result from database server")
		return
	}

	row, err := result.RowsAffected()
	if err != nil {
		return
	}

	if row == 0 {
		err = sql.ErrNoRows
	}

	return
}
