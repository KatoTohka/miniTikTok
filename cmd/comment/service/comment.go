package service

import (
	"context"
	"gorm.io/gorm"
	"miniTikTok/cmd/comment/dal/db"
	"miniTikTok/kitex_gen/comment"
	"miniTikTok/middleware"
	"miniTikTok/pkg/errno"
	"miniTikTok/pkg/snowflake"
)

type CommentService struct {
	ctx context.Context
}

func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{
		ctx: ctx,
	}
}

func (s *CommentService) Comment(req *comment.DouyinCommentActionRequest) (*comment.Comment, error) {
	actionType := req.ActionType
	videoId := req.VideoId
	_, claim, _ := middleware.ParseToken(req.Token)
	userId := claim.UserID
	switch actionType {
	case 1:
		snow, _ := snowflake.NewSnowflake(1)
		commentId := snow.Generate()
		c := db.Comment{
			Model:   gorm.Model{ID: commentId},
			Content: *req.CommentText,
			VideoId: videoId,
			UserId:  userId,
		}
		err := db.PublishComment(context.Background(), &c)
		if err != nil {
			return &comment.Comment{}, err
		}
		err = db.AddCommentById(context.Background(), videoId)
		if err != nil {
			return &comment.Comment{}, err
		}
		cmt, err := db.QueryCommentById(context.Background(), commentId)
		//log.Println("7777777", len(cmt))
		createDate := cmt[0].CreatedAt
		sCreateDate := createDate.Format("2006-01-02 15:04:05")[5:10]
		if err != nil {
			return &comment.Comment{}, err
		}
		u, err := db.QueryUserByID(context.Background(), userId)
		if err != nil {
			return &comment.Comment{}, err
		}
		return &comment.Comment{
			Id: int64(commentId),
			User: &comment.User{
				Id:            int64(u[0].ID),
				Name:          u[0].UserName,
				FollowCount:   &u[0].FollowCount,
				FollowerCount: &u[0].FollowerCount,
				IsFollow:      false,
			},
			Content:    *req.CommentText,
			CreateDate: sCreateDate,
		}, err

	case 2:
		commentId := req.CommentId
		err := db.DeleteCommentById(context.Background(), uint(*commentId))
		if err != nil {
			return &comment.Comment{}, err
		}
		err = db.DecCommentById(context.Background(), videoId)
		if err != nil {
			return &comment.Comment{}, err
		}
		return &comment.Comment{}, nil
	default:
		return &comment.Comment{}, errno.ParamErr
	}
}

func (s *CommentService) CommentList(req *comment.DouyinCommentListRequest) ([]*comment.Comment, error) {

	comments, err := db.QueryCommentByVideoId(context.Background(), uint(req.VideoId))

	if err != nil {
		return []*comment.Comment{}, err
	}
	var resp []*comment.Comment
	for _, cmt := range comments {
		var c comment.Comment
		createDate := cmt.CreatedAt
		sCreateDate := createDate.Format("2006-01-02 15:04:05")[5:10]
		u, err := db.QueryUserByID(context.Background(), cmt.UserId)
		if err != nil {
			return []*comment.Comment{}, err
		}
		var user comment.User
		user.Id = int64(u[0].ID)
		user.Name = u[0].UserName
		user.FollowerCount = &u[0].FollowerCount
		user.FollowCount = &u[0].FollowCount
		user.IsFollow = false
		c.Id = int64(cmt.ID)
		c.CreateDate = sCreateDate
		c.Content = cmt.Content
		c.User = &user
		resp = append(resp, &c)
	}

	return resp, nil
}
