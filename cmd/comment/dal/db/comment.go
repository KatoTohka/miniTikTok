package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"miniTikTok/pkg/constants"
	"miniTikTok/pkg/errno"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"column:content"`
	VideoId int64  `gorm:"column:video_id"`
	UserId  int64  `gorm:"column:user_id"`
}

func (c *Comment) TableName() string {
	return constants.CommentTableName
}

// PublishComment add comment
func PublishComment(ctx context.Context, c *Comment) error {
	var comments []*Comment = make([]*Comment, 1)
	comments[0] = c
	if DB.WithContext(ctx).Create(comments).Error != nil {
		return errno.InsertErr
	}
	return nil
}

// QueryCommentById  query comment info by  id
func QueryCommentById(ctx context.Context, commentId uint) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := DB.WithContext(ctx).Where("id = ?", fmt.Sprint(commentId)).Find(&res).Error; err != nil {
		return []*Comment{}, err
	}
	return res, nil
}

// DeleteCommentById delete comment
func DeleteCommentById(ctx context.Context, commentId uint) error {
	if err := DB.WithContext(ctx).Where("id = ?", fmt.Sprint(commentId)).Delete(&Comment{}).Error; err != nil {
		return err
	}
	return nil

}

// QueryCommentByVideoId  query comment info by video id
func QueryCommentByVideoId(ctx context.Context, videoId uint) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := DB.WithContext(ctx).Where("video_id = ?", fmt.Sprint(videoId)).Find(&res).Error; err != nil {
		return []*Comment{}, err
	}
	return res, nil
}
