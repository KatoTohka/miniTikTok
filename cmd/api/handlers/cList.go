package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"miniTikTok/cmd/api/rpc"
	"miniTikTok/kitex_gen/comment"
	"miniTikTok/middleware"
	"miniTikTok/pkg/errno"
	"strconv"
)

func CommentList(ctx context.Context, c *app.RequestContext) {
	hlog.Info("commentList log")
	var commentsVar CommentListParam
	if err := c.Bind(&commentsVar); err != nil {
		SendCommentListResponse(c, []Comments{}, errno.ParamErr)
		return
	}
	token, _, err := middleware.ParseToken(commentsVar.Token)
	if err != nil || !token.Valid {
		SendCommentListResponse(c, []Comments{}, errno.AuthorizationFailedErr)
		return
	}
	videoId, err := strconv.ParseInt(commentsVar.VideoID, 10, 64)
	if err != nil {
		SendCommentListResponse(c, []Comments{}, errno.ParamErr)
	}
	cmtList, err := rpc.CommentList(context.Background(), &comment.DouyinCommentListRequest{
		Token:   commentsVar.Token,
		VideoId: videoId,
	})
	var commentList []Comments
	for _, c := range cmtList {
		var cmt Comments
		var u User
		u.ID = c.User.Id
		u.UserName = c.User.Name
		u.FollowerCount = *c.User.FollowerCount
		u.FollowCount = *c.User.FollowCount
		u.IsFollow = c.User.IsFollow
		cmt.ID = c.Id
		cmt.User = u
		cmt.Content = c.Content
		cmt.CreateDate = c.CreateDate
		commentList = append(commentList, cmt)
	}
	SendCommentListResponse(c, commentList, errno.Success)
}
