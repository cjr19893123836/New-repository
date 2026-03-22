package http_models

import (
	"time"

	"gorm.io/gorm"
)

// CommentLike 留言点赞表模型
type CommentLike struct {
	gorm.Model
	CommentID int64     `gorm:"column:comment_id"` //评论ID（关联评论表）
	UserID    int64     `gorm:"column:user_id"`    //用户ID（关联用户表）
	LikeTime  time.Time `gorm:"column:like_time"`  //点赞时间
}

func (CommentLike) TableName() string {
	return "comment_like"
}
