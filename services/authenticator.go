package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nqvinh00/CakeAssignment/dao"
	"github.com/nqvinh00/CakeAssignment/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type Authenticator interface {
	CreateUser(ctx context.Context, user *model.User, password string) (err error)
	Login(ctx context.Context, username, password string) (token string, err error)
}

type authenticator struct {
	userDAO            dao.IUserDAO
	userSecDAO         dao.IUserSecDAO
	voucherDistributor VoucherDistributor
	jwtSecretKey       string
}

func NewAuthenticator(userDAO dao.IUserDAO, userSecDAO dao.IUserSecDAO, jwtSecretKey string) Authenticator {
	return &authenticator{
		userDAO:      userDAO,
		userSecDAO:   userSecDAO,
		jwtSecretKey: jwtSecretKey,
	}
}

func (a *authenticator) CreateUser(ctx context.Context, user *model.User, password string) (err error) {
	u, err := a.userDAO.SelectToSignup(ctx, user.Username, user.Email, user.PhoneNumber)
	if err != nil {
		log.Err(err).Msg("failed to select user")
		return
	}

	if u != nil {
		return model.ErrUserAlreadyExists
	}

	id, err := a.userDAO.Insert(ctx, user)
	if err != nil {
		log.Err(err).Msg("failed to insert user")
		return
	}

	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Err(err).Msg("failed to generate password")
		return
	}

	if err = a.userSecDAO.Insert(ctx, &model.UserSec{
		UserID:   id,
		Password: p,
	}); err != nil {
		log.Err(err).Msg("failed to insert user sec")
		return
	}

	return
}

func (a *authenticator) Login(ctx context.Context, username, password string) (token string, err error) {
	user, err := a.userDAO.SelectToLogin(ctx, username)
	if err != nil {
		log.Err(err).Msg("failed to select user")
		return
	}

	if user == nil {
		err = model.ErrUserNotFound
		return
	}

	if user.Status != model.UserActivatedStatus {
		err = errors.New("account deactived")
		return
	}

	userSec, err := a.userSecDAO.Select(ctx, user.ID)
	if err != nil {
		log.Err(err).Msg("failed to select user sec")
		return
	}

	if userSec == nil {
		err = model.ErrUserNotFound
		return
	}

	if err = bcrypt.CompareHashAndPassword(userSec.Password, []byte(password)); err != nil {
		user.LoginAttempt--
		if user.LoginAttempt <= 0 {
			user.Status = model.UserDeactivatedStatus
		}

		if er := a.userDAO.Update(ctx, user); er != nil {
			log.Err(er).Msg("failed to update user")
			return
		}

		err = fmt.Errorf("assword is incorrect, you only have %d times left to retry", user.LoginAttempt)
		return
	}

	// Null last login = first time login
	if !user.LastLogin.Valid {
		go func(user *model.User) {
			if user.CampaignID == 0 {
				return
			}

			voucher, err := a.voucherDistributor.CreateVoucher(ctx, user.CampaignID, user.ID)
			if err != nil {
				log.Err(err).Msgf("create voucher for userId %d failed", user.ID)
				return
			}

			if voucher != "" {
				if err := a.voucherDistributor.AddVoucherForUser(ctx, user, voucher); err != nil {
					log.Err(err).Msgf("create voucher for userId %d failed", user.ID)
				}
			}

			// TODO: notify
		}(user)
	}

	user.LoginAttempt = 3
	user.LastLogin = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	if err = a.userDAO.Update(ctx, user); err != nil {
		log.Err(err).Msg("failed to update user")
		return
	}

	expire := time.Now().Add(time.Hour * 24)
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.Claim{
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
		},
	})

	token, err = claim.SignedString([]byte(a.jwtSecretKey))
	if err != nil {
		log.Err(err).Msg("failed to sign string")
		return
	}
	return
}
