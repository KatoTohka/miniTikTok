package main

import (
	"context"
	"miniTikTok/cmd/feed/service"
	feed "miniTikTok/kitex_gen/feed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.DouyinFeedRequest) (*feed.DouyinFeedResponse, error) {
	var resp feed.DouyinFeedResponse
	videoList, nextTime, err := service.NewFeedService(ctx).FeedVideo(req)
	if err != nil {
		return nil, err
	}
	status := "Err"

	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = &status
		resp.VideoList = nil
		resp.NextTime = nextTime
		return &resp, err
	}
	status = "VideoListSuccess"
	resp.StatusCode = 0
	resp.StatusMsg = &status
	resp.VideoList = videoList
	return &resp, nil
}
