package main

import (
	"context"
	"miniTikTok/cmd/video/service"
	video "miniTikTok/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *video.DouyinPublishActionRequest) (*video.DouyinPublishActionResponse, error) {
	// true call register
	err := service.NewVideoPublishService(ctx).PublishVideo(req)
	var resp video.DouyinPublishActionResponse
	status := "PublishUploadErr"
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = &status
		return &resp, err
	}
	status = "PublishUploadSuccess"
	resp.StatusCode = 0
	resp.StatusMsg = &status
	return &resp, nil
}

// ListVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) ListVideo(ctx context.Context, req *video.DouyinPublishListRequest) (*video.DouyinPublishListResponse, error) {
	var resp video.DouyinPublishListResponse
	// true call register
	videoList, err := service.NewVideoListService(ctx).ListVideo(req)

	status := "VideoListErr"
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = &status
		resp.VideoList = nil
		return &resp, err
	}
	status = "VideoListSuccess"
	resp.StatusCode = 0
	resp.StatusMsg = &status
	resp.VideoList = videoList
	return &resp, nil
}
