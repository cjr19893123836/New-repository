package http_models

import (
	"time"

	"gorm.io/gorm"
)

// ProductComment 商品留言表模型
type ProductComment struct {
	gorm.Model
	UserID    int64     `gorm:"column:user_id"`           //用户ID
	ProductID int64     `gorm:"column:product_id"`        //商品ID
	Content   string    `gorm:"column:content;size:1500"` //留言
	CreatedAt time.Time `gorm:"column:created_at"`        //创建时间
	Status    int       `gorm:"column:status;default:1"`  //留言状态（1 正常，0 删除，-1 违规）
}

func (ProductComment) TableName() string {
	return "product_comments"
}
