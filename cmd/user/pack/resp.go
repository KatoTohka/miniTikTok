package pack

import (
	"errors"
	"miniTikTok/kitex_gen/user"
	"miniTikTok/pkg/errno"
)

func BuildRegisterResp(err error, userID int64, token string) *user.DouyinUserRegisterResponse {
	var resp user.DouyinUserRegisterResponse
	//两种种类的err
	if err == nil {
		resp.StatusCode = errno.Success.ErrCode
		resp.StatusMsg = &errno.Success.ErrMsg
		resp.UserId = userID
		resp.Token = token
		return &resp
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		resp.StatusCode = e.ErrCode
		resp.StatusMsg = &e.ErrMsg
		resp.UserId = userID
		resp.Token = token
		return &resp
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	resp.StatusCode = s.ErrCode
	resp.StatusMsg = &s.ErrMsg
	resp.UserId = userID
	resp.Token = token
	return &resp
}

func BuildLoginResp(err error, userID int64, token string) *user.DouyinUserLoginResponse {
	var resp user.DouyinUserLoginResponse
	if err == nil {
		resp.StatusCode = errno.Success.ErrCode
		resp.StatusMsg = &errno.Success.ErrMsg
		resp.UserId = userID
		resp.Token = token
		return &resp
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		resp.StatusCode = e.ErrCode
		resp.StatusMsg = &e.ErrMsg
		resp.UserId = userID
		resp.Token = token
		return &resp
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	resp.StatusCode = s.ErrCode
	resp.StatusMsg = &s.ErrMsg
	resp.UserId = userID
	resp.Token = token
	return &resp
}
func BuildInfoResp(err error, u *user.User) *user.DouyinUserResponse {
	var resp user.DouyinUserResponse
	if err == nil {
		resp.StatusCode = errno.Success.ErrCode
		resp.StatusMsg = &errno.Success.ErrMsg
		resp.User = u
		return &resp
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		resp.StatusCode = e.ErrCode
		resp.StatusMsg = &e.ErrMsg
		resp.User = u
		return &resp
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	resp.StatusCode = s.ErrCode
	resp.StatusMsg = &s.ErrMsg
	resp.User = u
	return &resp

}
