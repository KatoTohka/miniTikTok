package service

import (
	"context"
	"miniTikTok/cmd/user/dal/db"
	"miniTikTok/kitex_gen/user"
	"miniTikTok/pkg/errno"
)

type UserInfoService struct {
	ctx context.Context
}

func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

func (s *UserInfoService) InfoUser(req *user.DouyinUserRequest) (*user.User, error) {
	users, err := db.QueryUserByID(s.ctx, req.UserId)
	if err != nil {
		return nil, errno.ServiceErr
	}
	if len(users) == 0 {
		return nil, errno.UserNotExistErr
	}
	var u user.User
	u.Id = int64(users[0].ID)
	u.Name = users[0].UserName
	u.FollowCount = &users[0].FollowCount
	u.FollowerCount = &users[0].FollowerCount
	//TODO 修改follow逻辑
	u.IsFollow = false
	return &u, nil
}
