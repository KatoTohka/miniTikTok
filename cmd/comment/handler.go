package main

import (
	"context"
	"miniTikTok/cmd/comment/service"
	comment "miniTikTok/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// PublishComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) PublishComment(ctx context.Context, req *comment.DouyinCommentActionRequest) (*comment.DouyinCommentActionResponse, error) {
	var resp comment.DouyinCommentActionResponse
	c, err := service.NewCommentService(ctx).Comment(req)
	status := "PublishCommentErr"
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = &status
		resp.Comment = &comment.Comment{}
		return &resp, err
	}
	status = "PublishSuccess"
	resp.StatusCode = 0
	resp.StatusMsg = &status
	resp.Comment = c
	return &resp, nil
}

// ListComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) ListComment(ctx context.Context, req *comment.DouyinCommentListRequest) (*comment.DouyinCommentListResponse, error) {
	var resp comment.DouyinCommentListResponse
	comments, err := service.NewCommentService(ctx).CommentList(req)
	status := "ListCommentErr"
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = &status
		resp.CommentList = []*comment.Comment{{}}
		return &resp, err
	}
	status = "ListSuccess"
	resp.CommentList = comments
	resp.StatusCode = 0
	resp.StatusMsg = &status
	return &resp, nil
}
