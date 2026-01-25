package http_models

import (
	"time"

	"gorm.io/gorm"
)

// CommentReply 留言回复表模型
type CommentReply struct {
	gorm.Model
	CommentID     int64     `gorm:"column:comment_id"`        //评论ID（关联评论表）
	UserID        int64     `gorm:"column:user_id"`           //用户ID（关联用户表）
	Content       string    `gorm:"column:content;type:text"` //回复内容
	ReplyTime     time.Time `gorm:"column:reply_time"`        //回复时间
	RepliedUserID int64     `gorm:"column:replied_user_id"`   //被回复的用户ID
}

func (CommentReply) TableName() string {
	return "comment_reply"
}
