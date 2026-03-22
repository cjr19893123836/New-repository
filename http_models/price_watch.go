package http_models

import (
	"time"

	"gorm.io/gorm"
)

// PriceWatch 蹲降价表模型
type PriceWatch struct {
	gorm.Model
	UserID      int64     `gorm:"column:user_id"`                                      //用户ID（关联用户表）
	GoodsID     int64     `gorm:"column:goods_id"`                                     //商品ID（关联商品表）
	ExpectPrice float64   `gorm:"column:expect_price;type:decimal(10,2);default:null"` //期望降价
	IsNotified  bool      `gorm:"column:is_notified;default:false"`                    //是否已通知（0-未通知/1-已通知）
	CreatedTime time.Time `gorm:"column:created_time"`                                 //创建时间
}

func (PriceWatch) TableName() string {
	return "price_watch"
}
