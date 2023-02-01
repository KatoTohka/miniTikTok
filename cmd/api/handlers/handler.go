package handlers

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"miniTikTok/pkg/errno"
)

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type RegisterResponse struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
	ID      int64  `json:"user_id"`
	Token   string `json:"token"`
}

func SendRegisterResponse(c *app.RequestContext, err error, id int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, RegisterResponse{
		Code:    int64(Err.ErrCode),
		Message: Err.ErrMsg,
		ID:      id,
		Token:   token,
	})
}

type LoginResponse struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
	ID      int64  `json:"user_id"`
	Token   string `json:"token"`
}

func SendLoginResponse(c *app.RequestContext, err error, id int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, LoginResponse{
		Code:    int64(Err.ErrCode),
		Message: Err.ErrMsg,
		ID:      id,
		Token:   token,
	})
}

// UserInfoParam get请求 tag要设置query
// https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/binding-and-validate/
type UserInfoParam struct {
	ID    string `json:"user_id" form:"user_id" query:"user_id"`
	Token string `json:"token" form:"token" query:"token"`
}

type User struct {
	ID            int64  `json:"id"`
	UserName      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type InfoResponse struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg"`
	User    User   `json:"user"`
}

func SendInfoResponse(c *app.RequestContext, err error, user User) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, InfoResponse{
		Code:    int64(Err.ErrCode),
		Message: Err.ErrMsg,
		User:    user,
	})
}
