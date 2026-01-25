package http_models

import (
	"time"

	"gorm.io/gorm"
)

// Comment
type Comment struct {
	gorm.Model
	GoodsID     int64     `gorm:"column:goods_id"`                //商品ID（关联商品表）
	UserID      int64     `gorm:"column:user_id"`                 //用户ID（关联用户表）
	IsSeller    bool      `gorm:"column:is_seller"`               //是否是卖家（0-普通用户/1-卖家）
	Content     string    `gorm:"column:content;type:text"`       //留言内容（支持表情包）
	PublishTime time.Time `gorm:"column:publish_time"`            //发布时间
	Region      string    `gorm:"column:region"`                  //留言用户地区
	LikeCount   int64     `gorm:"column:like_count;default:0"`    //点赞数
	IsHidden    bool      `gorm:"column:is_hidden;default:false"` //是否隐藏（0-显示/1-隐藏）
}

func (Comment) TableName() string {
	return "comment"
}
