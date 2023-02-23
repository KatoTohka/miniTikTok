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

func Comment(ctx context.Context, c *app.RequestContext) {
	hlog.Info("comment log")
	var commentVar CommentParam
	if err := c.Bind(&commentVar); err != nil {
		SendCommentResponse(c, Comments{}, errno.NewErrNo(777,"1"))
		return
	}
	if len(commentVar.Token) == 0 || len(commentVar.VideoID) == 0 || len(commentVar.ActionType) == 0 {
		SendCommentResponse(c, Comments{}, errno.NewErrNo(777,"2"))
		return
	}
	//JWT
	token, _, err := middleware.ParseToken(commentVar.Token)
	if err != nil || !token.Valid {
		SendCommentResponse(c, Comments{}, errno.AuthorizationFailedErr)
		return
	}
	videoId, err := strconv.ParseInt(commentVar.VideoID, 10, 64)
	if err != nil {
		SendCommentResponse(c, Comments{}, errno.NewErrNo(777,"3"))
		return
	}
	var id int64
	if len(commentVar.ID) == 0 {
		id = 0
	} else {
		id, err = strconv.ParseInt(commentVar.ID, 10, 64)
		if err != nil {
			SendCommentResponse(c, Comments{}, errno.ParamErr)
			return
		}
	}

	actionType, err := strconv.Atoi(commentVar.ActionType)
	if err != nil {
		SendCommentResponse(c, Comments{}, errno.ParamErr)
		return
	}
	cmt, err := rpc.Comment(context.Background(), &comment.DouyinCommentActionRequest{
		Token:       commentVar.Token,
		VideoId:     videoId,
		ActionType:  int32(actionType),
		CommentText: &commentVar.CommentText,
		CommentId:   &id,
	})
	if err != nil {
		SendCommentResponse(c, Comments{}, errno.ParamErr)
		return
	}
	if actionType == 1 {
		u := User{
			ID:            cmt.User.Id,
			UserName:      cmt.User.Name,
			FollowCount:   *cmt.User.FollowCount,
			FollowerCount: *cmt.User.FollowerCount,
			IsFollow:      cmt.User.IsFollow,
		}
		comments := Comments{
			ID:         cmt.Id,
			User:       u,
			Content:    cmt.Content,
			CreateDate: cmt.CreateDate,
		}
		SendCommentResponse(c, comments, errno.Success)
		return
	}
	SendCommentResponse(c, Comments{}, errno.Success)
}
