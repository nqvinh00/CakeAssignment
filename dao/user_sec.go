package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/nqvinh00/CakeAssignment/model"
	"github.com/pkg/errors"
)

var (
	sqlInserUserSec = `
	INSERT INTO user_security (user_id, password, created_at, updated_at)
	VALUES (
		?,?,?,?
	)
	`

	// Might use for changing password, not use now
	sqlUpdateUserSec = `
	UPDATE user_security
	SET
		password = ?, updated_at ?
	WHERE user_id = ?
	`

	sqlSelectUserSec = `
	SELECT user_id, password, created_at, updated_at
	FROM user_security
	WHERE user_id = ?
	`
)

type IUserSecDAO interface {
	Insert(ctx context.Context, us *model.UserSec) (err error)
	Update(ctx context.Context, us *model.UserSec) (err error)
	Select(ctx context.Context, userID uint64) (us *model.UserSec, err error)
}

type userSecDAO struct {
	db *sql.DB
}

func NewUserSecDAO(db *sql.DB) IUserSecDAO {
	return &userSecDAO{
		db: db,
	}
}

func (u *userSecDAO) Insert(ctx context.Context, us *model.UserSec) (err error) {
	now := time.Now()
	result, err := u.db.ExecContext(ctx, sqlInserUserSec,
		us.UserID, us.Password, now, now,
	)
	if err != nil {
		return
	} else if result == nil {
		err = errors.New("invalid result from database server")
		return
	}

	return
}

func (u *userSecDAO) Update(ctx context.Context, us *model.UserSec) (err error) {
	result, err := u.db.ExecContext(ctx, sqlUpdateUserSec,
		us.Password, time.Now(),
		us.UserID,
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

func (u *userSecDAO) Select(ctx context.Context, userID uint64) (us *model.UserSec, err error) {
	us = &model.UserSec{}
	if err = u.db.QueryRowContext(ctx, sqlSelectUserSec, userID).Scan(us); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return
	}

	return us, nil
}
