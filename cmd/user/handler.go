package main

import (
	"context"
	"miniTikTok/cmd/user/pack"
	"miniTikTok/cmd/user/service"
	"miniTikTok/kitex_gen/user"
	"miniTikTok/middleware"
	"miniTikTok/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildRegisterResp(errno.ParamErr, -1, "")
		return resp, nil
	}
	// true call register
	var userId int64
	userId, err = service.NewUserRegisterService(ctx).RegisterUser(req)
	if err != nil {
		resp = pack.BuildRegisterResp(err, -1, "")
		return resp, err
	}
	// generate token
	token, err := middleware.CreateToken(userId)
	if err != nil {
		resp = pack.BuildRegisterResp(err, -1, "")
	}
	resp = pack.BuildRegisterResp(nil, userId, token)
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildLoginResp(errno.ParamErr, -1, "")
		return resp, nil
	}
	// true call register
	var userId int64
	userId, err = service.NewUserLoginService(ctx).CheckUser(req)
	if err != nil {
		resp = pack.BuildLoginResp(err, -1, "")
		return resp, err
	}
	// generate token
	token, err := middleware.CreateToken(userId)
	if err != nil {
		resp = pack.BuildLoginResp(err, -1, "")
	}
	resp = pack.BuildLoginResp(nil, userId, token)
	return resp, nil
}

// QueryUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUser(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	if len(req.Token) == 0 {
		resp = pack.BuildInfoResp(errno.ParamErr, nil)
		return resp, nil
	}
	// true call register
	var u *user.User
	u, err = service.NewUserInfoService(ctx).InfoUser(req)
	if err != nil {
		resp = pack.BuildInfoResp(err, nil)
		return resp, err
	}
	if err != nil {
		resp = pack.BuildInfoResp(err, nil)
	}
	resp = pack.BuildInfoResp(nil, u)
	return resp, nil
}
