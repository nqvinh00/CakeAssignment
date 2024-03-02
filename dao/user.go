package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/nqvinh00/CakeAssignment/model"
	"github.com/pkg/errors"
)

var (
	sqlInsertUser = `
	INSERT INTO user (fullname, phone_number, email, username, campaign_id,
		birthday, created_at, updated_at)
	VALUES (
		?,?,?,?,?,
		?,?,?
	)
	`

	sqlUpdateUser = `
	UPDATE user
	SET
		status = ?, login_attempt = ?, last_login = ?, updated_at = ?
	WHERE id = ?
	`

	sqlSelectUserTemplate = `
	SELECT id, fullname, phone_number, email, username,
	campaign_id, status, login_attempt, birthday,
	last_login, created_at, updated_at
	FROM user
	WHERE %s
	`
)

type IUserDAO interface {
	Insert(ctx context.Context, user *model.User) (id uint64, err error)
	Update(ctx context.Context, user *model.User) (err error)
	SelectToLogin(ctx context.Context, loginValue string) (user *model.User, err error)
	SelectToSignup(ctx context.Context, username, email, phoneNumber string) (user *model.User, err error)
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
		user.Birthday, now, now,
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
		user.Status, user.LoginAttempt, user.LastLogin, time.Now(),
		user.ID,
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

func (u *userDAO) SelectToLogin(ctx context.Context, loginValue string) (user *model.User, err error) {
	user = &model.User{}
	query := fmt.Sprintf(sqlSelectUserTemplate, "(username = ? OR email = ? OR phone_number = ?) AND status = ?")

	row := u.db.QueryRowContext(ctx, query, loginValue, loginValue, loginValue, model.UserActivatedStatus)
	err = row.Scan(&user.ID, &user.Fullname, &user.PhoneNumber, &user.Email, &user.Username,
		&user.CampaignID, &user.Status, &user.LoginAttempt, &user.Birthday,
		&user.LastLogin, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return
	}

	return user, nil
}

func (u *userDAO) SelectToSignup(ctx context.Context, username, email, phoneNumber string) (user *model.User, err error) {
	user = &model.User{}

	where := ""
	if username != "" {
		where += fmt.Sprintf(`username = "%s"`, username)
	}

	if email != "" {
		where += fmt.Sprintf(`OR email = "%s"`, email)
	}

	if phoneNumber != "" {
		where += fmt.Sprintf(`OR phone_number = "%s"`, phoneNumber)
	}

	row := u.db.QueryRowContext(ctx, fmt.Sprintf(sqlSelectUserTemplate, where))
	err = row.Scan(&user.ID, &user.Fullname, &user.PhoneNumber, &user.Email, &user.Username,
		&user.CampaignID, &user.Status, &user.LoginAttempt, &user.Birthday,
		&user.LastLogin, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return
	}

	return user, nil
}
