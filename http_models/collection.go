package http_models

import (
	"time"

	"gorm.io/gorm"
)

// Collection 收藏表模型
type Collection struct {
	gorm.Model
	UserID         int64     `gorm:"column:user_id"`         //用户ID（关联用户表）
	GoodsID        int64     `gorm:"column:goods_id"`        //商品ID（关联商品表）
	CollectionTime time.Time `gorm:"column:collection_time"` //收藏时间
}

func (Collection) TableName() string {
	return "collection"
}
