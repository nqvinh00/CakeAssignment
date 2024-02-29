package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/nqvinh00/CakeAssignment/model"
	"github.com/pkg/errors"
)

var (
	sqlInsertUser = `
	INSERT INTO user (fullname, phone_number, email, username, campaign_id,
		checksum, created_at, updated_at)
	VALUES (
		?,?,?,?,?,
		?,?,?
	)
	`

	sqlUpdateUser = `
	UPDATE user
	SET
		status = ?, login_attempt = ?, updated_at = ?
	WHERE id = ?
	`

	sqlSelectUser = `
	SELECT id, fullname, phone_number, email, username,
		campaign_id, status, login_attempt, checksum, created_at,
		updated_at
	WHERE phone_number = ?
	OR email = ?
	OR username = ?
	`
)

type IUserDAO interface {
	Insert(ctx context.Context, user *model.User) (id uint64, err error)
	Update(ctx context.Context, user *model.User) (err error)
	Select(ctx context.Context, loginValue string) (user *model.User, err error)
}

type userDAO struct {
	db *sql.DB
}

func NewUserDAO(db *sql.DB) IUserDAO {
	return &userDAO{
		db: db,
	}
}

func (u *userDAO) Insert(ctx context.Context, user *model.User) (id uint64, err error) {
	now := time.Now()
	result, err := u.db.ExecContext(ctx, sqlInsertUser,
		user.Fullname, user.PhoneNumber, user.Email, user.Username, user.CampaignID,
		user.Checksum, now, now,
	)
	if err != nil {
		return
	} else if result == nil {
		err = errors.New("invalid result from database server")
		return
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return
	}

	id = uint64(userID)
	return
}

func (u *userDAO) Update(ctx context.Context, user *model.User) (err error) {
	result, err := u.db.ExecContext(ctx, sqlUpdateUser,
		user.Status, user.LoginAttempt, time.Now(),
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

func (u *userDAO) Select(ctx context.Context, loginValue string) (user *model.User, err error) {
	user = &model.User{}
	if err = u.db.QueryRowContext(ctx, sqlSelectUser, loginValue, loginValue, loginValue).Scan(user); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return
	}

	return user, nil
}
