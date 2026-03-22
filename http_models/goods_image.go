package http_models

import "gorm.io/gorm"

// GoodsImage 商品图片表模型
type GoodsImage struct {
	gorm.Model
	GoodsID  int64  `gorm:"column:goods_id"`           //商品ID（关联商品表）
	ImageUrl string `gorm:"column:image_url;size:255"` //图片储存地址
	SortNum  int    `gorm:"column:sort_num"`           //图片展示顺序
}

func (GoodsImage) TableName() string {
	return "goods_image"
}
