package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"miniTikTok/cmd/user/dal/db"
	"miniTikTok/kitex_gen/user"
	"miniTikTok/pkg/errno"
)

type UserLoginService struct {
	ctx context.Context
}

func NewUserLoginService(ctx context.Context) *UserLoginService {
	return &UserLoginService{
		ctx: ctx,
	}
}

func (s *UserLoginService) CheckUser(req *user.DouyinUserLoginRequest) (int64, error) {
	userName := req.Username
	res, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return -1, errno.AuthorizationFailedErr
	}
	if len(res) == 0 {
		return -1, errno.AuthorizationFailedErr
	}
	u := res[0]
	err = bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(req.Password))
	if err != nil {
		return -1, errno.AuthorizationFailedErr
	}
	return int64(u.ID), nil
}
