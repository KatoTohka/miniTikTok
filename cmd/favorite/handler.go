package main

import (
	"context"
	"miniTikTok/cmd/favorite/service"
	favorite "miniTikTok/kitex_gen/favorite"
	"miniTikTok/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// ActionFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) ActionFavorite(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	resp = new(favorite.DouyinFavoriteActionResponse)
	if req.ActionType != 1 && req.ActionType != 2 {
		resp.StatusCode = errno.BadReqErrCode
		resp.StatusMsg = errno.BadReqErr.ErrMsg
		return resp, nil
	}
	err = service.NewFavoriteActionService(ctx).ActionFavorite(req)
	if err != nil {
		resp.StatusCode = errno.ActionFavoriteErr.ErrCode
		resp.StatusMsg = errno.ActionFavoriteErr.ErrMsg
		return resp, err
	}
	resp.StatusCode = errno.SuccessCode
	resp.StatusMsg = errno.Success.ErrMsg
	return resp, nil
}

// ListFavorite implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) ListFavorite(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	resp = new(favorite.DouyinFavoriteListResponse)
	videoList, err := service.NewFavoriteListService(ctx).ListFavorite(req)
	if err != nil {
		resp.StatusCode = errno.ListFavoriteErrCode
		resp.StatusMsg = errno.ListFavoriteErr.ErrMsg
		resp.VideoList = videoList
		return resp, err
	}
	resp.StatusCode = errno.SuccessCode
	resp.StatusMsg = errno.Success.ErrMsg
	resp.VideoList = videoList
	return resp, nil
}
