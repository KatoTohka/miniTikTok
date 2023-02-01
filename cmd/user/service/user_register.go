package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"miniTikTok/cmd/user/dal/db"
	"miniTikTok/kitex_gen/user"
	"miniTikTok/pkg/errno"
)

type UserRegisterService struct {
	ctx context.Context
}

// NewUserRegisterService new UserRegisterService
func NewUserRegisterService(ctx context.Context) *UserRegisterService {
	return &UserRegisterService{
		ctx: ctx,
	}
}

// RegisterUser call create User in db
func (s *UserRegisterService) RegisterUser(req *user.DouyinUserRegisterRequest) (userID int64, err error) {
	// query user exist or not
	users, err := db.QueryUser(s.ctx, req.Username)

	if err != nil {
		return -1, err
	}
	if len(users) != 0 {
		return -1, errno.UserAlreadyExistErr
	}

	// password use md5 encrypt
	h := md5.New()

	if _, err = io.WriteString(h, req.Password); err != nil {
		return -1, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	// register
	err = db.RegisterUser(s.ctx, []*db.User{
		{
			UserName: req.Username,
			PassWord: passWord,
		},
	})
	if err != nil {
		return -1, err
	}
	// query id
	queryUser, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return -1, err
	}
	return int64(queryUser[0].ID), nil
}
